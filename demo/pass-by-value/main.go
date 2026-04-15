package main

import "fmt"

// This is still pass-by-value.
// It looks like pass-by-reference because the value of a changes, but under the hood, Go is doing exactly what it always does: copying.
// Here is why this is technically "pass-by-value" (specifically, passing a pointer by value):

// 1. The Copy: When you call modify(&a), Go takes the memory address of a (e.g., 0x123) and copies that address into a new variable x inside the function.
// 2. The Pointer: x is a separate variable that just happens to hold the same memory address as &a.
// 3. The Dereference: When you do *x = 100, you are telling Go: "Go to the address stored in x and change the data there."
func modify(x *int) {
	*x = 100 // Directly changes the original variable
}

// The Ultimate Proof:
// In a true pass-by-reference language (like C++), the parameter x is an alias for a. In Go, x is a distinct variable. You can prove this by checking if you can change where x points.
// If Go were pass-by-reference, reassigning x inside the function would potentially reassign a in the main scope. Since it doesn't, it's just a value copy of a pointer.
func modify1(x *int) {
	newVal := 200
	x = &newVal // You are changing the LOCAL copy of the pointer
	// This will NOT change 'a' in main() to 200
}

var globalNum = 500

// In Go, you can change the data at an address, but you cannot change the original variable to point somewhere else by passing it into a function
func reassign(p *int) {
	// 'p' is a COPY of the pointer.
	// We are only changing the local copy to point to globalNum.
	p = &globalNum
}

func main() {
	a := 10
	modify(&a)
	println(a) // Output: 100

	a = 10
	modify1(&a)
	println(a) // Output: 10, not 200

	localNum := 10
	ptr := &localNum
	fmt.Println("Before:", *ptr) // Prints 10
	reassign(ptr)
	fmt.Println("After:", *ptr) // Still prints 10
}
