down:
	docker-compose down -v

build:
	docker-compose build

protoc:
	protoc protoc/example.proto --go_out=plugins=grpc:.
