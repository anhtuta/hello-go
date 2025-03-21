# Go tutorial

Ref: https://go.dev/doc/

This tutorial uses version: go1.22.1 darwin/arm64

- 2024-12-31: upgraded to: go1.23.4 darwin/arm64

Install and config Go:

- Install using brew: `brew install go`
- Config go path: add it to `.zshrc` file:
  ```bash
  # go
  export GOPATH=$HOME/dev
  export PATH=$PATH:$GOPATH/bin
  ```
- Create the needed directories: `mkdir -p $GOPATH/{src,pkg,bin}`
- Các binary file được cài bằng lệnh `go install` sẽ được lưu ở thư mục `$GOPATH/bin`
- Các package được cài bằng lệnh `go get` sẽ được lưu ở thư mục `$GOPATH/pkg`
- Đang dùng máy mac nên thư mục go path sẽ là: `/Users/tuta/dev`

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

## Cấu trúc của project này

Project này sẽ gồm nhiều module

- Mỗi module là 1 tutorial nhỏ về 1 topic nào đó trên internet
- Mỗi module sẽ có 1 thư mục riêng (có thể nằm trong 1 thư mục khác để dễ group và dễ quản lý)
- Mỗi module sẽ có 1 file `go.mod`, dòng đầu tiên của file đó sẽ define tên module, e.g. `module hello-go/greetings`. Tên module và tên thư mục KHÔNG phải trùng nhau
- Mỗi module thường sẽ có 1 file `main.go` ở thư mục root của module nó (cùng level với `go.mod`), thường thì sẽ run file này

```
hello-go/
    ├── module1/
    │   ├── go.mod
    │   ├── main.go
    ├── greetings/
    │   ├── go.mod
    │   ├── main.go
    ├── parent-folder
    │   ├── module3/
    │   │   ├── go.mod
    │   │   ├── main.go
    │   ├── module4/
    │   │   ├── go.mod
    │   │   ├── main.go
```

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

Dereference pointer: giả sử function có 1 param kiểu con trỏ, thì khi gọi field/method của biến đó, ta phải dereference con trỏ đó

```go
func getSlothsFavoriteSnack(c *ZooClient) (string, error) {
    // Dereference pointer
    res, err := (*c).ListAnimalFacts(...)
}
```

Detail code: [getSlothsFavoriteSnack](./test-dev-to/mock_web_demo/zoo_service.go)

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

# Goroutines vs Thread in Java

Goroutines in Go are similar to threads in Java in that they both allow for concurrent execution of code. However, there are some key differences between them:

Similarities

- Concurrent Execution: Both goroutines and Java threads enable concurrent execution of code, allowing multiple tasks to run simultaneously.
- Lightweight: Both are designed to be lightweight and efficient in terms of resource usage. (nhưng goroutines lightweight hơn)

Differences

- Creation and Management:
  - Goroutines: Managed by the Go runtime, which **handles scheduling and execution**. They are **extremely lightweight**, with a **small initial stack size** that **grows and shrinks as needed**.
  - Java Threads: Managed by the Java Virtual Machine (JVM) and the underlying operating system. They have a **fixed stack size** and are generally **more resource-intensive** than goroutines.
- Syntax:

  - Goroutines: Created using the `go` keyword
    ```go
    go say("world")
    ```
  - Java Threads: Created by extending the `Thread` class or implementing the `Runnable` interface.
    ```java
    Thread t = new Thread(new Runnable() {
        public void run() {
            say("world");
        }
    });
    t.start();
    ```

- Communication:

  - Goroutines: Communicate using **channels**, which provide a way to **send and receive values between goroutines**.
    ```go
    ch := make(chan int)
    ch <- 1 // Send value to channel
    val := <-ch // Receive value from channel
    ```
  - Java Threads: Communicate using **shared variables, synchronized blocks**, and other concurrency utilities provided by the java.util.concurrent package.
    ```java
    synchronized(sharedObject) {
        // Access shared object
    }
    ```

- Blocking Behavior:
  - Goroutines: By default, sends and receives on channels block until the other side is ready. This allows for synchronization without explicit locks.
  - Java Threads: Blocking behavior is managed using synchronization primitives like `wait()`, `notify()`, and `notifyAll()`, or higher-level constructs like `CountDownLatch` and `Semaphore`.

Ref: Copilot

# Steps to work with Testify Mock

1. Create an interface that defines the behavior of the nondeterministic code
2. Create a struct that implements the interface
3. Use the struct in the code you want to test
4. Create a mock of the interface in the test
5. Set the expected behavior of the mock
6. Use the mock in the test

Ref: Copilot

Note: Mọi file test phải có tên kết thúc bằng `_test.go`, nếu không khi run cmd `go test`, nó sẽ không detect được để run

# Terminology

- `go.mod`: Defines the module's properties and dependencies (NOT for package). It's placed at the root of the module directory.
- Package:
  - A collection of source files in the same directory that are compiled together.
  - `package main`: A special package that defines a standalone executable program. It must contain a main function, which is the entry point of the program.
- Module: A collection of related Go packages that are versioned together.

Example:

```
my-module/
    go.mod
    main.go
    pkg1/
        file1.go
        file1_1.go
    pkg2/
        file2.go
```

Hiểu đơn giản thì:

- 1 module sẽ được đặt trong 1 folder (chẳng hạn `my-module` như ở trên), bên trong này có thể chứa nhiều package khác nhau
- File `go.mod` được đặt ở thư mục root của module (như hình trên), dùng để định nghĩa dependencies cho module
- Mỗi package sẽ có 1 folder riêng, bên trong sẽ gồm các file `.go`, chúng đều khai báo chung 1 package ở đầu file
- The `main.go` file contains the **entry point of the program**, nó trông như này:

  ```go
  package main

  import (
      "fmt"
      "my-module/pkg1"
      "my-module/pkg2"
  )

  func main() {
      fmt.Println("Hello, World!")
      pkg1.Function1()
      pkg2.Function2()
  }
  ```

- Các file bên trong 1 package sẽ trông như này:

  ```go
  // pkg1/file1.go
  package pkg1

  import "fmt"

  func Function1() {
      fmt.Println("Function1 from pkg1")
  }
  ```

# Type Inference

2 cách sau đều dùng để cast biến `mockHandler` sang kiểu `http.HandlerFunc`, khác nhau ở chỗ:

- Explicit Type Declaration: `var mockHandler http.HandlerFunc = handler`

  - Explicitly states the type of the variable.
  - Can be **more readable** and clear, especially for those unfamiliar with type inference.

- Type Inference: `var mockHandler = http.HandlerFunc(handler)`

  - Uses type inference to determine the type of the variable.
  - More concise but relies on understanding the type returned by `http.HandlerFunc(handler)`.

Code: [httptest_demo](./test-dev-to/httptest_demo/server_test.go)

Ref: Copilot

# Value types vs Reference types

Go có 2 thể loại kiểu dữ liệu: kiểu giá trị và kiểu tham chiếu

Value Types

- Integers (`int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`)
- Floating-point numbers (`float32`, `float64`)
- Booleans (`bool`)
- Strings (`string`)
- Arrays
- Structs

Reference Types

- Slices
- Maps
- Channels
- Pointers
- Functions

## All reference types in detail with examples

Full code: [reference_types.go](./tour/moretypes/reference_types.go)

### 1. Slices

- **Description**: Slices are dynamically-sized, flexible views into the elements of an array. They are reference types because they contain a pointer to the underlying array, along with a length and capacity.
- **Example**:

  ```go
  func modifySlice(s []int) {
      s[0] = 42
  }

  func main() {
      slice := []int{1, 2, 3}
      modifySlice(slice)
      fmt.Println(slice) // Output: [42, 2, 3]
  }
  ```

### 2. Maps

- **Description**: Maps are collections of key-value pairs. They are reference types because they contain a pointer to the underlying hash table.
- **Example**:

  ```go
  func modifyMap(m map[string]int) {
      m["key"] = 42
  }

  func main() {
      m := map[string]int{"key": 1}
      modifyMap(m)
      fmt.Println(m) // Output: map[key:42]
  }
  ```

### 3. Channels

- **Description**: Channels are used for communication between goroutines. They are reference types because they contain a pointer to the underlying channel data structure.
- **Example**:

  ```go
  func sendMessage(ch chan string) {
      ch <- "Hello, World!"
  }

  func main() {
      ch := make(chan string)
      go sendMessage(ch)
      msg := <-ch
      fmt.Println(msg) // Output: Hello, World!
  }
  ```

### 4. Pointers

- **Description**: Pointers hold the memory address of a value. They are explicitly reference types.
- **Example**:

  ```go
  func modifyValue(p *int) {
      *p = 42
  }

  func main() {
      value := 1
      modifyValue(&value)
      fmt.Println(value) // Output: 42
  }
  ```

### 5. Functions

- **Description**: Functions are reference types because they contain a pointer to the function's code and environment.
- **Example**:

  ```go
  func modifyFunction(f func() int) func() int {
      return func() int {
          return f() + 1
      }
  }

  func main() {
      f := func() int { return 1 }
      f = modifyFunction(f)
      fmt.Println(f()) // Output: 2
  }
  ```

### Summary

- **Slices**: Reference types that point to an underlying array.
- **Maps**: Reference types that point to an underlying hash table.
- **Channels**: Reference types that point to an underlying channel data structure.
- **Pointers**: Explicit reference types that hold memory addresses.
- **Functions**: Reference types that point to function code and environment.

## Function nên return kiểu nào

- Với kiểu giá trị thì nên return con trỏ, bởi vì:

  - Efficiency: avoids copying the entire object: không phải copy kết quả sang 1 object khác rồi mới return
  - Modification: Có thể modify response
  - Nil Check: vì nó là con trỏ nên có thể check `nil`

- Với kiểu tham chiếu: **không cần return con trỏ**, tức là khi return 1 giá trị kiểu map, ta cũng đã ngầm định return 1 tham chiếu trỏ tới hashmap rồi (chứ KHÔNG phải return a copy of the entire map)
- Example:

  ```go
  type MyInterface interface {
    Func1(ctx context.Context, request *EmployeeRequest) (*EmployeeResponse, error)
    Func2(ctx context.Context, request *EmployeeRequest) (map[EmployeeResponse]string, error)
  }
  ```

## Compare to Java

Value Types

- Primitive types: `byte`, `short`, `int`, `long`, `float`, `double`, `char`, `boolean`

Reference Types

- Arrays
- Classes (including custom classes)
- Interfaces
- Enums
- `String`
- `Object`

Ref: Copilot

# Object dùng như function nếu nó có kiểu function type

```go
// 1 kiểu function type
// PayrollProviderFactoryFunc is a function to create a PayrollProvider service for a given provider
type PayrollProviderFactoryFunc func(provider string) (PayrollService, error)

// Ta có object sau thuộc kiểu function type, nó có thể là 1 field của 1 struct
payrollProviderFactory       PayrollProviderFactoryFunc

// Giờ có thể gọi biến payrollProviderFactory như gọi function
payrollService, err := pr.payrollProviderFactory("Provider name")
if err != nil {
  return nil, errors.NewNotImplementedError("provider not supported").Error()
}
```

TODO update this...

# Transitioning from Java to Go: terms and concepts

Here are some terms and concepts that a Java developer should be familiar with when working with Go:

1. **Goroutines**:

- **Description**: Lightweight threads managed by the Go runtime.
- **Java Equivalent**: Similar to threads but more lightweight and managed by the Go runtime.
- **Example**:
  ```go
  go func() {
      fmt.Println("Hello from a goroutine")
  }()
  ```

2. **Channels**:

- **Description**: Used for communication between goroutines.
- **Java Equivalent**: Similar to blocking queues or other concurrency utilities.
- **Example**:
  ```go
  ch := make(chan int)
  go func() {
      ch <- 42
  }()
  value := <-ch
  fmt.Println(value)
  ```

3. **Interfaces**:

- **Description**: Define a set of method signatures that a type must implement.
- **Java Equivalent**: Similar to interfaces in Java.
- **Example**:
  ```go
  type Stringer interface {
      String() string
  }
  ```

4. **Structs**:

- **Description**: Composite data types that group together variables.
- **Java Equivalent**: Similar to classes but without inheritance.
- **Example**:
  ```go
  type Person struct {
      Name string
      Age  int
  }
  ```

5. **Slices**:

- **Description**: Dynamically-sized arrays.
- **Java Equivalent**: Similar to ArrayLists but more efficient.
- **Example**:
  ```go
  var numbers []int
  numbers = append(numbers, 1, 2, 3)
  ```

6. **Maps**:

- **Description**: Key-value pairs.
- **Java Equivalent**: Similar to HashMaps.
- **Example**:
  ```go
  var m map[string]int
  m = make(map[string]int)
  m["one"] = 1
  ```

7. **Pointers**:

- **Description**: Variables that hold memory addresses.
- **Java Equivalent**: Java does not have explicit pointers, but references in Java are somewhat similar.
- **Example**:
  ```go
  var p *int
  i := 42
  p = &i
  ```

8. **Packages**:

- **Description**: Group related Go files together.
- **Java Equivalent**: Similar to packages in Java.
- **Example**:

  ```go
  package main

  import "fmt"

  func main() {
      fmt.Println("Hello, World!")
  }
  ```

9. **Error Handling**:

- **Description**: Go uses explicit error handling rather than exceptions.
- **Java Equivalent**: Similar to checked exceptions but handled explicitly.
- **Example**:
  ```go
  func divide(a, b int) (int, error) {
      if b == 0 {
          return 0, fmt.Errorf("division by zero")
      }
      return a / b, nil
  }
  ```

10. **Function Types**:

- **Description**: Functions can be assigned to variables and passed around as arguments. Functions can be used as first-class citizens.
- **Java Equivalent**: Similar to functional interfaces or lambda expressions.
- **Example**:

  ```go
  type Adder func(int, int) int

  func main() {
      var add Adder = func(a, b int) int {
          return a + b
      }
      fmt.Println(add(1, 2))
  }
  ```

11. **Defer**:

- **Description**: Schedules a function call to be run after the function completes.
- **Java Equivalent**: Similar to finally blocks but more flexible.
- **Example**:
  ```go
  func main() {
      defer fmt.Println("This will be printed last")
      fmt.Println("This will be printed first")
  }
  ```

12. **Go Modules**:

- **Description**: Dependency management system for Go.
- **Java Equivalent**: Similar to Maven or Gradle.
- **Example**:
  ```sh
  go mod init mymodule
  go get github.com/some/dependency
  ```

Ref: Copilot
