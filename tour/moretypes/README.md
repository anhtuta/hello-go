How to run:

```bash
go run .
```

If we run `go run main.go`, we will get an error:

```sh
# command-line-arguments
./main.go:11:2: undefined: pointerAndStruct
./main.go:12:2: undefined: arrayAndSlice
./main.go:13:2: undefined: sliceLenCapDemo1
./main.go:14:2: undefined: sliceLenCapDemo2
./main.go:15:2: undefined: forRange
./main.go:16:2: undefined: mapDemo
./main.go:17:2: undefined: functionsAreValues
./main.go:18:2: undefined: referenceTypes
```

This happens because `go run .` compiles and runs all Go files in the current directory, so all functions defined in other files (like `pointerAndStruct`, `arrayAndSlice`, etc.) are included.

However, `go run main.go` only compiles and runs `main.go`, so it doesn't see the function definitions in the other files. That's why you get undefined errors for those functions.
`
To summarize:

- `go run .`: Runs all Go files in the directory (works as expected).
- `go run main.go`: Runs only `main.go` (fails if `main.go` depends on functions in other files).

To run your program successfully, use `go run .` or `go run *.go`.
