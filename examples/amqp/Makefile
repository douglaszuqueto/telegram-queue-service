# Makefile

include .env

.EXPORT_ALL_VARIABLES:	

CGO_ENABLED=0

build:
	go build -ldflags "${XFLAGS} -s -w" -a -o ./bin/telegram-amqp-amd64
	GOARCH=arm go build -ldflags "${XFLAGS} -s -w" -a -o ./bin/telegram-amqp-arm

dev:
	go run main.go

.PHONY: dev
