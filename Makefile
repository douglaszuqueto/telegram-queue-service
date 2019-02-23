# Makefile

include .env

.EXPORT_ALL_VARIABLES:	

CGO_ENABLED=0

build:
	go build -ldflags "${XFLAGS} -s -w" -a -o ./bin/telegram-service-amd64
	go build -ldflags "${XFLAGS} -s -w" -a -o ./bin/telegram-service-arm

dev:
	go run main.go

docker-build:
	docker build -t telegram-service .

docker-run: docker-build

	@docker run -it --rm -e TELEGRAM_TOKEN="${TELEGRAM_TOKEN}" -e TELEGRAM_CHATID="${TELEGRAM_CHATID}" -e RABBITMQ_USERNAME="${RABBITMQ_USERNAME}" -e RABBITMQ_PASSWORD="${RABBITMQ_PASSWORD}" -e RABBITMQ_IP="${RABBITMQ_IP}" -e RABBITMQ_PORT="${RABBITMQ_PORT}" telegram-service

deploy-rpi: build
	scp ./bin/telegram-service-arm ${RASPBERRY_USER}@${RASPBERRY_IP}:~/

.PHONY: dev
