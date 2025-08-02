package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var ops int64

func increaseCount() {
	atomic.AddInt64(&ops, 1)
}

func main() {
	for i := 0; i < 1000; i++ {
		go increaseCount()
	}

	time.Sleep(time.Second)

	fmt.Println("ops:", ops)
}
