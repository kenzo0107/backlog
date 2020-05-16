package backlog

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

// BacklogErrorResponse : backlog error response
type BacklogErrorResponse struct {
	Errors []BacklogError `json:"errors"`
}

type BacklogError struct {
	Message  string `json:"message"`
	Code     int    `json:"code"`
	MoreInfo string `json:"moreInfo"`
}

func (t BacklogError) Err() error {
	if strings.TrimSpace(t.Message) == "" {
		return nil
	}

	return errors.New(t.Message)
}

type statusCodeErrors struct {
	Errors []statusCodeError `json:"errors"`
}

// StatusCodeError represents an http response error.
// StatusCodeError represents an http response error.
// type httpStatusCode interface { HTTPStatusCode() int } to handle it.
type statusCodeError struct {
	Code   int
	Status string
}

func (t statusCodeError) Error() string {
	return fmt.Sprintf("backlog server error: %s", t.Status)
}

func (t statusCodeError) HTTPStatusCode() int {
	return t.Code
}

func (t statusCodeError) Retryable() bool {
	if t.Code >= 500 || t.Code == http.StatusTooManyRequests {
		return true
	}
	return false
}

func getResource(ctx context.Context, client httpClient, endpoint string, values url.Values, intf interface{}, d debug) error {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}

	req.URL.RawQuery = values.Encode()

	return doPost(ctx, client, req, newJSONParser(intf), d)
}

func doPost(ctx context.Context, client httpClient, req *http.Request, parser responseParser, d debug) error {
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = checkStatusCode(resp, d)
	if err != nil {
		return err
	}

	return parser(resp)
}

func checkStatusCode(resp *http.Response, d debug) error {
	if resp.StatusCode != http.StatusOK {
		logResponse(resp, d)
		return statusCodeError{Code: resp.StatusCode}
	}

	return nil
}

type responseParser func(*http.Response) error

func newJSONParser(dst interface{}) responseParser {
	return func(resp *http.Response) error {
		return json.NewDecoder(resp.Body).Decode(dst)
	}
}

func logResponse(resp *http.Response, d debug) error {
	if d.Debug() {
		text, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return err
		}
		d.Debugln(string(text))
	}

	return nil
}
