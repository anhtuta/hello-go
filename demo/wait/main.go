package main

import (
	"fmt"
	"sync"
	"time"
)

// This is the function we'll run in every goroutine.
// Note that a WaitGroup must be passed to functions by pointer.
func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", id)

	// Sleep to simulate an expensive task.
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
	// On return, notify the WaitGroup that we're done.
	defer wg.Done()
}

// Ref: https://batnamv.medium.com/go-concurrency-d%C3%A0nh-cho-java-developers-c7709f1f8752
func main() {

	// This WaitGroup is used to wait for all the goroutines launched here to finish.
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup counter for each.
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Block until the WaitGroup counter goes back to 0; all the workers notified they're done.
	wg.Wait()

	fmt.Println("All workers done")
}
