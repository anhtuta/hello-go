package httptest_demo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// ========== Testing your handler with httptest ==========
// handleSlothfulMessage is a regular Go function, NOT a web server,
// therefore we can run test without any actual HTTP server
func TestHandleSlothfulMessage(t *testing.T) {
	// We make our net/http ResponseWriter implementation with NewRecorder.
	// Note type của wr là con trỏ struct httptest.ResponseRecorder, trong khi param đầu tiên của method
	// handleSlothfulMessage là 1 kiểu interface http.ResponseWriter. Điều này không sai, vì ResponseRecorder
	// implement cái interface đó (check source code sẽ thấy)
	var wr *httptest.ResponseRecorder = httptest.NewRecorder() // dùng var để thấy rõ kiểu

	// Make a Request object pointed at our /sloth endpoint using NewRequest
	var req *http.Request = httptest.NewRequest(http.MethodGet, "/sloth", nil)

	// Test
	// Note: param đầu tiên có kiểu http.ResponseWriter, đây là 1 interface, nên khi gọi có thể truyền vào bất kỳ kiểu
	// nào mà implement interface đó, bao gồm cả 1 kiểu con trỏ của struct, nếu như struct đó dùng pointer receiver.
	// If a type ResponseRecorder implements the http.ResponseWriter interface, then *ResponseRecorder can also
	// be used as an http.ResponseWriter if the methods are defined with pointer receivers.
	handleSlothfulMessage(wr, req)

	// Verify
	if wr.Code != http.StatusOK {
		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
	}

	if !strings.Contains(wr.Body.String(), "Stay slothful!") {
		t.Errorf(
			`response body "%s" does not contain "Stay slothful!"`,
			wr.Body.String(),
		)
	}
}

// Test POST method
func TestHandleSlothfulMessagePost(t *testing.T) {
	// We make our net/http ResponseWriter implementation with NewRecorder.
	// Note type của wr là con trỏ struct httptest.ResponseRecorder, trong khi param đầu tiên của method
	// handleSlothfulMessage là 1 kiểu interface http.ResponseWriter. Điều này không sai, vì ResponseRecorder
	// implement cái interface đó (check source code sẽ thấy)
	var wr *httptest.ResponseRecorder = httptest.NewRecorder() // dùng var để thấy rõ kiểu

	// Make a Request object pointed at our /sloth endpoint using NewRequest
	var req *http.Request = httptest.NewRequest(http.MethodPost, "/sloth-post", nil)

	// Test
	// Note: param đầu tiên có kiểu http.ResponseWriter, đây là 1 interface, nên khi gọi có thể truyền vào bất kỳ kiểu
	// nào mà implement interface đó, bao gồm cả 1 kiểu con trỏ của struct, nếu như struct đó dùng pointer receiver.
	// If a type ResponseRecorder implements the http.ResponseWriter interface, then *ResponseRecorder can also
	// be used as an http.ResponseWriter if the methods are defined with pointer receivers.
	handleSlothfulMessagePost(wr, req)

	// Verify
	if wr.Code != http.StatusOK {
		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
	}

	if !strings.Contains(wr.Body.String(), "Stay slothful!") {
		t.Errorf(
			`response body "%s" does not contain "Stay slothful!"`,
			wr.Body.String(),
		)
	}
}

// ========== Testing the httptest server ==========
// When you want to do functional tests or integration tests for a web app

// The client we're making has a GetSlothfulMessage that sends an HTTP request to the /sloth of its baseURL.
// We can test this by creating a real HTTP server that listens on a random port, and then sending a request to it.
type Client struct {
	httpClient *http.Client
	baseURL    string
}

type SlothfulMessage struct {
	Message string `json:"message"`
}

func NewClient(httpClient *http.Client, baseURL string) Client {
	return Client{
		httpClient: httpClient,
		baseURL:    baseURL,
	}
}

func (c *Client) GetSlothfulMessage() (*SlothfulMessage, error) {
	res, err := c.httpClient.Get(c.baseURL + "/sloth")
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"got status code %d", res.StatusCode,
		)
	}

	var m SlothfulMessage
	if err := json.NewDecoder(res.Body).Decode(&m); err != nil {
		return nil, err
	}
	return &m, nil
}

// Ta sẽ tạo 1 http server thật, dùng handler đã được define ở file code để handle request
func TestGetSlothfulMessage(t *testing.T) {
	// Start a new server, on a randomized port, but use the same handler
	router := http.NewServeMux()
	router.HandleFunc("/sloth", handleSlothfulMessage)
	svr := httptest.NewServer(router)
	defer svr.Close()

	// Set up our client and have it test an HTTP roundtrip to our /sloth endpoint
	c := NewClient(http.DefaultClient, svr.URL)
	m, err := c.GetSlothfulMessage()
	if err != nil {
		t.Fatalf("error in GetSlothfulMessage: %v", err)
	}
	if m.Message != "Stay slothful!" {
		t.Errorf(
			`message %s should contain string "Sloth"`,
			m.Message,
		)
	}
}
