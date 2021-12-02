package raft

import "time"

const (
	// NewEntryTimeout 500ms for 4 times
	NewEntryTimeout     = time.Second / 2
	// LeaderRentTimeout when time is due, state should change from leader to follower
	LeaderRentTimeout   = time.Second * 5
	// StatusLoggerTimeout report instance's status periodically
	StatusLoggerTimeout = time.Second * 2
)
