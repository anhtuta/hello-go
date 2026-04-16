package defer1

import (
	"fmt"
	"os"
	"path/filepath"
)

// Defer is used to ensure that a function call is performed later in a program’s execution,
// usually for purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages.
func DeferDemo() {
	fmt.Print("\n======= Defer =======\n\n")

	path := filepath.Join(os.TempDir(), "defer.txt")
	f := createFile(path)

	// We defer the closing of that file.
	// This will be executed at the end of the enclosing function (DeferDemo), after writeFile has finished.
	defer closeFile(f)

	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("file created at", p)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	// It’s important to check for errors when closing a file, even in a deferred function
	if err != nil {
		panic(err)
	}
}
