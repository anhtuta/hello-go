Init this module: `go mod init env-demo`

Chúng ta sẽ có vài cách để thiết lập giá trị cho các biến môi trường từ bên ngoài:

- Sử dụng câu lệnh export (dành cho Linux/Mac) trong Terminal:
  ```bash
  export MYSQL_CONNECTION="root:my-root-pass@tcp(127.0.0.1:3306)/todo_db?charset=utf8mb4&parseTime=True&loc=Local"
  go run main.go
  ```
- Set trực tiếp biến môi trường ngay trong câu lệnh chạy chương trình:
  ```bash
  MYSQL_CONNECTION="root:my-root-pass@tcp(127.0.0.1:3306)/todo_db?charset=utf8mb4&parseTime=True&loc=Local" go run main.go
  ```
- Thiết lập biến môi trường để chạy debug trong Visual Studio Code: tạo file `launch.json` như sau, sau đó debug như thường
  ```json
  {
    "version": "0.2.0",
    "configurations": [
      {
        "name": "Launch Package",
        "type": "go",
        "request": "launch",
        "mode": "auto",
        "program": "${workspaceFolder}",
        "env": {
          "MYSQL_CONNECTION": "root:my-root-pass@tcp(127.0.0.1:3306)/todo_db?charset=utf8mb4&parseTime=True&loc=Local"
        }
      }
    ]
  }
  ```

Ref: https://200lab.io/blog/golang-environment-variable-bien-moi-truong-trong-golang/
