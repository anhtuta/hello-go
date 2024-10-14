package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"hello-go/greetings"

	"rsc.io/quote"
)

func main() {
	// Set properties of the predefined Logger, including the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("[greetings]: ")
	log.SetFlags(0)

	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	// A slice of names.
	names := []string{"Gladys", "", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	// If an error was returned, print it to the console and exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned map of messages to the console..
	fmt.Println(messages)

	fmt.Println("\n============ Context demo ============")
	// Create a context with a timeout of 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Simulate a long-running operation
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("1. Operation completed")
	case <-ctx.Done():
		fmt.Println("1. Operation timed out:", ctx.Err())
	}
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("2. Operation completed")
	case <-ctx.Done():
		fmt.Println("2. Operation timed out:", ctx.Err())
	}

	// Create a context with a value
	ctx1 := context.WithValue(context.Background(), "key", "value")

	// Pass the context to a function
	process(ctx1)
	fmt.Println("Context demo completed")
}

func process(ctx context.Context) {
	// Retrieve the value from the context
	if val, ok := ctx.Value("key").(string); ok {
		fmt.Println("Value from context:", val)
	} else {
		fmt.Println("Key not found in context")
	}
}
