package goroutines

import (
	"demo/util"
	"fmt"
	"time"
)

// The done channel will be used to notify another goroutine that this function’s work is done.
func worker(done chan bool) {
	fmt.Println("working...", util.GetGID())
	time.Sleep(time.Second)
	fmt.Println("done", util.GetGID())

	// Send a value to notify that we’re done
	done <- true
}

func ChannelSyncDemo() {
	fmt.Print("\n======= Channel Synchronization =======\n\n")

	// Start a worker goroutine, giving it the channel to notify on
	fmt.Println("Starting worker from main...", util.GetGID())
	done := make(chan bool, 1)
	go worker(done)
	fmt.Println("Waiting for worker to finish...", util.GetGID())

	// Block until we receive a notification from the worker on the channel
	<-done
}
