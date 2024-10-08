# Intro to automated testing in Go

Ref: https://dev.to/salesforceeng/intro-to-automated-testing-in-go-4mjl

## Setup

All commands assume that you are in `test-dev-to` folder

Init this module: `go mod init test-dev-to`

Create test folder: `mkdir sloths/test-files`

Install testify mock: `go get github.com/stretchr/testify/mock`

## Run tests

Use one of these commands, run in `sloths` folder:

```
go test
go test -v -test.short
```
