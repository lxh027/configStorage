# syntax=docker/dockerfile:1

FROM golang:1.17-alpine as builder

RUN apk add build-base

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY ./ ./

RUN go mod tidy \
  && go build -o ./cmd/accessor/accessor ./cmd/accessor/main.go

FROM alpine:latest as prod

RUN apk add tzdata

WORKDIR /root/

COPY --from=0 /app/cmd/accessor/accessor .
COPY --from=0 /app/config/ ./config
COPY --from=0 /app/web/ ./web

EXPOSE 5000

ENTRYPOINT ["/bin/sh", "-c", "./accessor -env dev"]
