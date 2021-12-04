package raft

import "time"

const (
	// StartUpTimeout waiting for other client to start
	StartUpTimeout = time.Second * 2
	// NewEntryTimeout 500ms for 4 times
	NewEntryTimeout = time.Millisecond * 500
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
