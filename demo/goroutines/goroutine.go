package goroutines

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

// https://gobyexample.com/goroutines
func GoroutineDemo() {
	fmt.Print("\n======= Goroutines =======\n\n")

	// running it synchronously
	f("direct")

	// This new goroutine will execute concurrently with the calling one
	go f("goroutine")

	// start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
