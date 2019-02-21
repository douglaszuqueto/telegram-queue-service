#!/bin/bash

CGO_ENABLED=0 go build -ldflags "${XFLAGS} -s -w" -a -o ./bin/telegram-service-amd64
CGO_ENABLED=0 GOARCH=arm go build -ldflags "${XFLAGS} -s -w" -a -o ./bin/telegram-service-arm