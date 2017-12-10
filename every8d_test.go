package every8d

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the Apple Music client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

// setup sets up a test HTTP server along with a every9d.Client that is configured to talk to that test server.
func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// EVERY8D client configured to use test server
	client = NewClient("username", "password", nil)
	u, _ := url.Parse(server.URL)
	client.BaseURL = u
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testURLParseError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func TestNewClient(t *testing.T) {
	c := NewClient("username", "password", nil)

	if got, want := c.BaseURL.String(), defaultBaseURL; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}
	if got, want := c.UserAgent, defaultUserAgent; got != want {
		t.Errorf("NewClient UserAgent is %v, want %v", got, want)
	}
}

func TestClient_NewRequest(t *testing.T) {
	c := NewClient("username", "password", nil)

	inURL, outURL := "/foo", defaultBaseURL+"foo"
	inBody, outBody := "Hello, 世界", "Hello, 世界"
	req, _ := c.NewRequest("GET", inURL, strings.NewReader(inBody))

	// test that relative URL was expanded
	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL is %v, want %v", inURL, got, want)
	}

	// test that body
	body, _ := ioutil.ReadAll(req.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("NewRequest(%q) Body is %v, want %v", inBody, got, want)
	}

	// test that default user-agent is attached to the request
	if got, want := req.Header.Get("User-Agent"), c.UserAgent; got != want {
		t.Errorf("NewRequest() User-Agent is %v, want %v", got, want)
	}
}

func TestClient_NewRequest_badURL(t *testing.T) {
	c := NewClient("", "", nil)
	_, err := c.NewRequest("GET", ":", nil)
	testURLParseError(t, err)
}

func TestClient_Do(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}
		fmt.Fprint(w, "Hello, 世界")
	})

	req, _ := client.NewRequest("GET", "/", nil)
	fn := func(r io.Reader, v interface{}) error {
		bs, _ := ioutil.ReadAll(r)
		*v.(*string) = string(bs)
		return nil
	}
	body := new(string)
	client.Do(context.Background(), req, fn, body)

	want := "Hello, 世界"
	if !reflect.DeepEqual(*body, want) {
		t.Errorf("Response body = %v, want %v", *body, want)
	}
}

// Test that an error caused by the internal http client's Do() function does not leak the client PWD.
func TestDo_sanitizeURL(t *testing.T) {
	client := NewClient("username", "password", nil)
	client.BaseURL = &url.URL{Scheme: "http", Host: "127.0.0.1:0", Path: "/"} // Use port 0 on purpose to trigger a dial TCP error, expect to get "dial tcp 127.0.0.1:0: connect: can't assign requested address".
	req, err := client.NewRequest("GET", ".", nil)
	if err != nil {
		t.Fatalf("NewRequest returned unexpected error: %v", err)
	}
	_, err = client.Do(context.Background(), req, nil, nil)
	if err == nil {
		t.Fatal("Expected error to be returned.")
	}
	if strings.Contains(err.Error(), "PWD=password") {
		t.Errorf("Do error contains password, should be redacted:\n%q", err)
	}
}

func TestClient_Do_noContent(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	_, err := client.Do(context.Background(), req, nil, nil)
	if err == nil {
		t.Errorf("Expected error response.") // read EOF
	}
}

func TestClient_Do_parseError(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Hello, 世界")
	})

	got := "parse error"

	req, _ := client.NewRequest("GET", "/", nil)
	fn := func(r io.Reader, v interface{}) error {
		return errors.New(got)
	}
	_, err := client.Do(context.Background(), req, fn, new(string))
	if err == nil {
		t.Errorf("Expected error response.")
	}
	if want := err.Error(); got != want {
		t.Errorf("Error = %v, want %v", got, want)
	}
}

func TestSanitizeURL(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"/?a=b", "/?a=b"},
		{"/?PWD=password&a=b", "/?PWD=REDACTED&a=b"},
		{"/?PWD=password&a=b&client_id=id", "/?PWD=REDACTED&a=b&client_id=id"},
	}

	for _, tt := range tests {
		inURL, _ := url.Parse(tt.in)
		want, _ := url.Parse(tt.want)

		if got := sanitizeURL(inURL); !reflect.DeepEqual(got, want) {
			t.Errorf("sanitizeURL(%v) returned %v, want %v", tt.in, got, want)
		}
	}
}

func TestCheckResponse(t *testing.T) {
	u, _ := url.Parse("/")
	resp := &http.Response{
		Request: &http.Request{
			Method: "GET",
			URL:    u,
		},
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(strings.NewReader("-300, 帳號密碼不得為空值。")),
	}

	got := CheckResponse(resp).(*ErrorResponse)
	if got == nil {
		t.Errorf("Expected error response.")
	}
	want := &ErrorResponse{
		Response:  resp,
		ErrorCode: StatusUsernameAndPasswordAreRequired,
		Message:   "帳號密碼不得為空值。",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Error = %#v, want %#v", got, want)
	}
	if got, want := got.Error(), "GET /: 200 -300 帳號密碼不得為空值。"; got != want {
		t.Errorf("Error = %v, want %v", got, want)
	}
}

func TestCheckResponse_unexpectedStatusCode(t *testing.T) {
	resp := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusBadRequest,
		Body:       nil,
	}

	err := CheckResponse(resp)
	if err == nil {
		t.Errorf("Expected error response.")
	}
	if got, want := err.Error(), fmt.Sprintf("unexpected status code: %d", http.StatusBadRequest); got != want {
		t.Errorf("Error = %v, want %v", got, want)
	}
}

func TestCheckResponse_invalidMessageFormat(t *testing.T) {
	resp := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(strings.NewReader("-")),
	}

	err := CheckResponse(resp)
	if err == nil {
		t.Errorf("Expected error response.")
	}
	if got, want := err.Error(), fmt.Sprint("invalid message format"); got != want {
		t.Errorf("Error = %v, want %v", got, want)
	}
}
