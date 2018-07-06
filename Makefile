APP_NAME := vaccichain.exe

all:

build-server:
	@go build -o ./server/$(APP_NAME) ./server/*.go
	@ls -la ./server/$(APP_NAME)

run-server:
	@cd ./server/ && ./$(APP_NAME)

get-server-deps:
	go get -u google.golang.org/grpc
	go get -u github.com/golang/protobuf/protoc-gen-go

build-protocol:
	@protoc --proto_path=./protocol/ ./protocol/vaccichain.proto --go_out=plugins=grpc:protocol