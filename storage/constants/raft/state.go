package raft

type State uint8

const (
	Shutdown 	State = iota
	Follower
	Candidate
	Leader
)
