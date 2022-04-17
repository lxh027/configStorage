# syntax=docker/dockerfile:1

FROM golang:1.17-alpine as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY ./ ./

RUN go mod tidy \
  && go build -o ./cmd/raft/raft_peer ./cmd/raft/main.go

FROM alpine:latest as prod

WORKDIR /root/

COPY --from=0 /app/cmd/raft/raft_peer .
COPY --from=0 /app/config/ ./config

EXPOSE 2000 3000

ENTRYPOINT ["/bin/sh", "-c", "./raft_peer -env dev -raft-id $RFID -raft-host $RFHOST -raft-port $RFPT -client-port $CPT"]
