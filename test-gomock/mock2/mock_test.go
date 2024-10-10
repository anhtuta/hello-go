package mock2

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
