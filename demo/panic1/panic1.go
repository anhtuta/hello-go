package panic1

import (
	"fmt"
	"os"
	"path/filepath"
)

// A panic typically means something went unexpectedly wrong.
// Mostly we use it to fail fast on errors that shouldn’t occur during normal operation,
// or that we aren’t prepared to handle gracefully.
// A common use of panic is to abort if a function returns an error value that we don’t know how to (or want to) handle.
// Note: unlike some languages which use exceptions for handling of many errors,
// in Go it is idiomatic to use error-indicating return values wherever possible
func PanicDemo() {
	fmt.Print("\n======= Panic =======\n\n")

	// When first panic in main fires, the program exits without reaching the rest of the code.
	// panic("a problem")

	path := filepath.Join(os.TempDir(), "file")
	_, err := os.Create(path)
	if err != nil {
		// Panic if we get an unexpected error when creating a new file.
		panic(err)
	} else {
		fmt.Println("Created file:", path)
	}
}
