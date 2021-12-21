.PHONY: all build clean run check

RAFT_PEER=raft_peer

all: check build

build:
	@go mod tidy
	@go build -o "${RAFT_PEER}" ./cmd/raft/main.go

check:
	@go fmt ./...
	@go vet ./...

run:
	nohup ./"${RAFT_PEER}" -env dev -id 0 >>1.log &
	nohup ./"${RAFT_PEER}" -env dev -id 1 >>2.log &
	nohup ./"${RAFT_PEER}" -env dev -id 2 >>3.log &

clean:
	rm ./"${RAFT_PEER}"
