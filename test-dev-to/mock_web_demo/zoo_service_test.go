package mock_web_demo

// 2. Writing a mock implementation of our client

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type mockClient struct{ mock.Mock }

func newMockClient() *mockClient { return &mockClient{} }

// Mock implementation of the ListAnimalFacts method of the ZooClient interface
func (c *mockClient) ListAnimalFacts(q AnimalFactsQuery) (*AnimalFactsResponse, error) {
	args := c.Called(q)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*AnimalFactsResponse), args.Error(1)
}

var allPageAPIReq = AnimalFactsQuery{
	AnimalName: "sloth",
	PageToken:  "ALL", // Phải giống hệt với bên service để mock đúng
}

// 3. Using our client in a test
func TestGetSlothsFavoriteSnack_FoundFavoriteSnack(t *testing.T) {
	// Create a new mock client, this one implements the ZooClient interface and mocks the ListAnimalFacts method
	c := newMockClient()

	// Stub the ListAnimalFacts method to return a response with a fact about sloths' favorite snack
	c.On("ListAnimalFacts", allPageAPIReq).Return(&AnimalFactsResponse{
		Facts: []string{
			"Sloths' slowness is actually used as a form of camouflage",
			"Baby sloths make the cutest li'l squeak 🥰",
			"Sloths' favorite snack is hibiscus flowers",
		},
		AreThereMore:  false,
		NextPageToken: "",
	}, nil)

	// Test
	favSnack, err := getSlothsFavoriteSnack(c)

	// Assert
	if err != nil {
		t.Fatalf("got error getting sloths' favorite snack: %v", err)
	}
	if favSnack != "hibiscus flowers" {
		t.Errorf(
			"expected favorite snack to be hibiscus flowers, got %s",
			favSnack,
		)
	}
}

func TestGetSlothsFavoriteSnack_NotFoundFavoriteSnack(t *testing.T) {
	// Create a new mock client, this one implements the ZooClient interface and mocks the ListAnimalFacts method
	c := newMockClient()

	// Stub the ListAnimalFacts method to return a response with a fact about sloths' favorite snack
	c.On("ListAnimalFacts", allPageAPIReq).Return(&AnimalFactsResponse{
		Facts: []string{
			"Sloths' slowness is actually used as a form of camouflage",
			"Baby sloths make the cutest li'l squeak 🥰",
			"Sloths detest people who don't use semicolons",
		},
		AreThereMore:  false,
		NextPageToken: "",
	}, nil)

	// Test
	favSnack, err := getSlothsFavoriteSnack(c)

	// Assert
	if err == nil {
		t.Fatalf("expected error getting sloths' favorite snack, but got nil")
	}
	if favSnack != "" {
		t.Errorf(
			"expected favorite snack to be empty, got %s",
			favSnack,
		)
	}
}

func TestGetSlothsFavoriteSnack_500Error(t *testing.T) {
	// Create a new mock client
	c := newMockClient()

	// Stub the ListAnimalFacts method to return an error
	c.On("ListAnimalFacts", allPageAPIReq).Return(
		(*AnimalFactsResponse)(nil), &ErrorResponse{
			StatusCode: 500,
			Message:    "server error",
		},
	)

	// Test
	_, err := getSlothsFavoriteSnack(c)

	// Assert
	if err == nil {
		t.Fatal("got nil error from getSlothsFavoriteSnack, but expected an error")
	}

	// The type assertion err.(*ErrorResponse) checks if err is of type *ErrorResponse.
	errRes, ok := err.(*ErrorResponse)
	if !ok {
		t.Fatalf("expected error to be ErrorResponse, got %T", err)
	}
	if status := errRes.StatusCode; status != 500 {
		t.Errorf("expected 500, got %d status code", status)
	}
}
