package main

import "fmt"

type Vertex1 struct {
	Lat, Long float64
}

func mapDemo() {
	fmt.Println("\n========== Map demo ==========")

	// Giống hashmap trong Java, gồm key và value.
	// The make function returns a map of the given type, initialized and ready for use
	var m = make(map[string]Vertex1) // type of m: map[string]Vertex1
	m["Bell Labs"] = Vertex1{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// Map literals are like struct literals, but the keys are required.
	// If the top-level type is just a type name, you can omit it from the elements of the literal
	var m1 = map[string]Vertex1{
		"Bell Labs": Vertex1{
			40.68433, -74.39967,
		},
		"Google": { // KHÔNG cần khai báo kiểu của struct
			37.42202, -122.08408,
		},
	}
	fmt.Println(m1) // map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]

	// Mutating Maps
	// Insert or update an element in map m:
	m1["Bell Labs"] = Vertex1{1, 2}

	// Retrieve an element:
	fmt.Println("The value:", m1["Google"]) // The value: {37.42202 -122.08408}

	// Delete an element:
	delete(m1, "Google")
	fmt.Println("The value:", m1["Google"]) // The value: {0 0}

	// Test that a key is present with a two-value assignment:
	// If key is in m, ok is true. If not, ok is false.

	// Element "Google" is deleted, so ok = false
	v, ok := m1["Google"]
	fmt.Println("The value:", v, "Present?", ok) // The value: {0 0} Present? false

}
