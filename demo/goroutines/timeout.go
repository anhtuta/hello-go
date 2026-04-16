package goroutines

import (
	"fmt"
	"time"
)

// Timeouts are important for programs that connect to external resources or that otherwise need to bound execution time
func TimeoutDemo() {
	fmt.Print("\n======= Timeout =======\n\n")

	// Execute an external call that returns its result on a channel c1 after 2s
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// the select implementing a timeout
	select {
	case res := <-c1: // await the result
		fmt.Println("res from c1:", res)
	case <-time.After(1 * time.Second): // await a value to be sent after the timeout of 1s
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	// If we allow a longer timeout of 3s, then the receive from c2 will succeed and we’ll print the result
	select {
	case res := <-c2:
		fmt.Println("res from c2:", res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
