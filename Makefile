.PHONY: all build clean run check

RAFT_PEER=raft_peer
SCHEDULER=scheduler
ACCESSOR=accessor
PORT1=2001
PORT2=2002
PORT3=2003
C_PORT1=3001
C_PORT2=3002
C_PORT3=3003

all: check build-raft build-scheduler

clean:
	rm ./"${RAFT_PEER}"

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    	./api/raftrpc/raft.proto

	protoc --go_out=. --go_opt=paths=source_relative \
    	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
        ./api/register/register.proto

check:
	@go fmt ./...
	@go vet ./...

build-raft:
	@go mod tidy
	@go build -o ./cmd/raft/"${RAFT_PEER}" ./cmd/raft/main.go

run-raft:
	nohup ./cmd/raft/"${RAFT_PEER}" -env dev -raft-port ${PORT1} -client-port ${C_PORT1} >>logs/1.log &
	nohup ./cmd/raft/"${RAFT_PEER}" -env dev -raft-port ${PORT2} -client-port ${C_PORT2} >>logs/2.log &
	nohup ./cmd/raft/"${RAFT_PEER}" -env dev -raft-port ${PORT3} -client-port ${C_PORT3} >>logs/3.log &

build-scheduler:
	@go mod tidy
	@go build -o ./cmd/scheduler/"${SCHEDULER}" ./cmd/scheduler/main.go

run-scheduler:
	nohup ./cmd/scheduler/"${SCHEDULER}" -env dev >> logs/scheduler.log &

build-accessor:
	@go mod tidy
	@go build -o ./cmd/accessor/"${ACCESSOR}" ./cmd/accessor/main.go &

run-accessor:
	nohup ./cmd/accessor/"${ACCESSOR}" -env dev >> log/accessor.log &
