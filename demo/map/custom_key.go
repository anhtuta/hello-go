package main

import "fmt"

// Go's runtime handles hashing internally. You don't define or expose a hash function —
// Go computes the hash of a struct key automatically based on its fields.
type Address struct {
	Name    string
	city    string
	Pincode int
}

func mapCustomKey() {
	a2 := Address{Name: "Ram", city: "Delhi", Pincode: 2400}
	a1 := Address{"Pam", "Dehradun", 2200}
	a3 := Address{Name: "Sam", city: "Lucknow", Pincode: 1070}

	sampleMap := map[Address]int{a1: 10, a2: 20, a3: 3}
	fmt.Println(sampleMap)
}
