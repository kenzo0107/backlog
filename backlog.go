package backlog

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// httpClient defines the minimal interface needed for an http.Client to be implemented.
type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Client : backlog client
type Client struct {
	apiKey     string
	endpoint   string
	baseURL    *url.URL
	debug      bool
	log        ilogger
	httpclient httpClient
}

// Option defines an option for a Client
type Option func(*Client)

// OptionHTTPClient - provide a custom http client to the backlog client.
func OptionHTTPClient(client httpClient) func(*Client) {
	return func(c *Client) {
		c.httpclient = client
	}
}

// OptionDebug enable debugging for the client
func OptionDebug(b bool) func(*Client) {
	return func(c *Client) {
		c.debug = b
	}
}

// OptionLog set logging for client.
func OptionLog(l logger) func(*Client) {
	return func(c *Client) {
		c.log = internalLog{logger: l}
	}
}

// New builds a backlog client from the provided token, baseURL and options
func New(apiKey, endpoint string, options ...Option) *Client {
	baseURL, _ := url.Parse(endpoint)
	s := &Client{
		apiKey:     apiKey,
		endpoint:   endpoint,
		baseURL:    baseURL,
		httpclient: &http.Client{},
		log:        log.New(os.Stderr, "kenzo0107/backlog", log.LstdFlags|log.Lshortfile),
	}

	for _, opt := range options {
		opt(s)
	}

	return s
}

// Debugf print a formatted debug line.
func (c *Client) Debugf(format string, v ...interface{}) {
	if c.debug {
		if err := c.log.Output(2, fmt.Sprintf(format, v...)); err != nil {
			c.Debugln(err)
		}
	}
}

// Debugln print a debug line.
func (c *Client) Debugln(v ...interface{}) {
	if c.debug {
		if err := c.log.Output(2, fmt.Sprintln(v...)); err != nil {
			c.Debugln(err)
		}
	}
}

// Debug returns if debug is enabled.
func (c *Client) Debug() bool {
	return c.debug
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if strings.HasSuffix(c.baseURL.Path, "/") {
		return nil, fmt.Errorf("baseURL must not have a trailing slash, but %q does", c.baseURL)
	}

	u, err := c.baseURL.Parse(c.baseURL.Path + urlStr)
	if err != nil {
		return nil, err
	}

	// add 'apiKey' to the url path
	q := u.Query()
	q.Set("apiKey", c.apiKey)
	u.RawQuery = q.Encode()

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		if er := enc.Encode(body); er != nil {
			return nil, er
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

// UploadMultipartFile uploads multipart file
func (c *Client) UploadMultipartFile(ctx context.Context, method, urlStr, fpath, field string, v interface{}) (err error) {
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

	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)
	errc := make(chan error)
	go func() {
		defer func() {
			if er := pw.Close(); er != nil {
				errc <- er
			}
		}()
		ioWriter, er := mw.CreateFormFile(field, filepath.Base(fpath))
		if er != nil {
			errc <- er
			return
		}
		_, errcp := io.Copy(ioWriter, file)
		if errcp != nil {
			errc <- errcp
			return
		}
		if errcl := mw.Close(); errcl != nil {
			errc <- errcl
			return
		}
	}()

	if strings.HasSuffix(c.baseURL.Path, "/") {
		return fmt.Errorf("baseURL must not have a trailing slash, but %q does", c.baseURL)
	}

	u, err := c.baseURL.Parse(c.baseURL.Path + urlStr)
	if err != nil {
		return err
	}

	// add 'apiKey' to the url path
	q := u.Query()
	q.Set("apiKey", c.apiKey)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(method, u.String(), pr)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", mw.FormDataContentType())
	if err := c.Do(ctx, req, &v); err != nil {
		return err
	}
	return nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it. If rate limit is exceeded and reset time is in the future,
// Do returns *RateLimitError immediately without making a network API call.
//
// The provided ctx must be non-nil, if it is nil an error is returned. If it is canceled or times out,
// ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) error {
	if ctx == nil {
		return errors.New("context must be non-nil")
	}

	req = req.WithContext(ctx)

	resp, err := c.httpclient.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		return err
	}
	defer func() {
		if er := resp.Body.Close(); er != nil {
			err = er
		}
	}()

	err = checkStatusCode(resp, c)
	if err != nil {
		return err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			if _, er := io.Copy(w, resp.Body); er != nil {
				return er
			}
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return err
}

// AddOptions adds the parameters in opt as URL query parameters to s. opt
// must be a struct whose fields may contain "url" tags.
func (c *Client) AddOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()

	return u.String(), nil
}

type p struct {
	Count int
}
