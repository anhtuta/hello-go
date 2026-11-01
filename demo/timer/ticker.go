package timer

import (
	"fmt"
	"time"
)

// Timers are for when you want to do something once in the future
// Tickers are for when you want to do something repeatedly at regular intervals
func TickerDemo() {
	fmt.Print("\n======= Ticker =======\n\n")

	// Tickers use a similar mechanism to timers: a channel that is sent values
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t) // Value will arrive every 500ms
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
