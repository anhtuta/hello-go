package main

import (
	"fmt"
	"sync"
	"time"
)

//========== Mutex ==========
// We've seen how channels are great for communication among goroutines.
// What if we just want to make sure only one goroutine can access a variable at a time to avoid conflicts?
// This concept is called mutual exclusion, and the conventional name for the data structure that provides it is mutex.
// Go's standard library provides mutual exclusion with sync.Mutex and its two methods: Lock, Unlock
// We can define a block of code to be executed in mutual exclusion by surrounding it with a call to Lock and Unlock as shown on the Inc method.
// We can also use defer to ensure the mutex will be unlocked as in the Value method

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mutex  sync.Mutex
	strMap map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mutex.Lock()
	// Lock so only one goroutine at a time can access the map c.strMap.
	c.strMap[key]++
	fmt.Println("Working on", key, "with value", c.strMap[key])
	time.Sleep(500 * time.Millisecond) // Simulate some work
	c.mutex.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mutex.Lock()
	// Lock so only one goroutine at a time can access the map c.strMap.
	defer c.mutex.Unlock()
	return c.strMap[key]
}
