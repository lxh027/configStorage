package logger

/**
to add prefix and suffix for raft instance's log
TODO upload logs to monitor
*/

import (
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"log"
	"strings"
)

type Logger struct {
	params []interface{}
	prefix string
}

func NewZapLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to initialize zap logger: %v", err)
	}
	grpc_zap.ReplaceGrpcLoggerV2(logger)
	return logger
}

func NewLogger(params []interface{}, prefix string) *Logger {
	logger := Logger{params: params, prefix: prefix}
	return &logger
}

func (logger *Logger) Printf(format string, v ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	v = append(logger.params, v...)
	log.Printf(logger.prefix+format, v...)
}

func (logger *Logger) Panicf(format string, v ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	v = append(logger.params, v...)
	log.Panicf(logger.prefix+format, v...)
}

func (logger *Logger) Fatalf(format string, v ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	v = append(logger.params, v...)
	log.Fatalf(logger.prefix+format, v...)
}
