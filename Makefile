up:
	docker-compose up -d

down:
	docker-compose down -v

build:
	docker-compose build

restart:
	docker-compose restart

protoc:
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./protoc/example.proto

ser:
	go run server/server.go

cli:
	go run client/client.go