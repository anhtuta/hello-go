# Logger service

## Commands used (don't need to run anymore)

Install dependencies:

```bash
go get github.com/youmark/pkcs8
go get google.golang.org/grpc
go get google.golang.org/protobuf
```

## Generate code (rerun after modifying proto files)

Generate Go files from proto files: run in `/logger-service` folder

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    logs/logs.proto
```
