package backlog

import (
	"bytes"
	"log"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestLogging(t *testing.T) {
	buf := bytes.NewBufferString("")
	logger := internalLog{logger: log.New(buf, "", 0|log.Lshortfile)}
	logger.Println("test line 123")
	assert.Equal(t, "logger_test.go:15: test line 123\n", buf.String())
	buf.Truncate(0)
	logger.Print("test line 123")
	assert.Equal(t, "logger_test.go:18: test line 123\n", buf.String())
	buf.Truncate(0)
	logger.Printf("test line 123\n")
	assert.Equal(t, "logger_test.go:21: test line 123\n", buf.String())
	buf.Truncate(0)
	if err := logger.Output(1, "test line 123\n"); err != nil {
		log.Println(err)
	}
	assert.Equal(t, "logger_test.go:24: test line 123\n", buf.String())
	buf.Truncate(0)
}

type mockLogger struct {
	shouldError bool
}

func (m *mockLogger) Output(_ int, _ string) error {
	if m.shouldError {
		return errors.New("mock error")
	}
	return nil
}

func TestInternalLog_PrintlnError(t *testing.T) {
	originalLogFatal := logFatal
	defer func() {
		logFatal = originalLogFatal
	}()

	var fatalCalled bool
	logFatal = func(_ ...interface{}) {
		fatalCalled = true
	}

	logger := internalLog{logger: &mockLogger{shouldError: true}}
	logger.Println("test")
	assert.True(t, fatalCalled)
}
