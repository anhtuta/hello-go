package mock1

import "errors"

type Database interface {
	Get(key string) (string, error)
	Set(key, value string) error
}

type MockDatabase struct {
	Data map[string]string
}

func (db *MockDatabase) Get(key string) (string, error) {
	value, ok := db.Data[key]
	if !ok {
		return "", errors.New("key not found")
	}
	return value, nil
}

func (db *MockDatabase) Set(key, value string) error {
	db.Data[key] = value
	return nil
}
