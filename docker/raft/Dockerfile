# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY ./ ./

RUN go mod tidy \
  && go build -o ./cmd/raft/raft_peer ./cmd/raft/main.go

EXPOSE 2000 3000

ENTRYPOINT ["/bin/sh", "-c", "./cmd/raft/raft_peer -env dev -raft-id $RFID -raft-port $RFPT -client-port $CPT >> logs/cluster0-1.log"]
