APP_NAME := vaccichain.exe

all:

build-server:
	@go build -o ./server/$(APP_NAME) ./server/*.go

run-server:
	@cd ./server/ && ./$(APP_NAME)
