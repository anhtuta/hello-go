package main

import "fmt"

func changeValue(val int) {
	// val is a copy of x.
	val = 20
}

func changeValueViaPointer(ptr *int) {
	// ptr is a copy of the address in p.
	// *ptr deferences the address, changing the original data.
	*ptr = 20
}

func tryToReassignPointer(ptr *int) {
	// ptr is a copy of the address in p.
	y := 30
	// This line ONLY changes the local copy 'ptr'.
	// It now points to y.
	// The caller's original 'p' is unaffected.
	ptr = &y
}

// Ref: Gemini 2.5 Pro
func main() {
	x := 10
	p := &x // p holds the memory address of x

	// 1. Pass by value (primitive)
	changeValue(x)
	fmt.Println("After changeValue:", x) // Prints 10

	// 2. Pass by value (of a pointer) - This modifies the original data.
	changeValueViaPointer(p)
	fmt.Println("After changeValueViaPointer:", x) // Prints 20

	// 3. The "True" Test: Try to change the caller's pointer
	x = 10 // Reset x
	p = &x // Reset p

	tryToReassignPointer(p)
	fmt.Println("After tryToReassignPointer:", *p) // Still prints 10
}
