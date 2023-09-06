# segment-service-api 

<details>
<summary> 
	Project tree
</summary>

```
📦 segment-service-api
├─ .github
│  └─ workflows
│     └─ go.yml
├─ .gitignore
├─ Dockerfile
├─ Makefile
├─ README.md
├─ api
│  ├─ segment
│  │  └─ segment.proto
│  └─ user
│     └─ user.proto
├─ cmd
│  └─ server
│     └─ segment_service.go
├─ config.yml
├─ database
│  └─ init.sql
├─ docker-compose.yml
├─ go.mod
├─ go.sum
├─ internal
│  ├─ api
│  │  ├─ segment_v1
│  │  │  ├─ add_segment.go
│  │  │  ├─ add_segment_test.go
│  │  │  ├─ remove_segment.go
│  │  │  ├─ remove_segment_test.go
│  │  │  └─ segment_v1.go
│  │  └─ user_v1
│  │     ├─ add_user.go
│  │     ├─ add_user_test.go
│  │     ├─ get segments.go
│  │     ├─ get_history.go
│  │     ├─ get_history_test.go
│  │     ├─ get_segments_test.go
│  │     ├─ modify_segments.go
│  │     ├─ modify_segments_test.go
│  │     ├─ remove_user.go
│  │     ├─ remove_user_test.go
│  │     ├─ set_expire_time.go
│  │     ├─ set_expire_time_test.go
│  │     └─ user_v1.go
│  ├─ app
│  │  ├─ app.go
│  │  └─ service_provider.go
│  ├─ client
│  │  └─ db
│  │     ├─ client.go
│  │     ├─ db.go
│  │     ├─ mocks_db
│  │     │  └─ mock_db.go
│  │     ├─ mocks_tx
│  │     │  └─ mock_tx.go
│  │     └─ transaction
│  │        └─ transaction.go
│  ├─ config
│  │  └─ config.go
│  ├─ convert
│  │  └─ convert.go
│  ├─ model
│  │  └─ model.go
│  ├─ repository
│  │  ├─ mocks
│  │  │  ├─ segment_mocks
│  │  │  │  └─ segment_service_repository.go
│  │  │  └─ user_mocks
│  │  │     └─ user_service_repository.go
│  │  ├─ segment
│  │  │  ├─ add_segment.go
│  │  │  ├─ remove_segment.go
│  │  │  └─ segment.go
│  │  ├─ table
│  │  │  └─ tables.go
│  │  └─ user
│  │     ├─ add_user.go
│  │     ├─ get_history.go
│  │     ├─ get_segment_id.go
│  │     ├─ get_segments.go
│  │     ├─ get_user.go
│  │     ├─ modify_segments.go
│  │     ├─ remove_user.go
│  │     ├─ set_expire_time.go
│  │     └─ user.go
│  └─ service
│     ├─ segment
│     │  ├─ add_segment.go
│     │  ├─ remove_segment.go
│     │  └─ segment.go
│     └─ user
│        ├─ add_user.go
│        ├─ get_history.go
│        ├─ get_segments.go
│        ├─ modify_segments.go
│        ├─ remove_user.go
│        ├─ set_expire_time.go
│        └─ user.go
├─ pkg
│  ├─ segment_api
│  │  ├─ api.swagger.json
│  │  ├─ segment.pb.go
│  │  ├─ segment.pb.gw.go
│  │  ├─ segment.pb.validate.go
│  │  └─ segment_grpc.pb.go
│  └─ user_api
│     ├─ api.swagger.json
│     ├─ user.pb.go
│     ├─ user.pb.gw.go
│     ├─ user.pb.validate.go
│     └─ user_grpc.pb.go
└─ readme_assets
   └─ logo-avito.png
```
</details>

## Brief description

<p align="justify">
	
Это сервис для динамического назначения пользователей сегментам в рамках проекта для отбора в Avito. Все текстовые значения, передаваемые сервису не должны превышать длину строки 100 (пустые строки запрещены). Массивы значений должны содержать не повторяющиеся значения. Для выполнения этих условий организована валидация. Сервис принимает запросы как с помощью gRPC, так и с помощью HTTP, в нём частично реализлована идея model-view-control, однако слой-адаптер был сокращен ввиду простоты выполняемых операций. Так как в задании, выданном авито, операция модификации сегментов, привязанных к пользователю, обозначена не атомарной (одновременно и добавление и удаление сегментов), была применена оболочка, позволяющая превращать операции в транзакции (в методе ModifySegments service-слоя). <br /> **Дополнительное задание 1**. Было реализовано с помощью хранения времени добавления записи о связи пользователя и сегмента в таблицу-связку (многие-ко-многим). При этом при отвязке сегмента от пользователя сменяется флаг и запись превращается в архивную (id клиента, id сегмента, время создания связи, время ее удаления, статус). При удалении пользователя или сегмента, связанные с ним записи в таблице-связке каскадно удаляются. <br /> **Дополнительное задание 2**. Было реализовано с помощью поля, отвечающего за время удаления (в случае операции удаления в это поле записывается текущее время, а в случае назначения Time-To-Live, используется отдельный метод, который назначает время отвязки сегмента от пользователя). Процесс исполнения этой отвязки вне рамок задания. <br /> Для работы сервиса минимально необходим [Docker](https://www.docker.com/)  а также Linux-система, либо WSL для запуска сервиса в контейнере. Все методы покрыты тестами, кроме метода, работающего с потоком буферизованных байтов.
</justify> <br /> 
 <img  src="./readme_assets/logo-avito.png" width="100%">

## Project setup


### Installation

<p align="justify">
	
Для сборки сервиса на базе данного репозитория нужны следующие утилиты:
- Makefile
- Protocol Buffer Compiler ([protoc](https://github.com/protocolbuffers/protobuf/releases))
- [Docker](https://www.docker.com/)
- [Golang](https://go.dev/dl/)
	
Параметры соединения с базой данных находятся в файле **config.yml**, а также в **docker-compose**. При изменении конфигурационного файла, необходимо также внести изменения в файл **docker-compose**. При локальном запуске сервиса имя хоста базы данных localhost, в случае контейнеризованного запуска -- postgres. Для запуска сервиса нужно ввести следующие команды:
```
git clone https://github.com/nikitads9/segment-service-api.git
cd segment-service-api/
make deps
make vendor-proto
make generate
docker-compose up -d
```
- `make deps` команда устанавливает зависимости для работы проекта.
- `make vendor-proto` скачивает необходимые для работы protobuf и validate структур пакеты. После выполнения этой команды в корне проекта появится папка со всеми необходимыми `.proto` файлами.
- `make generate` создает: `grpc.pb.go`, `pb.go`, `pb.gw.go` и `pb.validate.go` для сервисов **user** и **segment** на основании контрактов, описанных в **user_v1.proto** и **segment_v1.proto**. Эти файлы содержат структуры, интерфейсы и методы для работы API, сгенерированные с помощью Protobuffer.
- `docker-compose up -d` скачивает образ **alpine3.15** с DockerHub (если он еще не загружен), собирает исполняемый файл и два контейнера: один для серверного приложения с API, а второй для непосредственно сервера базы данных. Оба контейнера организуются в связанную сеть докер, которая позволяет им обращаться друг к другу по имени. Флаг -d позволяет вернуть управление терминалом пользователю без остановки контейнера. <br />
Для генерации html файла с результатами покрытия тестами кода необходимо ввести следующую команду:
```
make test-coverage
```

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
<details>
<summary> 
8. метод GetUserHistoryCsv
</summary>
  
**GET** `host:port/user/download-history/{id}` <br />
У запроса нет тела, передается только id пользователя. В ответ приходит такая структура, в которой содержится массив байтов csv файла с историей сегментов пользователя:
```
	"result": {
		"chunk": "string"
	}
}
```
</details>
