# Intro to automated testing in Go

Ref: https://dev.to/salesforceeng/intro-to-automated-testing-in-go-4mjl

## Setup

All commands assume that you are in `test-dev-to` folder

Init this module: `go mod init test-dev-to`

Create test folder: `mkdir sloths/test_files`

Install testify mock: `go get github.com/stretchr/testify/mock`

## Run tests

Use one of these commands, run in `sloths` folder, or `mock_demo` folder:

```
go test
go test -v -test.short
```

## List of tutorials

1. Intro to automated testing in Go: [sloths.go](./sloths/sloths.go)
2. Subtesting, skipping, and cleanup in the Go testing.T: [sloths.go](./sloths/sloths.go), [gopher.go](./sloths/gopher.go)
3. Mocks in Go tests with Testify Mock: [rand.go](./mock_demo/rand.go)
