# A tour of Go

Ref: https://go.dev/tour/list

Cách tạo (init) từng module

```sh
➜  tour git:(main) ✗ mkdir basic
➜  tour git:(main) ✗ cd basic
➜  basic git:(main) ✗ go mod init tour/basic
➜  basic git:(main) ✗ touch main.go
```

Cách module khác làm tương tự

List of tutorials

- [Packages, variables, and functions](./basic/main.go): Learn the basic components of any Go program.
- [Flow control statements: for, if, else, switch and defer](./flowcontrol/main.go): Learn how to control the flow of your code with conditionals, loops, switches and defers
- [More types: structs, slices, and maps](./moretypes/main.go): Learn how to define types based on existing ones: this lesson covers structs, arrays, slices, and maps
- [Methods and interfaces](./methods/main.go): This lesson covers methods and interfaces, the constructs that define objects and their behavior
- [Generics](./generics/main.go): Go supports generic programming using type parameters. This lesson shows some examples for employing generics in your code
