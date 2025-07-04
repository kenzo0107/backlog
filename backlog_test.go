package backlog

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
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

func TestClient_Debug(t *testing.T) {
	client := New("test-token", "https://example.com", OptionDebug(true))
	assert.True(t, client.Debug())

	client = New("test-token", "https://example.com", OptionDebug(false))
	assert.False(t, client.Debug())
}

func TestClient_Debugf(t *testing.T) {
	buf := bytes.NewBufferString("")
	client := New("test-token", "https://example.com",
		OptionDebug(true),
		OptionLog(log.New(buf, "", 0)))

	client.Debugf("test %s", "message")
	assert.Contains(t, buf.String(), "test message")
}

func TestClient_Debugln(t *testing.T) {
	buf := bytes.NewBufferString("")
	client := New("test-token", "https://example.com",
		OptionDebug(true),
		OptionLog(log.New(buf, "", 0)))

	client.Debugln("test", "message")
	assert.Contains(t, buf.String(), "test message")
}

func TestNewRequestWithTrailingSlashError(t *testing.T) {
	client := New("test-token", "https://example.com/api/v2/")

	_, err := client.NewRequest("GET", "/test", nil)
	if err == nil {
		t.Error("Expected error for baseURL with trailing slash")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}
}

func TestNewRequestWithInvalidURL(t *testing.T) {
	client := New("test-token", "https://example.com/api/v2")

	// 無効な文字を含むURLを使用してbaseURL.Parseを失敗させる
	_, err := client.NewRequest("GET", string([]byte{0x00, 0x01, 0x02}), nil)
	if err == nil {
		t.Error("Expected error for invalid URL")
	}
}

func TestNewRequestWithJSONEncodeError(t *testing.T) {
	client := New("test-token", "https://example.com/api/v2")

	// JSONエンコードできないボディ（チャンネル型）を使用
	invalidBody := make(chan int)
	_, err := client.NewRequest("POST", "/test", invalidBody)
	if err == nil {
		t.Error("Expected error for invalid JSON body")
	}
}

func TestNewRequestWithInvalidHTTPMethod(t *testing.T) {
	client := New("test-token", "https://example.com/api/v2")

	// 無効なHTTPメソッドを使用してhttp.NewRequestを失敗させる
	_, err := client.NewRequest("INVALID\nMETHOD", "/test", nil)
	if err == nil {
		t.Error("Expected error for invalid HTTP method")
	}
}
