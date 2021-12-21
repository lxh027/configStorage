package raft_peer

import (
	"configStorage/pkg/config"
	"configStorage/pkg/raft"
	"path"
)

func StartRaftPeer(env string, id int) {
	basePath := path.Join("./config", env)

	raftPath := path.Join(basePath, "raft.yml")
	raftCfg := config.NewRaftRpcConfig(raftPath)

	flag := false
	for _, cfg := range raftCfg.RaftPeers {
		if int(cfg.ID) == id {
			if flag {
				panic("raft.yml config error: peers' id duplicated")
			}
			raftCfg.RaftRpc = cfg
			flag = true
			break
		}
	}
	if !flag {
		panic("param error: id is not existed")
	}

	raftCfg.RaftRpc = raftCfg.RaftPeers[id]
	peer := raft.NewRaftInstance(raftCfg)

	peer.Start()
}
