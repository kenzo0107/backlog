package backlog

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
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

// Errs : error
func (t ErrorResponse) Errs() error {
	s := []string{}
	for _, err := range t.Errors {
		s = append(s, fmt.Sprintf("code:%d message:%s moreInfo:%s", err.Code, err.Message, err.MoreInfo))
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

func getResource(ctx context.Context, client httpClient, endpoint string, values url.Values, intf interface{}, d debug) error {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}

	req.URL.RawQuery = values.Encode()

	return doPost(ctx, client, req, newJSONParser(intf), d)
}

func postForm(ctx context.Context, client httpClient, method, endpoint string, values url.Values, intf interface{}, d debug) error {
	reqBody := strings.NewReader(values.Encode())
	req, err := http.NewRequest(method, endpoint, reqBody)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return doPost(ctx, client, req, newJSONParser(intf), d)
}

func postLocalWithMultipartResponse(ctx context.Context, client httpClient, endpoint, fpath, fieldname string, values url.Values, intf interface{}, d debug) (err error) {
	fullpath, err := filepath.Abs(fpath)
	if err != nil {
		return err
	}

	file, err := os.Open(filepath.Clean(fullpath))
	if err != nil {
		return err
	}
	defer func() {
		if er := file.Close(); er != nil {
			err = er
		}
	}()

	if err := postWithMultipartResponse(ctx, client, endpoint, filepath.Base(fpath), fieldname, values, file, intf, d); err != nil {
		return err
	}

	return nil
}

func postWithMultipartResponse(ctx context.Context, client httpClient, endpoint, filename, fieldname string, values url.Values, r io.Reader, intf interface{}, d debug) (err error) {
	pipeReader, pipeWriter := io.Pipe()
	wr := multipart.NewWriter(pipeWriter)
	errc := make(chan error)
	go func() {
		defer func() {
			if er := pipeWriter.Close(); er != nil {
				errc <- er
			}
		}()
		ioWriter, er := wr.CreateFormFile(fieldname, filename)
		if er != nil {
			errc <- er
			return
		}
		_, errcp := io.Copy(ioWriter, r)
		if errcp != nil {
			errc <- errcp
			return
		}
		if errcl := wr.Close(); errcl != nil {
			errc <- errcl
			return
		}
	}()

	req, err := http.NewRequest("POST", endpoint, pipeReader)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", wr.FormDataContentType())
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

	select {
	case err = <-errc:
		return err
	default:
		return newJSONParser(intf)(resp)
	}
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
	if err := newJSONParser(errorResponse)(resp); err == nil {
		return errorResponse.Errs()
	}

	return statusCodeError{Code: resp.StatusCode, Status: resp.Status}
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

func downloadFile(ctx context.Context, client httpClient, apiKey, downloadURL string, writer io.Writer, d debug) (err error) {
	if downloadURL == "" {
		return fmt.Errorf("received empty download URL")
	}

	req, err := http.NewRequest("GET", downloadURL, &bytes.Buffer{})
	if err != nil {
		return err
	}

	values := url.Values{}
	values.Add("apiKey", apiKey)
	req.URL.RawQuery = values.Encode()

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

	_, err = io.Copy(writer, resp.Body)

	return err
}
