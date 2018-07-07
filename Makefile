LOCAL_PATH  := /c/Users/leand/go
GITHUB_PATH := github.com/leandrosilva/vaccichain
PROTOC_PATH := $(GITHUB_PATH)/protocol
SERVER_PATH := $(GITHUB_PATH)/server
CLIENT_PATH := $(GITHUB_PATH)/client

SERVER_NAME := vaccichain

all:

build-server:
	@go build -o ./server/$(SERVER_NAME) ./server/*.go
	@ls -la ./server/$(SERVER_NAME)

run-server:
	@cd ./server/ && ./$(SERVER_NAME)

test-server:
	@cd ./server/ && go test -timeout 30s

get-server-deps:
	go get -u google.golang.org/grpc
	go get -u github.com/golang/protobuf/protoc-gen-go

build-protocol:
	@protoc --proto_path=./protocol/ ./protocol/vaccichain.proto --go_out=plugins=grpc:protocol

install-protocol: build-protocol
	@mkdir -p $(LOCAL_PATH)/src/$(PROTOC_PATH)
	@cp ./protocol/vaccichain.pb.go $(LOCAL_PATH)/src/$(PROTOC_PATH)
