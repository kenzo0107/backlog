package backlog

import (
	"fmt"
	"log"
)

var logFatal = log.Fatal

type logger interface {
	Output(int, string) error
}

// ilogger represents the internal logging api we use.
type ilogger interface {
	logger
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})
}

type debug interface {
	Debug() bool

	// Debugf print a formatted debug line.
	Debugf(format string, v ...interface{})
	// Debugln print a debug line.
	Debugln(v ...interface{})
}

// internalLog implements the additional methods used by our internal logging.
type internalLog struct {
	logger
}

// Println replicates the behaviour of the standard logger.
func (t internalLog) Println(v ...interface{}) {
	if err := t.Output(2, fmt.Sprintln(v...)); err != nil {
		logFatal(err)
	}
}

// Printf replicates the behaviour of the standard logger.
func (t internalLog) Printf(format string, v ...interface{}) {
	if err := t.Output(2, fmt.Sprintf(format, v...)); err != nil {
		logFatal(err)
	}
}

// Print replicates the behaviour of the standard logger.
func (t internalLog) Print(v ...interface{}) {
	if err := t.Output(2, fmt.Sprint(v...)); err != nil {
		logFatal(err)
	}
}
