FROM alpine

RUN apk update && apk add --no-cache ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app

COPY ./bin/telegram-amqp-amd64 /app

CMD ["./telegram-amqp-amd64"]