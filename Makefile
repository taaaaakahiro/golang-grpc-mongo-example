down:
	docker-compose down -v

build:
	docker-compose build

proto:
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./protoc/example.proto
