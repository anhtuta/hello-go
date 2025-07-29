package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var mutex sync.Mutex

func incrCount() {
	mutex.Lock()
	counter++
	mutex.Unlock()
}

func main() {
	for i := 0; i < 1000; i++ {
		go incrCount()
	}
	time.Sleep(time.Second)
	defer fmt.Println(counter) // 1000, as expected due to mutex protection
}
