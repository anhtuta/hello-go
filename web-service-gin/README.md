# Developing a RESTful API with Go and Gin

https://go.dev/doc/tutorial/web-service-gin

## Cách tạo 1 module step by step

Đầu tiên phải tạo thư mục `web-service-gin`, sau đó cd vào đó

Tiếp theo, khởi tạo module web-service-gin bằng lệnh `init` (cái tên neuvillette để cho hay chứ chả có ý nghĩa gì, có thể dùng tên khác tuỳ ý):

```bash
go mod init neuvillette/web-service-gin
```

Sau đó tạo file `main.go` và bắt đầu code thôi

## Start this module

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
