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

run-raft-multi-cluster:
	nohup ./cmd/raft/"${RAFT_PEER}" -env dev -raft-id raft000 -raft-port 2001 -client-port 3001 >> logs/cluster0-1.log &
	nohup ./cmd/raft/"${RAFT_PEER}" -env dev -raft-id raft000 -raft-port 2002 -client-port 3002 >> logs/cluster0-2.log &
	nohup ./cmd/raft/"${RAFT_PEER}" -env dev -raft-id raft000 -raft-port 2003 -client-port 3003 >> logs/cluster0-3.log &

	nohup ./cmd/raft/"${RAFT_PEER}" -env dev -raft-id raft001 -raft-port 2101 -client-port 3101 >> logs/cluster1-1.log &
	nohup ./cmd/raft/"${RAFT_PEER}" -env dev -raft-id raft001 -raft-port 2102 -client-port 3102 >> logs/cluster1-2.log &
	nohup ./cmd/raft/"${RAFT_PEER}" -env dev -raft-id raft001 -raft-port 2103 -client-port 3103 >> logs/cluster1-3.log &

	nohup ./cmd/raft/"${RAFT_PEER}" -env dev -raft-id raft002 -raft-port 2201 -client-port 3201 >> logs/cluster2-1.log &
	nohup ./cmd/raft/"${RAFT_PEER}" -env dev -raft-id raft002 -raft-port 2202 -client-port 3202 >> logs/cluster2-2.log &
	nohup ./cmd/raft/"${RAFT_PEER}" -env dev -raft-id raft002 -raft-port 2203 -client-port 3203 >> logs/cluster2-3.log &

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
