package main

import "fmt"

// Ref: Copilot
// Naming Conventions
// camelCase: Used for unexported (private) functions.
// PascalCase: Used for exported (public) functions.
func EmptyInterfaceDemo() {
	fmt.Println("\n========== EmptyInterfaceDemo ==========")
	// The empty interface can store any value, regardless of its type. This makes it a universal container.
	// Usage: It is often used in situations where you need to handle values of unknown or varying types, such as in generic data structures, functions that accept any type, or when working with JSON data.
	// Type Assertion: To retrieve the original value from an empty interface, you typically use type assertions or type switches.
	var i interface{}

	// Storing different types in the empty interface
	i = 42
	fmt.Println(i) // Output: 42

	i = "hello"
	fmt.Println(i) // Output: hello

	i = true
	fmt.Println(i) // Output: true

	// Type assertion to retrieve the original value
	if str, ok := i.(string); ok {
		fmt.Println("String value:", str) // Output: String value: hello
	} else {
		fmt.Println("Not a string")
	}

	// Type switch
	// A type switch can be used to handle different types stored in an empty interface:
	var i1 interface{} = "hello"

	switch v := i1.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	case bool:
		fmt.Println("Boolean:", v)
	default:
		fmt.Println("Unknown type")
	}

	doSomething()
}

func doSomething() {
	fmt.Println("doSomething")
}
