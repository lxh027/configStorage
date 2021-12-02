package raft

import "time"

const (
	// NewEntryTimeout 500ms for 4 times
	NewEntryTimeout = time.Second / 2
)
