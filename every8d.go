package every8d

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	libraryVersion   = "0.0.1"
	defaultBaseURL   = "https://oms.every8d.com/"
	defaultUserAgent = "go-every8d/" + libraryVersion
)

// A Client manages communication with the EVERY8D API.
type Client struct {
	client   *http.Client
	username string
	password string

	BaseURL   *url.URL
	UserAgent string
}

// NewClient returns a new EVERY8D API client.
func NewClient(username, password string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	return &Client{
		client:    httpClient,
		username:  username,
		password:  password,
		UserAgent: defaultUserAgent,
		BaseURL:   baseURL,
	}
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Response is a EVERY8D API response.
type Response struct {
	*http.Response
}

// newResponse creates a new Response for the provided http.Response.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}

	// TODO body

	return response
}

// Do sends an API request and returns the API response.
//
// The provided ctx must be non-nil. If it is canceled or time out, ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request) (*Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}
	defer resp.Body.Close()

	response := newResponse(resp)

	if err := CheckResponse(response); err != nil {
		return response, err
	}

	return response, err
}

// CheckResponse checks the API response for errors.
func CheckResponse(r *Response) error {
	c := r.StatusCode
	if 200 <= c && c <= 299 {
		return nil
	}

	// TODO parse body

	// EVERY8D API always return status code 200
	return fmt.Errorf("unexpected status code: %d", c)
}
