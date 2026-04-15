# Demo

Module này giống như project Demo bên repo [spring](https://github.com/anhtuta/spring) vậy, sẽ gồm các ví dụ, khái niệm đơn giản nhất về Go

Init this module: `go mod init demo`

Run any example (in the `demo` folder): `go run sub-folder-name/main.go`

## Code structure

Mỗi 1 Go module thường sẽ chỉ có 1 file `main.go`, it is the entry point for a Go program

Nhưng ta sẽ tạo nhiều thư mục trong module này, mỗi thư mục sẽ chứa 1 file `main.go`, do đó:

- Mỗi 1 thư mục là 1 ví dụ nhỏ, chúng đều là `package main` vì chứa file `main.go`
- Bên Java, mỗi 1 class đều có thể có hàm `main` và run file đó được, nhưng với Go, ta phải tạo nhiều thư mục nếu muốn có nhiều hàm `main` để chạy, vì
  - Mỗi 1 thư mục chỉ được phép khai báo 1 package, e.g. `package main` (ở đầu mỗi file.go)
  - Mỗi 1 thư mục chỉ được phép có 1 hàm `main()`

## How to run

```sh
cd demo
go run go-by-example-main/main.go
```

Note: Ko thể đứng ở thư mục cha rồi run đc, nếu ko sẽ lỗi:

```sh
➜  hello-go git:(main) ✗ go run demo/go-by-example-main/main.go
demo/go-by-example-main/main.go:4:2: package demo/array is not in std (/opt/homebrew/Cellar/go/1.26.2/libexec/src/demo/array)
demo/go-by-example-main/main.go:5:2: package demo/closure is not in std (/opt/homebrew/Cellar/go/1.26.2/libexec/src/demo/closure)
demo/go-by-example-main/main.go:6:2: package demo/constants is not in std (/opt/homebrew/Cellar/go/1.26.2/libexec/src/demo/constants)
demo/go-by-example-main/main.go:7:2: package demo/range-over-types is not in std (/opt/homebrew/Cellar/go/1.26.2/libexec/src/demo/range-over-types)
demo/go-by-example-main/main.go:8:2: package demo/slice is not in std (/opt/homebrew/Cellar/go/1.26.2/libexec/src/demo/slice)
demo/go-by-example-main/main.go:9:2: package demo/strings-and-runes is not in std (/opt/homebrew/Cellar/go/1.26.2/libexec/src/demo/strings-and-runes)
```

Ref:

- https://gobyexample.com
