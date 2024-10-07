# Comprehensive Guide to Testing in Go

Ref: https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/

## Setup

All commands assume that you are in `test-goland` folder

Init this module: `go mod init test-goland`

Install testify: `go get github.com/stretchr/testify`

After this, we will get some error:

```
/Users/tuta/dev/pkg/mod/github.com/stretchr/testify@v1.9.0/assert/assertions.go:20:2: missing go.sum entry for module providing package github.com/davecgh/go-spew/spew (imported by github.com/stretchr/testify/assert); to add:
        go get github.com/stretchr/testify/assert@v1.9.0
FAIL    test-goland [setup failed]
```

Maybe this library needs some dependencies, let install them by following its suggestion:

```
go get github.com/stretchr/testify/assert@v1.9.0
```

## Run tests

Use one of these commands:

```
go test
go test -v -test.short
go test -fuzz FuzzFooer -fuzztime=20s
```
