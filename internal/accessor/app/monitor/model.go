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

	RevoteTime int64 `json:"revote_time"`
	CommitTime int64 `json:"commit_time"`

	MemoryTotal     uint64    `json:"memory_total"`
	MemoryUsed      uint64    `json:"memory_used"`
	MemoryAvailable uint64    `json:"memory_available"`
	MemoryCur       uint64    `json:"memory_cur"`
	Time            time.Time `json:"time"`
}

type Peer struct {
	RaftID   string `json:"raft_id,omitempty"`
	PeerID   int    `json:"peer_id"`
	IsLeader bool   `json:"is_leader"`
}

type Data struct {
	Basic  BasicData  `json:"basic"`
	Memory MemoryData `json:"memory"`
	Raft   []RaftData `json:"raft"`
	Time   TimeData   `json:"time"`
}

type BasicData struct {
	RaftID        string `json:"raft_id"`
	PeerID        int    `json:"peer_id"`
	IsLeader      bool   `json:"is_leader"`
	Status        int    `json:"status"`
	ConfigVersion string `json:"config_version"`
}

type MemoryData struct {
	MemoryTotal     uint64 `json:"memory_total"`
	MemoryUsed      uint64 `json:"memory_used"`
	MemoryAvailable uint64 `json:"memory_available"`
	MemoryCur       uint64 `json:"memory_cur"`
}

type TimeData struct {
	RevoteTime int64 `json:"revote_time"`
	CommitTime int64 `json:"commit_time"`
}

type RaftData struct {
	Item string    `json:"item"`
	Data [][]int64 `json:"data"`
}
