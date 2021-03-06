package every8d

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
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

	// Base URL for API requests, should always be specified with a trailing slash.
	BaseURL *url.URL

	// User agent used when communicating with the EVERY8D API.
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
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
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

// NewFormRequest creates an API POST request.
func (c *Client) NewFormRequest(urlStr string, form url.Values) (*http.Request, error) {
	form.Set("UID", c.username)
	form.Set("PWD", c.password)

	req, err := c.NewRequest(http.MethodPost, urlStr, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

// Parser parses the response body into a object v.
type Parser func(body io.Reader, v interface{}) error

// Do sends an API request and returns the API response.
//
// The provided ctx must be non-nil. If it is canceled or time out, ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, fn Parser, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// If the error type is *url.Error, sanitize its URL before returning.
		if e, ok := err.(*url.Error); ok {
			if u, err := url.Parse(e.URL); err == nil {
				e.URL = sanitizeURL(u).String()
				return nil, e
			}
		}

		return nil, err
	}
	defer resp.Body.Close()

	if err := CheckResponse(resp); err != nil {
		return resp, err
	}

	if err := fn(resp.Body, v); err != nil {
		return resp, err
	}

	return resp, nil
}

// sanitizeURL redacts the PWD parameter from the URL which may be exposed to the user.
func sanitizeURL(uri *url.URL) *url.URL {
	if uri == nil {
		return nil
	}
	params := uri.Query()
	if len(params.Get("PWD")) > 0 {
		params.Set("PWD", "REDACTED")
		uri.RawQuery = params.Encode()
	}
	return uri
}

// ErrorResponse reports error caused by an API request.
type ErrorResponse struct {
	Response  *http.Response
	ErrorCode StatusCode
	Message   string
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %d %v",
		r.Response.Request.Method,
		sanitizeURL(r.Response.Request.URL),
		r.Response.StatusCode,
		r.ErrorCode,
		r.Message)
}

// CheckResponse checks the API response for errors.
func CheckResponse(r *http.Response) error {
	if r.StatusCode == 200 {
		reader := bufio.NewReader(r.Body)
		firstByte, err := reader.ReadByte()
		if err != nil {
			return err
		}

		reader.UnreadByte()

		if string(firstByte) == "-" {
			errorString, _ := reader.ReadString('\n')
			if matched, _ := regexp.MatchString("-\\d+,.+", errorString); matched == false {
				return fmt.Errorf("invalid message format")
			}
			errors := strings.Split(errorString, ",")
			errorCode, _ := strconv.Atoi(errors[0])

			return &ErrorResponse{
				Response:  r,
				Message:   strings.TrimSpace(errors[1]),
				ErrorCode: StatusCode(errorCode),
			}
		}

		// reset the response body to the original unread state
		r.Body = ioutil.NopCloser(reader)

		return nil
	}

	// EVERY8D API always return status code 200
	return fmt.Errorf("unexpected status code: %d", r.StatusCode)
}
