package raft

import (
	"configStorage/tools/md5"
	"log"
	"testing"
	"time"
)

var rpcConfigs []Rpc
var clientConfig ClientConfig

const instanceNum = 3

func init() {
	rpcConfigs = []Rpc{
		{ID: 0, Host: "localhost", Port: "2000", CPort: "3000"},
		{ID: 1, Host: "localhost", Port: "2001", CPort: "3001"},
		{ID: 2, Host: "localhost", Port: "2002", CPort: "3002"},
	}

	clientConfig = ClientConfig{
		Size: instanceNum,
		Addresses: []string{
			"localhost:3000",
			"localhost:3001",
			"localhost:3002",
		},
	}
}

var rafts []*Raft

// TODO update test files
func makeRaft() {
	rafts = make([]*Raft, instanceNum)
	for i := 0; i < instanceNum; i++ {
		go func(i int) {
			var cfg Config
			cfg.RaftRpc = rpcConfigs[i]
			cfg.RaftPeers = rpcConfigs
			cfg.LogPrefix = "[raft id: %v]"
			//rafts[i] = NewRaftInstance(cfg)
			rafts[i].Start(md5.GetRandomMd5())
		}(i)
	}

	time.Sleep(StartUpTimeout)
}

func TestStorage(t *testing.T) {
	makeRaft()
	time.Sleep(5 * time.Second)
	client := NewRaftClient(clientConfig)
	client.Set("key1", "value")
	client.Set("key2", "value2")
	client.Set("key4", "value4")
	client.Set("key5", "value5")
	time.Sleep(NewEntryTimeout)

	client.Get("key1")
	client.Get("key2")
	client.Get("key3")
	client.Get("key4")
	client.Get("key5")

	client.Set("key5", "v  kajsgalue2")
	time.Sleep(NewEntryTimeout)
	client.Get("key5")

	client.Del("key1")
	client.Del("key5")
	time.Sleep(NewEntryTimeout)

	client.Get("key5")

	for i := 0; i < instanceNum; i++ {
		rafts[i].Kill()
	}
}

func TestLeaderElection(t *testing.T) {
	makeRaft()
	// TestLeaderElection
	time.Sleep(time.Second * 2)
	ld := 0
	for i := 0; i < instanceNum; i++ {
		if rafts[i].state == Leader {
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
		if raft.state == Leader {
			ld = i
			raft.Kill()
			break
		}
	}
	// wait for reelection
	time.Sleep(2 * time.Second)

	cnt := 0
	for i := 0; i < instanceNum; i++ {
		if rafts[i].state == Leader {
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

	time.Sleep(5 * time.Second)

	for _, raft := range rafts {
		if raft.currentIndex != 10 {
			log.Fatalf("error appending logs")
		}
	}
	for i := 0; i < instanceNum; i++ {
		rafts[i].Kill()
	}
}

func (rf *Raft) appendLog(entry []byte) bool {
	rf.mu.Lock()
	defer rf.mu.Unlock()

	if rf.state != Leader {
		return false
	}

	logIndex := rf.currentIndex
	rf.currentIndex = rf.currentIndex + 1
	// new log
	l := Log{
		Entry:  entry,
		Term:   rf.currentTerm,
		Index:  logIndex,
		Status: true,
	}
	rf.logs = append(rf.logs, l)
	rf.logger.Printf("new log entry at term %d index %d", l.Term, l.Index)
	return true

}
