package raft

import (
	"errors"
	"time"
)

const (
	UnVoted     = -1
	NonLogIndex = -1

	Success = true
	Fail    = false
)

type LogType uint8

const (
	KV LogType = iota
	LoadPrefix
	RemovePrefix
)

type State uint8

const (
	Shutdown State = iota
	Follower
	Candidate
	Leader
)

const (
	// StartUpTimeout waiting for other client to start
	StartUpTimeout = time.Second * 2
	// NewEntryTimeout 500ms for 4 times
	NewEntryTimeout = time.Millisecond * 50
	// LeaderRentTimeout when time is due, state should change from leader to follower
	LeaderRentTimeout = time.Second * 2
	// StatusLoggerTimeout report instance's status periodically
	StatusLoggerTimeout = time.Second * 3
	// ElectionTimeout timeout for election
	ElectionTimeout = time.Millisecond * 500
	// AppendEntriesTimeout timeout for append entries
	AppendEntriesTimeout = time.Millisecond * 300
	// AppendEntriesDuration  append entry duration
	AppendEntriesDuration = time.Millisecond * 200
	// LatencyTimeout wait a timeout to continue
	LatencyTimeout = time.Millisecond * 500
	// HeartBeatTimeout follower heart beat timeout
	HeartBeatTimeout = time.Millisecond * 1000
	// RandTimeout a random timeout as a suffix
	RandTimeout = time.Millisecond * 300
)

type StorageError error

var KeyNotFoundErr StorageError = errors.New("key not found")
var CopyDataErr StorageError = errors.New("copy data error")
