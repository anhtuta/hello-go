package main

import "fmt"

// The type parameters of a function appear between brackets, before the function's arguments.
// This Index function returns the index of x in s, or -1 if not found.
// This declaration means that s is a slice of any type T that fulfills the built-in constraint comparable.
// x is also a value of the same type.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

// Note (Copilot): the built-in constraint comparable is satisfied by any type whose values can be compared using
// the == and != operators. This includes:
// Basic Types: All numeric types (e.g., int, float64), string, bool
// Pointer Types: Pointers to any type
// Array Types: Arrays of comparable types
// Interface Types: Interfaces that do not contain methods
// Struct Types: Structs where all fields are of comparable types
// Channel Types: Channels of comparable types

func typeParameters() {
	fmt.Println("\n========== Type parameters ==========")

	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}
