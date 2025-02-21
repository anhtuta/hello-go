# Demo

Module này giống như project Demo bên repo [spring](https://github.com/anhtuta/spring) vậy, sẽ gồm các ví dụ, khái niệm đơn giản nhất về Go

Init this module: `go mod init demo`

Run any example (in the `demo` folder): `go run sub-folder-name/main.go`

## Code structure

Mỗi 1 Go module thường sẽ chỉ có 1 file `main.go`, it is the entry point for a Go program

Nhưng ta sẽ tạo nhiều thư mục trong module này, mỗi thư mục sẽ chứa 1 file `main.go`, do đó:

- Mỗi 1 thư mục là 1 ví dụ nhỏ, chúng đều là `package main` vì chứa file `main.go`
- Bên Java, mỗi 1 class đều có thể có hàm `main` và run file đó được, nhưng với Go, ta phải tạo nhiều thư mục nếu muốn có nhiều hàm `main` để chạy (mỗi 1 thư mục chỉ có 1 hàm `main`)
