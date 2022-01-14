.PHONY: all build clean run check

RAFT_PEER=raft_peer
SCHEDULER=scheduler

all: check build

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
	@go build -o "${RAFT_PEER}" ./cmd/raft/main.go

run-raft:
	nohup ./"${RAFT_PEER}" -env dev -id 0 >>1.log &
	nohup ./"${RAFT_PEER}" -env dev -id 1 >>2.log &
	nohup ./"${RAFT_PEER}" -env dev -id 2 >>3.log &

build-scheduler:
	@go mod tidy
	@go build -o "${SCHEDULER}" ./cmd/scheduler/main.go

run-scheduler:
	nohup ./"${SCHEDULER}" -env dev >> scheduler.log &
