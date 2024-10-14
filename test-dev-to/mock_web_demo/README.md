# Using Testify Mock in web clients

Ban đầu ta có đoạn code như này:

```go
type ZooHTTPClient struct {
    baseURL string
    client *http.Client
}
func (c *ZooHTTPClient) ListAnimalFacts(q AnimalFactsQuery) (*AnimalFactsResponse, error) {
    // returning nil, nil just so the code compiles
    return nil, nil
}
func getSlothsFavoriteSnack(c *ZooHTTPClient) (string, error) {
    res, err := c.ListAnimalFacts(...)
    // ...
}
```

Function `getSlothsFavoriteSnack` có 1 dependency là 1 object kiểu `ZooHTTPClient` (kiểu concrete). Để test được nó, ta phải mock (stub) được function `ListAnimalFacts`

--> Phải dùng kiểu abstract thì mới có thể mock được `ListAnimalFacts`. Do đó ta sửa code thành như sau

```go
// Wrap method cần mock vào 1 interface
type ZooClient interface {
	ListAnimalFacts(q AnimalFactsQuery) (*AnimalFactsResponse, error)
}

// Struct ban đầu này sẽ được ngầm định là implement interface trên
type ZooHTTPClient struct {
    baseURL string
    client *http.Client
}
func (c *ZooHTTPClient) ListAnimalFacts(q AnimalFactsQuery) (*AnimalFactsResponse, error) {
    // returning nil, nil just so the code compiles
    return nil, nil
}

// ListAnimalFacts sẽ dùng 1 abstract dependency thay vì concrete dependency như trước
func getSlothsFavoriteSnack(c *ZooClient) (string, error) {
    res, err := c.ListAnimalFacts(...)
    // ...
}

// Bây giờ dễ dàng implement mock cho ZooClient bằng cách tạo 1 struct khác cũng implement ZooClient,
// sau đó truyền cho getSlothsFavoriteSnack và test
type mockClient struct{ mock.Mock }
func newMockClient() *mockClient { return &mockClient{} }
// Mock implementation of the ListAnimalFacts method of the ZooClient interface
func (c *mockClient) ListAnimalFacts(q AnimalFactsQuery) (*AnimalFactsResponse, error) {
	// Mock return here
}
```

Basic test with mock

1. Take the nondeterministic piece of functionality and wrap it in a Go interface type.
2. Write an implementation of the interface that uses Testify Mock.
3. In the tests, use the mock implementation to select deterministic results of the functions your interface calls.

Ref: https://dev.to/salesforceeng/using-testify-mock-in-web-clients-5amb
