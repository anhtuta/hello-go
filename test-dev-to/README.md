# Intro to automated testing in Go

Ref: https://dev.to/salesforceeng/intro-to-automated-testing-in-go-4mjl

Note: việc đặt tên file, package, module trong repo này có thể khác so với tác giả

## Steps to setup

Các bước đã dùng để setup tutorial này, không cần run lại nữa. All commands assume that you are in `test-dev-to` folder

Init this module: `go mod init test-dev-to`

Install testify mock: `go get github.com/stretchr/testify/mock`

## Run tests

First, create test folder: `mkdir sloths/test_files`

Use one of these commands, can run in each package (`sloths`, `mock_demo` or `httptest_demo` folder...):

```
go test
go test -v -test.short
```

## List of tutorials

1. Intro to automated testing in Go: [sloths.go](./sloths/sloths.go)
2. Subtesting, skipping, and cleanup in the Go testing.T: [sloths.go](./sloths/sloths.go), [gopher.go](./sloths/gopher.go)
3. Mocks in Go tests with Testify Mock: [rand.go](./mock_demo/rand.go)
4. Test your Go web apps with httptest: [server.go](./httptest_demo/server.go)
