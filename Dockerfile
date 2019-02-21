FROM alpine

RUN apk update && apk add --no-cache ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app

COPY ./bin/telegram-service-amd64 /app

CMD ["./telegram-service-amd64"]