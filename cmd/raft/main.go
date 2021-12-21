package main

import (
	"configStorage/internal/raft_peer"
	"flag"
)

func main() {
	var env string
	var id int

	flag.StringVar(&env, "env", "dev", "配置环境")
	flag.IntVar(&id, "id", 0, "RAFT PEER ID")
	flag.Parse()

	raft_peer.StartRaftPeer(env, id)
}
