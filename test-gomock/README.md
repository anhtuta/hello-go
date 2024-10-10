# Exploring Go mocking methods and the GoMock framework

Ref: https://blog.logrocket.com/exploring-go-mocking-methods-gomock-framework/, Copilot

Note: việc đặt tên file, package, module trong repo này có thể khác so với tác giả

## Steps to setup

Các bước đã dùng để setup tutorial này, không cần run lại nữa. All commands assume that you are in `test-gomock` folder

Init this module: `go mod init test-gomock`

Install testify mock: `go get github.com/stretchr/testify/mock`

Install gomock: `go get github.com/golang/mock/gomock`

Install mockgen (binary file): `go install github.com/golang/mock/mockgen@latest`

- With Go modules, `go install` is the recommended way to **install binaries**. It ensures that the tool is installed in the `$GOPATH/bin` or `$GOBIN` directory, making it available globally
- Vì `mockgen` là 1 thư viện binary, tức là ta sẽ run nó = terminal, do đó nên dùng `go install` để cài thay vì `go get`
- Sau khi run lệnh trên, `mockgen` sẽ được cài ở thư mục `$GOPATH/bin` hoặc `$GOBIN`
- Hiện tại máy t nó cài ở đây:
  ```bash
  /Users/tuta/dev/bin/mockgen
  ➜  test-gomock git:(main) ✗ echo $GOPATH
  /Users/tuta/dev
  ➜  test-gomock git:(main) ✗ where mockgen
  /Users/tuta/dev/bin/mockgen
  /Users/tuta/dev/bin/mockgen
  ```

## Mocking techniques in Go

Có nhiều cách để mock

### 1. Manually creating mock objects

Creating a struct that implements the same interface as the real object, but with mock data and behavior

Useful when dealing with simple objects that have a small number of methods

Code: [mock1](./mock1/mock_test.go)

### 2. Implementing interfaces

Implementing interfaces directly.

This technique is best used when dealing with more complex objects with many methods

Code: [mock2](./mock2/mock_test.go)

### 3. Using function callbacks

Useful when dealing with functions that are not part of an interface

Code: [mock3](./mock3/mock_test.go)

### 4. Using the GoMock framework for mock generation

When writing unit tests for code that depends on external dependencies, such as databases or web services, it can be difficult to ensure that the tests are repeatable and predictable

Theo t hiểu thì: thay vì tự tạo mock struct và implement interface cần mock như các cách trên, thì ta có thể dùng gomock để nó tự gen code mock struct cho ta (dùng thư viện `mockgen`)

To create a mock object for this interface using GoMock, we need to follow these steps:

- Create a new controller using `gomock.NewController()`
- Create a mock object using the controller's CreateMock method, passing in the interface you want to mock
- Set expectations on the mock object’s behavior using GoMock’s methods
- Call the controller's `Finish` method to indicate you are done with the mock object

Code example: [mock4](./mock4/fetch_test.go)

The following code snippet that illustrates how to record calls to a mock object using GoMock

```go
import (
    "testing"

    "github.com/golang/mock/gomock"
)

func TestMyFunction(t *testing.T) {
    // Create a new controller
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    // Create a mock object for the MyInterface interface
    mockObj := NewMockMyInterface(ctrl)

    // Record the expected calls to the mock object
    call1 := mockObj.EXPECT().MyMethod1()
    call2 := mockObj.EXPECT().MyMethod2("arg1", "arg2")

    // Set expectations on the behavior of the calls
    call1.Return(nil)
    call2.Return("result", nil)

    // Call the code under test
    myFunction(mockObj)

    // Verify that the expected calls were made
    if err := mockObj.AssertExpectationsWereMet(); err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
}
```

Note: trước khi viết được test, phải gen code đã, code này là mock implementation cho interface muốn mock. Command to gen code (run in `mock4` folder): `mockgen -source=fetcher.go -destination=mock_fetcher.go -package=mock4`
