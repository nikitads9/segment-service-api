# segment-service-api 

## Brief description

<p align="justify">
	
Это сервис для динамического назначения пользователей сегментам. Все текстовые значения, передаваемые сервису не должны превышать длину строки 100 (пустые строки запрещены). Массивы значений должны содержать не повторяющиеся значения. Для выполнения этих условий организована валидация. Сервис принимает запросы как с помощью gRPC, так и с помощью HTTP, в нём частично реализлована идея model-view-control, однако слой-адаптер был исключен ввиду простоты выполняемых операций. Так как в задании, выданном авито операция модификации сегментов, привязанных к пользователю обозначена не атомарной (одновременно и добавление и удаление сегментов), была применена оболочка, позволяющая превращать операции в транзакции (в методе ModifySegments service-слоя). <br /> **Дополнительное задание 1**. Было реализовано с помощью хранения времени добавления записи о связи пользователя и сегмента в таблицу-связку (многие-ко-многим). При этом при отвязке сегмента от пользователя сменяется флаг и запись превращается в архивную (id клиента, id сегмента, время создания связи, время ее удаления, статус). При удалении пользователя или сегмента, связанные с ним записи в таблице-связке каскадно удаляются. <br /> **Дополнительное задание 2**. Было реализовано с помощью поля, отвечающего за время удаления (в случае операции удаления в это поле записывается текущее время, а в случае назначения Time-To-Live, используется отдельный метод, который назначает время отвязки сегмента от пользователя. Процесс исполнения этой отвязки вне рамок задания. <br /> Для работы сервиса минимально необходим [Docker](https://www.docker.com/)  а также Linux-система, либо WSL для запуска сервиса в контейнере.
</justify>

## Project setup


### Installation

<p align="justify">
	
Для сборки сервиса на базе данного репозитория нужны следующие утилиты:
- Makefile
- Protocol Buffer Compiler ([protoc](https://github.com/protocolbuffers/protobuf/releases))
- [Docker](https://www.docker.com/)
- ([Golang](https://go.dev/dl/))
	
Параметры соединения с базой данных находятся в файле **config.yml**, а также в **docker-compose**. При изменении конфигурационного файла, необходимо также внести изменения в файл **docker-compose**. Для запуска сервиса нужно ввести следующие команды:
```
git clone https://github.com/nikitads9/segment-service-api.git
cd segment-service-api/
make deps
make vendor-proto
make generate
docker-compose up -d
```
- The `make deps` command installs dependencies required for this project.
- The `make vendor-proto` command downloads the required tools for protobuf and validate to work. Running this command will create proto folder in the root of the project with all necessary `.proto` files.
- The `make generate` command creates three files: `grpc.pb.go`, `pb.go`, `pb.gw.go` based on API description in **note_v1.proto**. These files contain golang structs, interfaces and golang methods generated on the basis of Protobuffer interface description.
- The `docker-compose up -d` command downloads **alpine3.15** image from Docker Hub (if you don't have it locally), builds a binary and creates two containers: one for server app which is the the API service itself and the second one acts as database server. Both containers are connected to default Docker network which enables the two containers to communicate successfully. 

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
