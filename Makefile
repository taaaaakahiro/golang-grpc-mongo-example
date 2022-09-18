up:
	docker-compose up -d

down:
	docker-compose down -v

build:
	docker-compose build

restart:
	docker-compose restart

proto:
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./protoc/example.proto

run:
	go run server/main.go