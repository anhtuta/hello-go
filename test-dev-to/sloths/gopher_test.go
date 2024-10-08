package sloths

import (
	"os"
	"path/filepath"
	"testing"
)

// Must create the test_files directory before running this test
func TestAddGopher(t *testing.T) {
	// set up file to add a gopher to: create a file if it doesn't exist
	path := filepath.Join("test_files", "gopher-added.txt")
	f, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}

	// write some text to the file (overriding any existing content)
	if _, err := f.WriteString("Go is awesome!"); err != nil {
		t.Fatal(err)
	}
	f.Close()

	// run addGopher and test that we now have a gopher emoji
	if err := addGopher(path); err != nil {
		t.Fatal(err)
	}

	// check that the file now contains the gopher emoji
	fileContents, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(fileContents) != "Go is awesome!🐹" {
		t.Errorf(
			`unexpected file contents %s`, string(fileContents),
		)
	}

	// clean up: remove the file.
	// like Jest in JavaScript, you can use t.Cleanup in scenarios similar to where you would use
	// the Jest afterAll function; the Cleanup function runs after a test and all its sub-tests complete
	t.Cleanup(func() {
		if err != os.Remove(path) {
			t.Fatalf("error cleaning up %s: %v", path, err)
		}
	})
}
