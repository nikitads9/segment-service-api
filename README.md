# segment-service-api 

## Brief description

<p align="justify">
	
Это сервис для динамического назначения пользователей сегментам. Все текстовые значения, передаваемые сервису не должны превышать длину строки 100 (пустые строки запрещены). Массивы значений должны содержать не повторяющиеся значения. Для выполнения этих условий организована валидация. Сервис принимает запросы как с помощью gRPC, так и с помощью HTTP, в нём частично реализлована идея model-view-control, однако слой-адаптер был исключен ввиду простоты выполняемых операций. Так как в задании, выданном авито операция модификации сегментов, привязанных к пользователю обозначена не атомарной (одновременно и добавление и удаление сегментов), была применена оболочка, позволяющая превращать операции в транзакции (в методе ModifySegments service-слоя). <br /> Дополнительное задание 1 было реализовано с помощью хранения времени добавления записи о связи пользователя и сегмента в таблицу-связку (многие-ко-многим). При этом при отвязке сегмента от пользователя сменяется флаг и запись превращается в архивную (id клиента, id сегмента, время создания связи, время ее удаления, статус). При удалении пользователя или сегмента, связанные с ним записи в таблице-связке каскадно удаляются. <br /> Дополнительное задание 2 было реализовано с помощью поля, отвечающего за время удаления (в случае операции удаления в это поле записывается текущее время, а в случае назначения Time-To-Live, используется отдельный метод, который назначает время отвязки сегмента от пользователя. Процесс исполнения этой отвязки вне рамок задания. <br /> This service requires at least [Docker](https://www.docker.com/) and [Goose](https://github.com/pressly/goose/) installed as well as using Linux or
WSL to set up the Note Service app and database in a container.
</p>

## Project setup

### Out of the box scenario

<p align="justify">
	
In case you want to just use this service out of the box, you need to verify the installation of goose and docker. If you don't have goose installed,
```
curl -fsSL \
    https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
    GOOSE_INSTALL=/usr/local/bin/.goose sh -s v3.5.0
sudo cp -r /home/$USER/.goose/bin/goose /usr/local/bin
```
Then you need to pull Docker images from my repository on DockerHub.
```
docker pull nikitads9/note-service:app
docker pull postgres:14-alpine3.15
```
When it is done, it's time to run containers using pulled images. If you want to specify your own database connection parameters, you should change the environment `-e` and port `-p` flags in the command featured below:
```
docker network create note-service-network
docker run -d -e POSTGRES_DB='notes_db' \
 -e POSTGRES_PASSWORD='notes_pass'\
 -e POSTGRES_USER='postgres'\
 -e PGDATA='/var/lib/postgresql/data/notification'\
 -p 5432:5432\
 -v postgres-volume:'/var/lib/postgresql/data'\
 --network note-service-network \
 --name postgres\
 postgres:14-alpine3.15
docker run -d --name app\
 -p 50051:50051\
 -p 8000:8000\
 -v 'app-volume:/var/lib/note-app/data'\
 --network note-service-network\
 nikitads9/note-service:app
```
**NB**: If you have changed the database configuration in `docker run` command, you should also edit the connection variables in **migration-local.sh** script file. 
And finally, when both containers are up, run this bash script for migration:
```
bash migration-local.sh
```
Now the database table is created and you can send HTTP and gRPC requests to the server app.
</p>

### Advanced installation

<p align="justify">
	
In case you want to build the service yourself, you will need to have these tools installed:
- Makefile
- Goose
- Protocol Buffer Compiler ([protoc](https://github.com/protocolbuffers/protobuf/releases))
- Docker
- Golang
	
If you are ok with that, be sure to edit database connection parameters in **config.yml** file among with **Dockerfile** and **migration-local.sh**. The commands to launch the server app and database are listed below:
```
git clone https://github.com/nikitads9/note-service-api.git
cd note-service-api/
make deps
make vendor-proto
make generate
docker-compose up -d
curl -fsSL \
    https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
    GOOSE_INSTALL=$HOME/.goose sh -s v3.5.0
sudo cp -r /home/$USER/.goose/bin/goose /usr/local/bin
bash migration-local.sh
```
- The `make deps` command installs dependencies required for this project.
- The `make vendor-proto` command downloads the required tools for protobuf and validate to work. Running this command will create proto folder in the root of the project with all necessary `.proto` files.
- The `make generate` command creates three files: `grpc.pb.go`, `pb.go`, `pb.gw.go` based on API description in **note_v1.proto**. These files contain golang structs, interfaces and golang methods generated on the basis of Protobuffer interface description.
- The `docker-compose up -d` command downloads **alpine3.15** image from Docker Hub (if you don't have it locally), builds a binary and creates two containers: one for server app which is the the API service itself and the second one acts as database server. Both containers are connected to default Docker network which enables the two containers to communicate successfully. 
- The `curl -fsSL...` command downloads goose tool for database migration and initiates the installation process. Another way of installing goose is to run ```go install github.com/pressly/goose/v3/cmd/goose@latest``` outside of this repository (goose binary will turn up at the **%GOPATH%/bin** folder).
- The `sudo cp -r /home/$USER/.goose/bin/goose /usr/local/bin` command copies goose binary file to `usr/local/bin` folder so that your Linux could run goose commands from anywhere.
- The `bash migration-local.sh` command starts the bash script, that completes database migration specified in `.sql` file in **/migrations** folder. The parameters required for database connection to complete migration are specified in **migration-local.sh**.

</justify>

## API use instruction

Этот сервис частично реализует концепцию CRUD. С его помощью возможно создавать и удалять сегменты и пользователей, а также назначать сегменты пользователям и прекращать участие пользователей в сегментах. При прекращении связи пользователя и сегмента запись об участии пользователя в сегменте сохраняется с флагом **state** `false` и пометкой о времени удаления из сегмента **time_of_expire**. При удалении сегмента, записи о входящих в него пользователях полностью удаляются. Инструкция, размещенная ниже, предназначена для запросов HTTP+JSON в RESTful API стиле. Для тестирования gRPC клиента см. контракты в файлах **segment.proto** и **user.proto**.
<details>
<summary> 
1. метод AddUser 
</summary>
  
**POST** `host:port/user/add-user` <br />
Объект JSON, передаваемый этому методу должен выглядеть так:
```
{
  "user_name": "user1"
}
```
Метод возвращает объект JSON с вложенным id добавленного пользователя
```
{
	"id": "1"
}
```
</details>
<details>
<summary> 
2. метод GetSegments
</summary>
  
**GET** `host:port/user/get-segments/{id}` <br />
Этот метод не нуждается в JSON объекте. Вместо этого требуется id сегмента.
Метод возвращает JSON с  массивом названий сегментов:
```
{
	"slugs": [
		"segment1"
	]
}
```
</details>
<details>
<summary> 
3. метод ModifySegments
</summary>
  
**PATCH** `host:port/user/modify-segments` <br />
Объект JSON, передаваемый этому методу должен выглядеть так:
```
{
  "id": "1",
  "slug_to_add": [
    "segment1"
  ],
  "slug_to_remove": [
    "segment2"
  ]
}
```
Возвращаемое значение должно выглядеть как пустой объект JSON.
</details>
<details>
<summary> 
4. метод RemoveUser 
</summary>
  
**DELETE** `host:port/user/remove-user/{id}` <br />
У запроса нет тела. Возвращаемое значение должно выглядеть как пустой объект JSON.
</details>
<details>
<summary> 
5. метод SetExpireTime
</summary>
  
**POST** `host:port/user/set-expire-time` <br />
Объект JSON, передаваемый этому методу должен выглядеть так:
```
{
  "id": "1",
  "slug": "segment1",
  "expiration_time": "2023-08-31T09:47:57.917Z"
}
```
Возвращаемое значение должно выглядеть как пустой объект JSON.
</details>
<details>
<summary> 
6. метод AddSegment
</summary>
  
**POST** `host:port/segment/add-segment` <br />
Объект JSON, передаваемый этому методу должен выглядеть так:
```
{
  "slug": "segment1"
}
```
Метод возвращает id добавленного сегмента. 
```
{
	"id": "1"
}
```
</details>
<details>
<summary> 
7. метод RemoveSegment
</summary>
  
**DELETE** `host:port/segment/remove-segment/{id}` <br />
У запроса нет тела. Возвращаемое значение должно выглядеть как пустой объект JSON.
</details>
