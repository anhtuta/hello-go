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
