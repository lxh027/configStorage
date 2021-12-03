package raft

import "time"

const (
	// StartUpTimeout waiting for other client to start
	StartUpTimeout = time.Second * 10
	// NewEntryTimeout 500ms for 4 times
	NewEntryTimeout = time.Millisecond * 500
	// LeaderRentTimeout when time is due, state should change from leader to follower
	LeaderRentTimeout = time.Second * 5
	// StatusLoggerTimeout report instance's status periodically
	StatusLoggerTimeout = time.Second * 2
	// ElectionTimeout timeout for election
	ElectionTimeout = time.Millisecond * 200
	// AppendEntriesTimeout timeout for append entries
	AppendEntriesTimeout = time.Millisecond * 200
	// AppendEntriesDuration  append entry duration
	AppendEntriesDuration = time.Millisecond * 100
	// LatencyTimeout wait a timeout to continue
	LatencyTimeout = time.Millisecond * 100
	// HeartBeatTimeout follower heart beat timeout
	HeartBeatTimeout = time.Millisecond * 500
	// RandTimeout a random timeout as a suffix
	RandTimeout = time.Millisecond * 300
)
