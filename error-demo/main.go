package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("========== Wrap error ==========")
	// The %w verb in Go's fmt package is used for wrapping errors. When you use %w in a formatted string,
	// it indicates that the error should be wrapped, allowing you to retain the original error while adding
	// additional context (Copilot).
	originalErr := errors.New("original error")
	wrappedErr := fmt.Errorf("additional context: %w", originalErr)

	fmt.Println(originalErr) // original error

	// Explicit Call to Error() Method
	fmt.Println(wrappedErr.Error()) // additional context: original error

	// Implicit Call to Error() Method: The fmt package will call the Error() method on wrappedErr
	// if it implements the error interface.
	// If wrappedErr is a custom error type that implements the fmt.Stringer interface or has a custom
	// Format method, fmt.Println might use those methods instead of just calling Error(). This can result
	// in additional formatting or information being included in the output.
	fmt.Println(wrappedErr) // additional context: original error

	// Check if the wrapped error contains the original error
	if errors.Is(wrappedErr, originalErr) {
		fmt.Println("The wrapped error contains the original error")
	}

	fmt.Println(errors.Unwrap(wrappedErr)) // original error

	// The %v verb in Go's fmt package is used for printing the value in a default format.
	// When you use %v with an error, it prints the error message.
	fmt.Printf("%v\n", wrappedErr) // additional context: original error

	fmt.Println("\n========== Another example of Wrap error ==========")
	businessID := "12345"
	originalErr1 := errors.New("original error")
	wrappedErr1 := fmt.Errorf("error occurred: %w. business_id: %s", originalErr1, businessID)
	fmt.Println(wrappedErr1) // error occurred: original error. business_id: 12345

}
