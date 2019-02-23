FROM golang:1.11 as builder
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -ldflags "${XFLAGS} -s -w" -a -o telegram-service-amd64

FROM alpine
RUN apk update && apk add --no-cache ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=builder /app/telegram-service-amd64 /app
CMD ["./telegram-service-amd64"]