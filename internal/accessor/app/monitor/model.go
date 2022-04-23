package monitor

import "time"

type Monitor struct {
	ID            int    `json:"id"`
	RaftID        string `json:"raft_id"`
	PeerID        int    `json:"peer_id"`
	IsLeader      bool   `json:"is_leader"`
	Status        int    `json:"status"`
	ConfigVersion string `json:"config_version"`
	CurrentTerm   int32  `json:"current_term"`
	CurrentIndex  int32  `json:"current_index"`
	CommitIndex   int32  `json:"commit_index"`

	MemoryTotal     uint64    `json:"memory_total"`
	MemoryUsed      uint64    `json:"memory_used"`
	MemoryAvailable uint64    `json:"memory_available"`
	MemoryCur       uint64    `json:"memory_cur"`
	Time            time.Time `json:"time"`
}
