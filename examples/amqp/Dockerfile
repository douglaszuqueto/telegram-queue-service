FROM golang:1.11 as builder
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -ldflags "${XFLAGS} -s -w" -a -o telegram-amqp-amd64

FROM alpine
RUN apk update && apk add --no-cache ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /app
COPY ./bin/telegram-amqp-amd64 /app
CMD ["./telegram-amqp-amd64"]