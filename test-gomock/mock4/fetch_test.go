package mock4_test

import (
	"bytes"
	"errors"
	"test-gomock/mock4"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestMyFunc_Success(t *testing.T) {
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

// TestMyFunc_EmptyResponse tests the case where the Fetcher.FetchData method returns an empty response
func TestMyFunc_EmptyResponse(t *testing.T) {
	// Create a new controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock object for the Fetcher interface
	mockFetcher := mock4.NewMockFetcher(ctrl)

	// Set expectations on the mock object's behavior to return an empty byte slice
	mockFetcher.EXPECT().FetchData().Return([]byte{}, nil)

	// Call the code under test
	data, err := mock4.MyFunc(mockFetcher)

	// Assert the results
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(data) != 0 {
		t.Errorf("Expected empty data, but got %v", data)
	}
}

// TestMyFunc_Failure tests the case where the Fetcher.FetchData method returns an error
func TestMyFunc_Failure(t *testing.T) {
	// Create a new controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock object for the Fetcher interface
	mockFetcher := mock4.NewMockFetcher(ctrl)

	// Set expectations on the mock object's behavior to return an error
	expectedError := errors.New("fetch error")
	mockFetcher.EXPECT().FetchData().Return(nil, expectedError)

	// Call the code under test
	data, err := mock4.MyFunc(mockFetcher)

	// Assert the results
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	if err != expectedError {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}
	if data != nil {
		t.Errorf("Expected nil data, but got %v", data)
	}
}
