package mock_web_demo

// File này là implementation của ZooClient interface

import (
	"net/http"
)

// Struct này sẽ implement ZooClient interface
type ZooHTTPClient struct {
	BaseURL string
	Client  *http.Client
}

// Function này sẽ được call bên service, nhưng bên service khi test mock sẽ stub nó,
// do đó không cần phải implement gì ở đây.
func (c *ZooHTTPClient) ListAnimalFacts(q AnimalFactsQuery) (*AnimalFactsResponse, error) {
	// HTTP implementation here; returns an
	// AnimalFactsResponse if the HTTP request succeeds,
	// or an error, of type ErrorResponse, if the request
	// gets a non-2xx HTTP status code.

	// returning nil, nil just so the code compiles
	return nil, nil
}

type AnimalFactsResponse struct {
	Facts         []string `json:"facts"`
	AreThereMore  bool     `json:"are_there_more"`
	NextPageToken string   `json:"next_page_token"`
}
