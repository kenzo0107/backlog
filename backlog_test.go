package backlog

import (
	"log"
	"net/http/httptest"
	"sync"

	"github.com/pkg/errors"
)

var (
	serverAddr string
	once       sync.Once
)

var (
	ErrIncorrectResponse = errors.New("Response is incorrect")
)

func startServer() {
	server := httptest.NewServer(nil)
	serverAddr = server.Listener.Addr().String()
	log.Print("Test WebSocket server listening on ", serverAddr)
}
