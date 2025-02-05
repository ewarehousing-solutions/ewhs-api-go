package ewhs

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/ewarehousing-solutions/ewhs-api-go/test/testdata"
	"github.com/stretchr/testify/assert"
)

var (
	tMux    *http.ServeMux
	tServer *httptest.Server
	tClient *Client
	tConf   *Config
)

var (
	noPre    = func() {}
	crashSrv = func() {
		u, _ := url.Parse(tServer.URL)
		tClient.BaseURL = u
	}
)

func setup() {
	tMux = http.NewServeMux()
	tServer = httptest.NewServer(tMux)

	tConf = NewConfig("test_username", "test_password", "test_wms", "test_customer", true)
	tClient, _ = NewClient(nil, tConf)

	u, _ := url.Parse(tServer.URL + "/")
	tClient.BaseURL = u
}

func teardown() {
	tServer.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testHeader(t *testing.T, r *http.Request, header string, want string) {
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %q, want %q", header, got, want)
	}
}

func testQuery(t *testing.T, r *http.Request, want string) {
	if r.URL.Query().Encode() != want {
		t.Errorf("Query().Encode() retuned unexpected values, want: %q, got %q", want, r.URL.Query().Encode())
	}
}

func unauthorizedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(testdata.WrongCredentialsResponse))
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(testdata.InternalServerErrorResponse))
}

func encodingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{hello: [{},]}`))
}

// <----- .Testing helpers ----->

func TestNewClient(t *testing.T) {
	var c = http.DefaultClient
	{
		c.Timeout = 25 * time.Second
	}

	tests := []struct {
		name   string
		client *http.Client
	}{
		{
			"nil returns a valid client",
			nil,
		},
		{
			"a passed client is decorated",
			c,
		},
	}

	conf := NewConfig("test_username", "test_password", "test_wms", "test_customer", true)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewClient(tt.client, conf)
			assert.Nil(t, err)
		})
	}
}

func TestCheckResponse(t *testing.T) {
	res1 := &http.Response{
		StatusCode: http.StatusNotFound,
		Status:     http.StatusText(http.StatusNotFound),
		Body:       io.NopCloser(strings.NewReader("not found ok")),
	}

	res2 := &http.Response{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Body:       io.NopCloser(strings.NewReader("success ok")),
	}

	res3 := &http.Response{
		StatusCode: http.StatusNotFound,
		Status:     http.StatusText(http.StatusNotFound),
		Body:       io.NopCloser(strings.NewReader("")),
	}

	tests := []struct {
		name string
		code string
		arg  *Response
	}{
		{
			"not found response",
			"Not Found",
			&Response{Response: res1},
		},
		{
			"successful response",
			"",
			&Response{Response: res2},
		},
		{
			"success with empty body",
			"",
			&Response{Response: res3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckResponse(tt.arg); err != nil {
				if !strings.Contains(err.Error(), tt.code) {
					t.Error(err)
				}
			}
		})
	}
}

func TestClient_Authorize(t *testing.T) {
	var c = http.DefaultClient
	{
		c.Timeout = 25 * time.Second
	}

	tests := []struct {
		name   string
		client *http.Client
	}{
		{
			"nil returns a valid client",
			nil,
		},
		{
			"a passed client is decorated",
			c,
		},
	}

	conf := NewConfig("test_username", "test_password", "test_wms", "test_customer", true)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewClient(tt.client, conf)
			assert.Nil(t, err)
		})
	}
}

func TestAddUserAgentString(t *testing.T) {
	tests := []struct {
		name           string
		initialAgent   string
		addition       string
		expectedAgent  string
	}{
		{
			name:          "adds string to existing user agent",
			initialAgent:  "ewhs-api-go/1.0.4 go/1.16 test_app/test_version",
			addition:      "custom/1.0",
			expectedAgent: "ewhs-api-go/1.0.4 go/1.16 test_app/test_version custom/1.0",
		},
		{
			name:          "adds string to empty user agent",
			initialAgent:  "",
			addition:      "custom/1.0",
			expectedAgent: " custom/1.0",
		},
		{
			name:          "adds empty string to user agent",
			initialAgent:  "ewhs-api-go/1.0.4",
			addition:      "",
			expectedAgent: "ewhs-api-go/1.0.4 ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test client with the initial user agent
			client := &Client{
				userAgent: tt.initialAgent,
			}

			// Call the function
			result := addUserAgentString(client, tt.addition)

			// Check the result
			assert.Equal(t, tt.expectedAgent, result.userAgent, "User agent string should match expected value")
		})
	}
}