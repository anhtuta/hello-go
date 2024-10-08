package mock_web_demo

import "fmt"

// Serialization of error response from zoo API service
type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (res *ErrorResponse) Error() string {
	return fmt.Sprintf("got %d error: %s", res.StatusCode, res.Message)
}
