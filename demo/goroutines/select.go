package goroutines

import (
	"fmt"
	"time"
)

// select lets you wait on multiple channel operations
func SelectDemo() {
	fmt.Print("\n======= Select =======\n\n")

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// for range 2 is valid syntax in Go 1.22+.
	// for range 2 means: loop exactly 2 times.
	// This is a feature introduced in Go 1.22 called "range over integers." It's a shorthand for looping N times without needing to track a loop variable.
	// It's equivalent to: for i := 0; i < 2; i++ { ... }
	for range 2 {
		// use select to await both of these values simultaneously, printing each one as it arrives
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
