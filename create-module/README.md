# Tutorial: Create a Go module

https://go.dev/doc/tutorial/create-module

Run code:

- `cd hello`
- `go run hello.go`

Build code: compile the code into an executable

- `cd hello`
- `go build`
- Run the built file (`hello` executable): `./hello`

Install the executable so you can run it without specifying its path

- First, check your `GOPATH`:
  ```bash
  ➜  hello-go git:(main) ✗ echo $GOPATH
  /Users/tuta/dev
  ```
- Next, install the package (in the `hello` folder): `go install`
- Run your application: you now can run it in anywhere, just open a new terminal and test it: `hello`
