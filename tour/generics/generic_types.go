package main

import "fmt"

// List represents a singly-linked list that holds values of any type.
// Exercise, add some functionality to this list implementation.
type List[T any] struct {
	next *List[T]
	val  T
}

// This implementation is written by Copilot
// Add appends a new element to the end of the list.
func (l *List[T]) Add(value T) {
	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = &List[T]{val: value}
}

// Print traverses the list and prints each element.
func (l *List[T]) Print() {
	current := l
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}
}

func genericTypes() {
	fmt.Println("\n========== Generic types ==========")

	// Create a new list with an initial element.
	head := &List[int]{val: 1}

	// Add elements to the list.
	head.Add(2)
	head.Add(3)
	head.Add(4)

	// Print the elements of the list.
	head.Print()
}
