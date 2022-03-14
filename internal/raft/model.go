package raft

import (
	"configStorage/api/raftrpc"
	"configStorage/pkg/logger"
	"google.golang.org/grpc"
	"sync"
)

// Raft object to implement raft state
type Raft struct {
	// rpc server implementation
	raftrpc.UnimplementedRaftServer
	raftrpc.UnimplementedStateServer

	// logger
	logger *logger.Logger

	// persister
	persister *Persister

	// configs
	cfg Config

	// storage
	storage Storage

	// rpc server
	rpcServer *grpc.Server
	// resServer
	stateServer *grpc.Server
	// mutex
	mu sync.Mutex

	// raft id for the instance
	id int32

	leaderID int32

	// peers is raft peer instance's host and port
	peers []Peer

	cfgVersion string

	// params to indicate instance status
	// currentTerm is term id of the instance
	// commitIndex is the Index of the last commit Entry
	// lastApplied is the Index of thr last Entry applied to the state machine
	currentTerm  int32
	currentIndex int32
	commitIndex  int32
	lastApplied  int32

	// the last vote
	votedFor int32
	// heartbeat receive
	heartbeat bool

	// state for the instance
	state       State
	stateChange chan State

	logs []Log
	// leaderState when the instance become leader, to record followers' state to maintain leadership
	leaderState struct {
		nextIndex  map[int32]int32
		matchIndex map[int32]int32
	}
}

// Log is the command sent by client
// including leader's Term and command's Index
// Entry is the command interface
// Status if the log is no longer useful
type Log struct {
	Entry  []byte
	Term   int32
	Index  int32
	Status bool
}

type Peer struct {
	client  raftrpc.RaftClient
	version string
}
