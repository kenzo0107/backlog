package backlog

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/pkg/errors"
)

const baseURLPath string = "/api/v2"

var (
	ErrIncorrectResponse = errors.New("response is incorrect")
)

// setup sets up a test HTTP server along with a backlog.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	return setupWithPath("")
}

// setupWithPath sets up a test HTTP server along with a backlog.Client with the path.
func setupWithPath(path string) (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	apiHandler := http.NewServeMux()
	apiHandler.Handle(baseURLPath+"/", http.StripPrefix(baseURLPath, mux))
	apiHandler.HandleFunc("/", func(_ http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(os.Stderr, "\t"+req.URL.String())
		fmt.Fprintln(os.Stderr)
	})

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(apiHandler)

	// client is the backlog client being tested and is
	// configured to use test server.
	client = New(
		"test-token",
		server.URL+path,
		OptionHTTPClient(&http.Client{}),
		OptionDebug(false),
		OptionLog(log.New(os.Stderr, "kenzo0107/backlog", log.LstdFlags|log.Lshortfile)),
	)

	return client, mux, server.URL, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func TestBaseURLWithPath(t *testing.T) {
	client, _, _, teardown := setupWithPath("/sub")
	defer teardown()

	if _, err := client.GetSpace(); err != nil {
		t.Fatal("Unexpected error", err)
	}
}
