# Makefile

include .env

.EXPORT_ALL_VARIABLES:	

CGO_ENABLED=0

build:
	go build -ldflags "${XFLAGS} -s -w" -a -o ./bin/telegram-amqp-amd64
	GOARCH=arm go build -ldflags "${XFLAGS} -s -w" -a -o ./bin/telegram-amqp-arm

dev:
	go run main.go

docker-build:
	docker build -t telegram-amqp .

docker-run: docker-build

	@docker run -it --rm -e RABBITMQ_USERNAME="${RABBITMQ_USERNAME}" -e RABBITMQ_PASSWORD="${RABBITMQ_PASSWORD}" -e RABBITMQ_IP="${RABBITMQ_IP}" -e RABBITMQ_PORT="${RABBITMQ_PORT}" telegram-amqp

.PHONY: dev
