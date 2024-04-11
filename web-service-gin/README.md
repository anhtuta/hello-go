Create a module via `init` command. This command creates a go.mod file in which dependencies you add will be listed for tracking:

```bash
go mod init neuvillette/web-service-gin
```

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