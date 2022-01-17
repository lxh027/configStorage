package global

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type Logger interface {
	DebugPrintf(string, ...interface{})
	Printf(string, ...interface{})
}

type logger struct{}

func NewLogger() Logger {
	return &logger{}
}

func (l *logger) DebugPrintf(format string, values ...interface{}) {
	if gin.IsDebugging() {
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		_, _ = fmt.Fprintf(gin.DefaultWriter, "[GIN-debug] "+format, values...)
	}
}

func (l *logger) Printf(format string, values ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	_, _ = fmt.Fprintf(gin.DefaultWriter, "[GIN-debug] "+format, values...)
}
