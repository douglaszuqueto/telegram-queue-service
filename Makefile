# Makefile

include .env

.EXPORT_ALL_VARIABLES:	

CGO_ENABLED=0

build:
	go build -ldflags "${XFLAGS} -s -w" -a -o ./bin/telegram-service-amd64
	go build -ldflags "${XFLAGS} -s -w" -a -o ./bin/telegram-service-arm

dev:
	go run main.go

.PHONY: dev
