package raft

import "sync"

type Persister struct {
	mu        sync.Mutex
	raftState []byte
	snapshot  []byte
}

func NewPersister() *Persister {
	return &Persister{}
}

func (p *Persister) SaveRaftState(state []byte) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.raftState = state
}

func (p *Persister) SaveSnapshot(sp []byte) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.snapshot = sp
}

func (p *Persister) ReadRaftState() []byte {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.raftState
}

func (p *Persister) SaveRaftStateAndSnapshot(state []byte, snapshot []byte) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.snapshot = snapshot
	p.raftState = state
}

func (p *Persister) ReadSnapshot() []byte {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.snapshot
}
