# syntax=docker/dockerfile:1

FROM golang:1.17-alpine as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY ./ ./

RUN go mod tidy \
    && go build -o ./cmd/scheduler/scheduler ./cmd/scheduler/main.go

FROM alpine:latest as prod

WORKDIR /root/

COPY --from=0 /app/cmd/scheduler/scheduler .
COPY --from=0 /app/config/ ./config

EXPOSE 2888 2889

ENTRYPOINT ["/bin/sh", "-c", "./scheduler -env dev"]
