package backlog

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

// httpClient defines the minimal interface needed for an http.Client to be implemented.
type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Client : backlog client
type Client struct {
	apiKey     string
	endpoint   string
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
	s := &Client{
		apiKey:     apiKey,
		endpoint:   endpoint,
		httpclient: &http.Client{},
		log:        log.New(os.Stderr, "kenzo0107/backlog", log.LstdFlags|log.Lshortfile),
	}

	for _, opt := range options {
		opt(s)
	}

	return s
}

// Debugf print a formatted debug line.
func (api *Client) Debugf(format string, v ...interface{}) {
	if api.debug {
		if err := api.log.Output(2, fmt.Sprintf(format, v...)); err != nil {
			api.Debugln(err)
		}
	}
}

// Debugln print a debug line.
func (api *Client) Debugln(v ...interface{}) {
	if api.debug {
		if err := api.log.Output(2, fmt.Sprintln(v...)); err != nil {
			api.Debugln(err)
		}
	}
}

// Debug returns if debug is enabled.
func (api *Client) Debug() bool {
	return api.debug
}

// get method
func (api *Client) getMethod(ctx context.Context, path string, values url.Values, intf interface{}) error {
	values.Add("apiKey", api.apiKey)
	return getResource(ctx, api.httpclient, api.endpoint+path, values, intf, api)
}

// post method
func (api *Client) postMethod(ctx context.Context, path string, values url.Values, intf interface{}) error {
	return postForm(ctx, api.httpclient, "POST", api.endpoint+path+"?apiKey="+api.apiKey, values, intf, api)
}

// patch method
func (api *Client) patchMethod(ctx context.Context, path string, values url.Values, intf interface{}) error {
	return postForm(ctx, api.httpclient, "PATCH", api.endpoint+path+"?apiKey="+api.apiKey, values, intf, api)
}

// delete method
func (api *Client) deleteMethod(ctx context.Context, path string, values url.Values, intf interface{}) error {
	return postForm(ctx, api.httpclient, "DELETE", api.endpoint+path+"?apiKey="+api.apiKey, values, intf, api)
}
