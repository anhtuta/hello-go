Create a module via `init` command. This command creates a go.mod file in which dependencies you add will be listed for tracking (lệnh này ban đầu dùng để khởi tạo module web-service-gin, cái tên neuvillette để cho hay chứ chả có ý nghĩa gì, có thể dùng tên khác tuỳ ý):

```bash
go mod init neuvillette/web-service-gin
```

Start this module (web server): `go run .`

Get albums, and get by ID:

```bash
curl http://localhost:9081/albums

# More detail:

curl http://localhost:9081/albums \
    --header "Content-Type: application/json" \
    --request "GET"

# By ID
curl http://localhost:9081/albums/2
```

Create a new album:

```bash
curl http://localhost:9081/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```