package backlog

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// ErrorResponse is backlog error response
type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

// Error is backlog error
type Error struct {
	Message  string `json:"message"`
	Code     int    `json:"code"`
	MoreInfo string `json:"moreInfo"`
}

// Err : error
func (t Error) Err() error {
	if strings.TrimSpace(t.Message) == "" {
		return nil
	}

	return errors.New(t.Message)
}

// Errs : error
func (t ErrorResponse) Errs() error {
	s := []string{}
	for _, err := range t.Errors {
		s = append(s, err.Message)
	}

	if len(s) == 0 {
		return nil
	}

	return errors.New(strings.Join(s, ", "))
}

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

func doPost(ctx context.Context, client httpClient, req *http.Request, parser responseParser, d debug) (err error) {
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if er := resp.Body.Close(); er != nil {
			err = er
		}
	}()

	err = checkStatusCode(resp, d)
	if err != nil {
		return err
	}

	return parser(resp)
}

func checkStatusCode(resp *http.Response, d debug) error {
	// return no error if response returns status code 2xx
	if resp.StatusCode/100 == 2 {
		return nil
	}

	if err := logResponse(resp, d); err != nil {
		return err
	}

	errorResponse := &ErrorResponse{}
	if err := newJSONParser(errorResponse)(resp); err != nil {
		return err
	}
	return errorResponse.Errs()
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

func projIDOrKey(projIDOrKey interface{}) (string, error) {
	var idOrKey string
	switch t := projIDOrKey.(type) {
	case int:
		idOrKey = strconv.Itoa(t)
	case string:
		idOrKey = t
	default:
		return idOrKey, fmt.Errorf("projectIDOrKey is int or string. you specify %t", t)
	}
	return idOrKey, nil
}
