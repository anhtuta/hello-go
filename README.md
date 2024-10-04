# Go tutorial

Ref: https://go.dev/doc/

# `go.mod` vs `pom.xml` file

The `go.mod` file in Go is somewhat similar to the `pom.xml` file in Java. Both are used for managing project dependencies

The `go.mod` file:

- This is a file used by Go Modules, which is Go's dependency management system.
- It keeps track of the module's dependencies and their versions.
- It's automatically updated when you add import statements to your Go code and then run go build or go get.

Ref: Copilot

# Multiple `go.mod` Files

In Go, the use of multiple `go.mod` files typically indicates that your project is structured as a **multi-module repository**. This is different from a **single-module repository**, which would have just one go.mod file at the root.

- Single-Module Repository
  - Java (`pom.xml`): A single `pom.xml` file is used to manage dependencies, build configurations, and project metadata for the entire project.
  - Go (`go.mod`): A single `go.mod` file at the root of the project manages dependencies for the entire project.
- Multi-Module Repository
  - Java (`pom.xml`): In a multi-module Maven project, you might have a parent `pom.xml` and multiple child `pom.xml` files in submodules.
  - Go (`go.mod`): In a multi-module Go project, each module has its own `go.mod` file. This allows each module to manage its own dependencies independently.

Why Use Multiple `go.mod` Files?

- Independent Versioning: Each module can be versioned independently, which is useful if different parts of your project are released separately.
- Dependency Isolation: Dependencies are isolated to the module that needs them, reducing the risk of dependency conflicts.
- Reusability: Modules can be reused across different projects without pulling in unnecessary dependencies.

Here's an example of a multi-module Go project structure:

```
my-multi-module-project/
    ├── module1/
    │   ├── go.mod
    │   ├── main.go
    ├── module2/
    │   ├── go.mod
    │   ├── main.go
    ├── go.mod (optional, if the root itself is a module)
```

Ref: Copilot

# Pointer

Khai báo giống C/C++: dùng dấu `*` để khai báo con trỏ, và dùng dấu `&` để lấy địa chỉ của biến

```go
a, b := 10, 20
type Vertex struct {
	X int
	Y int
}

// Khai báo biến p là con trỏ kiểu int
var p *int

// Cho p trỏ tới ô nhớ của biến a (nói ngắn gọn: cho p trỏ tới a)
p = &a

// Vừa khai báo con trỏ p kiểu int và trỏ tới a
p := &a

// Đọc p tức là đọc vị trí ô nhớ nó trỏ tới
// Đọc *p là đọc giá trị mà nó trỏ tới
fmt.Println("p = ", p)   // p lưu trữ địa chỉ của a
fmt.Println("*p = ", *p) // *p lấy giá trị của biến mà p trỏ tới

// Thao tác trên p là sẽ thay đổi vị trí mà nó sẽ trỏ tới,
// Thao tác trên *p là thay đổi giá trị của biến nó trỏ tới (không thay đổi vị trí)
p = &b // giờ trỏ sang b

// KHÔNG thể gán cho p giá trị 1 số int, mà chỉ có thể gán cho p địa chỉ của biến khác
// p = 5

*p = 15 // thay đổi b = 15

// Tạo con trỏ kiểu Vertex
v := Vertex{1, 2}
*p = &v

// Thay đổi field X của v thông qua con trỏ
(*p).X = 5

// Với con trỏ kiểu object, có thể bỏ qua * khi truy cập biến thành viên
p.X = 5
```

Detail code: [More types](./tour/moretypes/main.go)

# Composition

In Go, when you write `type A B`

- It means that you are defining a new type `A` that has the same underlying type as `B`.
- It does not mean that `A` is "extended" from `B` (ko giống như Java, C++)
- Instead, `A` is a distinct type that is based on the underlying type `B`.

In Go, there is **no concept of inheritance**. Instead, Go uses composition and interfaces to achieve similar functionality.

- Composition: Use struct embedding to reuse functionality
  - Go encourages the use of composition over inheritance. You can embed types within other types to reuse functionality.
  - This is often referred to as "struct embedding."
- Interfaces: Define behavior and achieve polymorphism
  - Go uses interfaces to define behavior. Any type that implements the methods of an interface is considered to implement that interface.
  - This allows for polymorphism without the need for inheritance.

Ref: Copilot

# Create and initialize objects in Go

Go does not have a `new` keyword, it uses a different approach to object creation and memory allocation.

Go provides built-in functions and idiomatic patterns for creating and initializing objects.

- Built-in Functions: Go provides the `new` and `make` built-in functions for memory allocation.

  ```go
  p := new(int) // p is of type *int

  *p = 42
  s := make([]int, 10) // s is a slice of 10 integers
  ```

- Composite Literals: Go uses composite literals for creating and initializing structs, arrays, slices, and maps.

  ```go
  // Struct Initialization:
  type Person struct {
      Name string
      Age  int
  }
  p := Person{Name: "Alice", Age: 30}

  // Array Initialization:
  a := [3]int{1, 2, 3}

  // Slice Initialization:
  s := []int{1, 2, 3}

  // Map Initialization:
  m := map[string]int{"one": 1, "two": 2}
  ```

- Idiomatic Patterns: Go encourages the use of constructor functions for more complex initialization. These functions typically **start with `New`** and return an initialized instance of a type.

  ```go
  type Person struct {
      Name string
      Age  int
  }

  func NewPerson(name string, age int) *Person {
      return &Person{Name: name, Age: age}
  }

  p := NewPerson("Alice", 30)
  ```

Ref: Copilot

# Explicitly return a pointer

Return an object from a function

- Java: Returning an object from a function **implicitly returns a reference** to that object (the caller receives a reference to the same object in memory, not a copy)
- Go: To return a reference to an object, you must **explicitly return a pointer**. Otherwise, Go will return a copy of the object (the caller can't modify the original object).

Example:

```java
public class Person {
    private String name;
    public Person(String name) {
        this.name = name;
    }
    public static Person createPerson(String name) {
        return new Person(name); // Returns a reference to the new Person object
    }
}
```

```go
package main

import "fmt"

type Person struct {
    Name string
}

// Function returning a pointer to a Person
func NewPerson(name string) *Person {
    return &Person{Name: name} // Returns a pointer to the new Person object
}

// Function returning a copy of a Person
func NewPersonCopy(name string) Person {
    return Person{Name: name} // Returns a copy of the new Person object
}

func main() {
    p := NewPerson("Alice")
    fmt.Println(p.Name) // Output: Alice

    pCopy := NewPersonCopy("Bob")
    fmt.Println(pCopy.Name) // Output: Bob
}
```

Nên return pointer khi dùng Go, vì

- Efficiency: avoid copying the entire struct, which can be more efficient, especially for large structs.
- Mutability: allow the caller to modify the struct's fields.
- Nil Checks: allow the caller to check for nil, which can be useful for error handling and initialization checks.
- Consistency: Many Go libraries and idiomatic patterns use pointers for struct types, making it easier to integrate with other code.

Java's implicit references: **all objects are reference types**, so we don't need explicit pointers (mặc định mọi kiểu object trong java là kiểu tham chiếu, nên khỏi phải dùng pointer cho nó mệt!)

Ref: Copilot
