package raft

import "sync"

type Persister struct {
	mu        sync.Mutex
	raftState []byte
	snapshot  map[string]map[string]string
}

func NewPersister() *Persister {
	return &Persister{}
}

func (p *Persister) SaveRaftState(state []byte) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.raftState = state
}

func (p *Persister) SaveSnapshot(sp map[string]map[string]string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.snapshot = sp
}

func (p *Persister) ReadRaftState() []byte {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.raftState
}

func (p *Persister) SaveRaftStateAndSnapshot(state []byte, snapshot map[string]map[string]string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.snapshot = snapshot
	p.raftState = state
}

func (p *Persister) ReadSnapshot() map[string]map[string]string {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.snapshot
}
