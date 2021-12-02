package logger

/**
to add prefix and suffix for raft instance's log
TODO upload logs to monitor
*/

import (
	"log"
	"storage/config"
)

const (
	prefix = "[raft id %d] "
	suffix = "\n"
)

var raftID int32

func init() {
	raftID = config.GetRpcConfig().RaftRpc.ID
}

func Printf(format string, v ...interface{}) {
	log.Printf(prefix+format+suffix, raftID, v)
}

func Panicf(format string, v ...interface{}) {
	log.Panicf(prefix+format+suffix, raftID, v)
}

func Fatalf(format string, v ...interface{}) {
	log.Fatalf(prefix+format+suffix, raftID, v)
}
