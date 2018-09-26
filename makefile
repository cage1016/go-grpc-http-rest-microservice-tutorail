# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build

build-server:
	cd cmd/server && $(GOBUILD) -o server

run-server: build-server
	./cmd/server/server \
	-grpc-port=9090 \
	-http-port=8080 \
	-db-host=127.0.0.1:3306 \
	-db-user=root \
	-db-password=@75dkYz9n \
	-db-schema=go-grpc-http-rest-microservice-tutorial \
	-swagger-dir=api/swagger/v1 \
	-log-level=-1 \
	-log-time-format=2006-01-02T15:04:05.999999999Z07:00

build-client-grpc:
	cd cmd/client-grpc/ && $(GOBUILD) -o client-grpc

run-client-grpc: build-client-grpc
	./cmd/client-grpc/client-grpc -server=localhost:9090

build-client-rest:
	cd cmd/client-rest/ && $(GOBUILD) -o client-rest

run-client-rest: build-client-rest
	./cmd/client-rest/client-rest -server=http://localhost:8080