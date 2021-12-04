package logger

/**
to add prefix and suffix for raft instance's log
TODO upload logs to monitor
*/

import (
	"log"
)

const (
	prefix = "[raft id %d] "
	suffix = "\n"
)

type Logger struct {
	raftID int32
}

func NewLogger(id int32) *Logger {
	logger := Logger{raftID: id}
	return &logger
}

func (logger *Logger) Printf(format string, v ...interface{}) {
	v = append([]interface{}{logger.raftID}, v...)
	log.Printf(prefix+format+suffix, v...)
}

func (logger *Logger) Panicf(format string, v ...interface{}) {
	v = append([]interface{}{logger.raftID}, v...)
	log.Panicf(prefix+format+suffix, v...)
}

func (logger *Logger) Fatalf(format string, v ...interface{}) {
	v = append([]interface{}{logger.raftID}, v...)
	log.Fatalf(prefix+format+suffix, v...)
}
