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

// Ref: https://batnamv.medium.com/go-concurrency-d%C3%A0nh-cho-java-developers-c7709f1f8752
func main() {
	for i := 0; i < 1000; i++ {
		go incrCount()
	}
	time.Sleep(time.Second)
	defer fmt.Println(counter) // 1000, as expected due to mutex protection
}
