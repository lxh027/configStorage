package raft

import (
	"log"
	"storage/config"
	constants "storage/constants/raft"
	"testing"
	"time"
)

var rpcConfigs []config.Rpc

const instanceNum = 3

func init() {
	rpcConfigs = []config.Rpc{
		{ID: 0, Host: "localhost", Port: "2000"},
		{ID: 1, Host: "localhost", Port: "2001"},
		{ID: 2, Host: "localhost", Port: "2002"},
	}
}

var rafts []*Raft

func makeRaft() {
	rafts = make([]*Raft, 3)
	for i := 0; i < instanceNum; i++ {
		go func(i int) {
			var cfg config.RpcConfig
			cfg.RaftRpc = rpcConfigs[i]
			cfg.RaftPeers = rpcConfigs
			rafts[i] = NewRaftInstance(cfg)
		}(i)
	}

	time.Sleep(constants.StartUpTimeout)
}

func TestLeaderElection(t *testing.T) {
	makeRaft()
	// TestLeaderElection
	time.Sleep(time.Second * 2)
	ld := 0
	for i := 0; i < instanceNum; i++ {
		if rafts[i].state == constants.Leader {
			ld++
		}
	}
	if ld == 0 {
		log.Fatalf("can't select leader")
	}
	if ld > 1 {
		log.Fatalf("there's over one leader")
	}
	for i := 0; i < instanceNum; i++ {
		rafts[i].Kill()
	}
}

func TestAppendEntries(t *testing.T) {
	makeRaft()
	time.Sleep(2 * time.Second)
	for i := 0; i < 10; i++ {
		for _, raft := range rafts {
			raft.appendLog([]byte("testLog"))
		}
	}

	time.Sleep(2 * time.Second)
	for _, raft := range rafts {
		if raft.currentIndex != 10 {
			log.Fatalf("error appending logs")
		}
	}
	for i := 0; i < instanceNum; i++ {
		rafts[i].Kill()
	}
}

func TestReelection(t *testing.T) {
	makeRaft()
	time.Sleep(2 * time.Second)

	for i := 0; i < 5; i++ {
		for _, raft := range rafts {
			raft.appendLog([]byte("testLog"))
		}
	}

	time.Sleep(2 * time.Second)

	// kill leader
	var ld int
	for i, raft := range rafts {
		if raft.state == constants.Leader {
			ld = i
			raft.Kill()
			break
		}
	}
	// wait for reelection
	time.Sleep(2 * time.Second)

	cnt := 0
	for i := 0; i < instanceNum; i++ {
		if rafts[i].state == constants.Leader {
			cnt++
		}
	}
	if cnt == 0 {
		log.Fatalf("can't select leader")
	}
	if cnt > 1 {
		log.Fatalf("there's over one leader")
	}

	// append logs
	for i := 0; i < 5; i++ {
		for _, raft := range rafts {
			raft.appendLog([]byte("testLog"))
		}
	}

	time.Sleep(2 * time.Second)

	// restart old leader
	rafts[ld].Restart(rpcConfigs[ld].Host, rpcConfigs[ld].Port)

	time.Sleep(2 * time.Second)

	for _, raft := range rafts {
		if raft.currentIndex != 10 {
			log.Fatalf("error appending logs")
		}
	}
	for i := 0; i < instanceNum; i++ {
		rafts[i].Kill()
	}
}

func (raft *Raft) appendLog(entry []byte) bool {
	raft.mu.Lock()
	defer raft.mu.Unlock()

	if raft.state != constants.Leader {
		return false
	}

	logIndex := raft.currentIndex
	raft.currentIndex = raft.currentIndex + 1
	// new log
	log := Log{
		Entry:  entry,
		Term:   raft.currentTerm,
		Index:  logIndex,
		Status: true,
	}
	raft.logs = append(raft.logs, log)
	raft.logger.Printf("new log entry at term %d index %d", log.Term, log.Index)
	return true

}
