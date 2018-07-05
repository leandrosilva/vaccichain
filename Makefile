APP_NAME := buckychain.exe

all: build

build:
	@go build -o $(APP_NAME) *.go

run:
	@./$(APP_NAME)
