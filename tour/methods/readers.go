package main

import (
	"fmt"
	"io"
	"strings"
)

func readers() {
	fmt.Println("\n========== Readers demo ==========")

	// The io package specifies the io.Reader interface, which represents the read end of a stream of data.
	// This example creates a strings.Reader and consumes its output 8 bytes at a time.
	r := strings.NewReader("Hello, Reader! Today is a good day.")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
