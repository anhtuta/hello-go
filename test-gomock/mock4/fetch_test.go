package mock4_test

import (
	"bytes"
	"test-gomock/mock4"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestMyFunc(t *testing.T) {
	// Create a new controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock object for the Fetcher interface
	mockFetcher := mock4.NewMockFetcher(ctrl)

	// Set expectations on the mock object's behavior
	mockFetcher.EXPECT().FetchData().Return([]byte("data"), nil)

	// Call the code under test
	data, err := mock4.MyFunc(mockFetcher)

	// Assert the results
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !bytes.Equal(data, []byte("data")) {
		t.Errorf("Unexpected data: %v", data)
	}
}
