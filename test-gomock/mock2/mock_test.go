package mock2

import (
	"fmt"
	"testing"
)

type Database interface {
	Get(key string) (string, error)
	Set(key, value string) error
}

type MockDatabase struct {
	GetFunc func(key string) (string, error)
	SetFunc func(key, value string) error
}

func (db *MockDatabase) Get(key string) (string, error) {
	return db.GetFunc(key)
}

func (db *MockDatabase) Set(key, value string) error {
	return db.SetFunc(key, value)
}

// Ref: Copilot
func TestDatabaseOperations(t *testing.T) {
	mockDB := &MockDatabase{
		GetFunc: func(key string) (string, error) {
			if key == "existingKey" {
				return "value", nil
			}
			return "", fmt.Errorf("key not found")
		},
		SetFunc: func(key, value string) error {
			if key == "" {
				return fmt.Errorf("key cannot be empty")
			}
			return nil
		},
	}

	// Test Get operation
	value, err := mockDB.GetFunc("existingKey")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != "value" {
		t.Errorf("Unexpected value: %v", value)
	}

	// Test the Get method with a non-existing key
	_, err = mockDB.GetFunc("nonExistingKey")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Test Set operation
	err = mockDB.SetFunc("newKey", "newValue")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
