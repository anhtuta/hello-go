package main

import (
	"fmt"
	"time"
)

//========== Goroutines ==========
// A goroutine is a lightweight thread managed by the Go runtime.
// go f(x, y, z): starts a new goroutine running f(x, y, z)\
// The evaluation of f, x, y, and z happens in the current goroutine and the execution of f happens in the new goroutine.

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}

// (typed conduit: đường ống có kiểu type nào đó???)

//========== Channels ==========
// Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
// ch <- v    // Send v to channel ch.
// v := <-ch  // Receive from ch, and assign value to v.
// (The data flows in the direction of the arrow.)
// Like maps and slices, channels must be created before use:
// ch := make(chan int)
// By default, sends and receives block until the other side is ready.
// This allows goroutines to synchronize without explicit locks or condition variables.

/*
========== Channels using Copilot ==========
Purpose:
- Channels are used for communication between goroutines.
- They provide a way to send and receive values between concurrently executing goroutines.

Blocking Behavior:
- Sending and receiving on a channel can block until the other side is ready.
- This allows for synchronization between goroutines without explicit locks.

Syntax: Channels are created using the make function and can be used to send and receive values.

Types: Channels are strongly typed, meaning a channel can only transport values of a specific type.

Concurrency: Channels are designed to work seamlessly with goroutines, providing a way to synchronize and communicate between them.
*/

/*
Compare with Java:
The term that is most similar to a channel in Go is a BlockingQueue
- Both are used for communication and synchronization between concurrently executing threads or goroutines.
- They provide a way to send and receive values, and they can block operations until the other side is ready.
*/

//========== Buffered Channels ==========
// Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:
// ch := make(chan int, 100)
// Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

//========== Range and Close ==========
// A sender can close a channel to indicate that no more values will be sent. Receivers can test
// whether a channel has been closed by assigning a second parameter to the receive expression: after
// v, ok := <-ch
// ok is false if there are no more values to receive and the channel is closed.
// Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	fmt.Println("========== Goroutines ==========")

	// Starts a new goroutine that calls the say function.
	// When you prefix a function call with the go keyword, it runs the function concurrently in a new goroutine.
	// This function call is executed immediately and synchronously in the main goroutine
	// Lời gọi hàm này được thực thi ngay lập tức ở 1 goroutine mới
	go say("world")

	// Immediately after starting the new goroutine (for above function call),
	// the say("hello") function is called in the main goroutine
	// Lời gọi hàm này được thực thi ngay lập tức ở main goroutine, sau cái goroutine mới ở trên.
	// Do 2 goroutine chạy song song nên thứ tự in ra có thể khác nhau
	say("hello")

	fmt.Println("\n========== Channels ==========")
	s := []int{7, 2, 8, -9, 4, 0}

	// The example code sums the numbers in a slice, distributing the work between two goroutines.
	// Once both goroutines have completed their computation, it calculates the final result.
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	fmt.Println("\n========== Buffered Channels ==========")
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3 // fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("\n========== Range and Close ==========")
	c1 := make(chan int, 10)
	go fibonacci(cap(c1), c1)
	// receives values from the channel repeatedly until it is closed
	for i := range c1 {
		fmt.Println(i)
	}

	// continue at https://go.dev/tour/concurrency/5
}
