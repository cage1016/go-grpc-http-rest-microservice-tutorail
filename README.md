# go grpc http rest microservice tutorail
> [Tutorial, Part 1] How to develop Go gRPC microservice with HTTP/REST endpoint, middleware… - https://is.gd/jiRNj3

基本就是照著 Tutorial 的說明自己實際操作一次

## Prerequisites
- 因為有使用到 go module, 所以需要使用 `go1.11`
- MySQL. 

## Getting Started

1. git clone

```sh
$ git clone git@github.com:cage1016/go-grpc-http-rest-microservice-tutorail.git
```

2. database

- database name: `go-grpc-http-rest-microservice-tutorial`
- database user: `qeek`
- database password: `qnap1688`

建立 database

```
$ docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=qnap1688 -e MYSQL_DATABASE=go-grpc-http-rest-microservice-tutorial -e MYSQL_USER=qeek -e MYSQL_PASSWORD=qnap1688 -v $(pwd)/configs:/docker-entrypoint-initdb.d -d mysql
```

執行 server 

```
$ make run-server
cd cmd/server && go build -o server
./cmd/server/server \
	-grpc-port=9090 \
	-http-port=8080 \
	-db-host=127.0.0.1:3306 \
	-db-user=qeek \
	-db-password=qnap1688 \
	-db-schema=go-grpc-http-rest-microservice-tutorial \
	-swagger-dir=api/swagger/v1 \
	-log-level=-1 \
	-log-time-format=2006-01-02T15:04:05.999999999Z07:00
{"level":"info","ts":"2018-09-26T16:19:39.655215+08:00","msg":"starting HTTP/REST gateway..."}
{"level":"info","ts":"2018-09-26T16:19:39.655299+08:00","msg":"starting gRPC server..."}
{"level":"info","ts":"2018-09-26T16:19:39.655338+08:00","msg":"ccResolverWrapper: sending new addresses to cc: [{localhost:9090 0  <nil>}]","system":"grpc","grpc_log":true}
{"level":"info","ts":"2018-09-26T16:19:39.655375+08:00","msg":"ClientConn switching balancer to \"pick_first\"","system":"grpc","grpc_log":true}
{"level":"info","ts":"2018-09-26T16:19:39.65545+08:00","msg":"pickfirstBalancer: HandleSubConnStateChange: 0xc00018c060, CONNECTING","system":"grpc","grpc_log":true}
{"level":"info","ts":"2018-09-26T16:19:39.657614+08:00","msg":"pickfirstBalancer: HandleSubConnStateChange: 0xc00018c060, READY","system":"grpc","grpc_log":true}
```

執行 grpc client

```
$ make run-client-grpc
cd cmd/client-grpc/ && go build -o client-grpc
./cmd/client-grpc/client-grpc -server=localhost:9090
2018/09/26 16:20:44 Create result: <api:"v1" id:32 >

2018/09/26 16:20:44 Read result: <api:"v1" toDo:<id:32 title:"title (2018-09-26T08:20:44.19278Z)" description:"description (2018-09-26T08:20:44.19278Z)" reminder:<seconds:1537950044 > > >

2018/09/26 16:20:44 Update result: <api:"v1" updated:1 >

2018/09/26 16:20:44 ReadAll result: <api:"v1" toDos:<id:18 title:"hello work" description:"string" reminder:<seconds:1537938582 > > toDos:<id:32 title:"title (2018-09-26T08:20:44.19278Z)" description:"description (2018-09-26T08:20:44.19278Z) + updated" reminder:<seconds:1537950044 > > >

2018/09/26 16:20:44 Delete result: <api:"v1" deleted:1 >
```

執行 rest client

```
$ make run-client-rest
cd cmd/client-rest/ && go build -o client-rest
./cmd/client-rest/client-rest -server=http://localhost:8080
2018/09/26 16:21:18 Create response: Code=200, Body={"api":"v1","id":"33"}

2018/09/26 16:21:18 Read response: Code=200, Body={"api":"v1","toDo":{"id":"33","title":"title (2018-09-26T08:21:18.371247Z)","description":"description (2018-09-26T08:21:18.371247Z)","reminder":"2018-09-26T08:21:18Z"}}

2018/09/26 16:21:18 Update response: Code=200, Body={"api":"v1","updated":"1"}

2018/09/26 16:21:18 ReadAll response: Code=200, Body={"api":"v1","toDos":[{"id":"18","title":"hello work","description":"string","reminder":"2018-09-26T05:09:42Z"},{"id":"33","title":"title (2018-09-26T08:21:18.371247Z) + updated","description":"description (2018-09-26T08:21:18.371247Z) + updated","reminder":"2018-09-26T08:21:18Z"}]}

2018/09/26 16:21:18 Delete response: Code=200, Body={"api":"v1","deleted":"1"}
```

Swagger

- localhost:8080/swagger/todo-service.swagger.json - http://localhost:8080/swagger/todo-service.swagger.json
- Swagger UI - http://localhost:8080/swagger-ui/#/
