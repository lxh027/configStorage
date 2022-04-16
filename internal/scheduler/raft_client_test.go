package scheduler

import (
	"configStorage/internal/raft"
	"log"
	"testing"
)

func TestStorage(t *testing.T) {
	cfg := raft.ClientConfig{
		Size: 3,
		Addresses: []string{
			"172.17.0.1:3001",
			"172.17.0.1:3002",
			"172.17.0.1:3003",
		},
	}
	client := raft.NewRaftClient(cfg)

	s := client.Get("key3")

	log.Printf(s)
}
