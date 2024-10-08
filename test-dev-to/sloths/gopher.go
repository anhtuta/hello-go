package sloths

import "os"

// A function that appends the gopher 🐹 (actually a hamster) emoji to a file:
func addGopher(filepath string) error {
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte("🐹"))
	return err
}
