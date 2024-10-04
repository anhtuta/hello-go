package main

import (
	"fmt"
	"strconv"
	"time"
)

func errors() {
	fmt.Println("\n========== Errors demo ==========")

	// Go programs express error state with error values.
	// Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.
	i, err := strconv.Atoi("42years")
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
	} else {
		fmt.Println("Converted integer:", i)
	}

	if err := run(); err != nil {
		fmt.Println(err)
	}
}

type MyError struct {
	When time.Time
	What string
}

// The error type is a built-in interface similar to fmt.Stringer:
//
//	type error interface {
//		Error() string
//	}
//
// Các kiểu dữ liệu implement interface error giống như implement interface Stringer.
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
