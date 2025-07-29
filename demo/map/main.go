package main

import (
	"fmt"
	"maps"
)

// Ref: https://gobyexample.com/maps, Copilot
func main() {
	// To create an empty map, use the builtin make: make(map[key-type]val-type).
	m := make(map[string]int)

	// Syntax: name[key] = val --> to set key/value pairs
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map: ", m) // map:  map[k1:7 k2:13]

	// Count the number of key/value pairs
	fmt.Println("len:", len(m)) // 2

	// Remove a key/value pair
	delete(m, "k2")
	delete(m, "non-key")   // No error if key doesn't exist
	fmt.Println("map:", m) // map: map[k1:7]

	// Remove all key/value pairs
	clear(m)
	fmt.Println("map:", m) // map: map[]

	// Get a value for a key with name[key], when key is not in the map, you get the zero value for the value type
	// Note: zero value = false for booleans, 0 for numeric types, "" for strings, and nil for pointers, functions, interfaces, slices, channels, and maps.
	v2 := m["k2"]
	fmt.Println("v2:", v2) // v2: 0

	// The second return value when getting a value from a map indicates if the key was present in the map.
	// (Param thứ 2 trả về true nếu key tồn tại, false nếu key không tồn tại)
	_, isExist := m["k2"]
	fmt.Println("isExist:", isExist) // isExist: false
	m["k2"] = 123
	v2, isExist = m["k2"]
	fmt.Println("v2:", v2)           // v2: 123
	fmt.Println("isExist:", isExist) // isExist: true

	// Declare and initialize a new map in the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n) // map: map[foo:1 bar:2]

	// Check if 2 maps are equal, using the Equal function in the maps package
	n1 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n1) {
		fmt.Println("n and n1 are equal") // will print this
	} else {
		fmt.Println("n and n1 are not equal")
	}

	fmt.Println("\n========== Similar to Map<String, Object> in Java ==========")
	n2 := make(map[string]interface{})
	n2["foo"] = 1
	n2["bar"] = true
	n2["baz"] = "baz"
	fmt.Println("map:", n2) // map: map[bar:true baz:baz foo:1]

	// Can use any instead
	n3 := map[string]any{
		"foo": 1,
		"bar": true,
		"baz": "baz",
	}
	if maps.Equal(n2, n3) {
		fmt.Println("n2 and n3 are equal") // will print this
	} else {
		fmt.Println("n2 and n3 are not equal")
	}

	fmt.Println("\n========== Nested maps ==========")
	n4 := map[string]map[string]int{
		"foo": {"foo1": 1, "foo2": 2},
		"bar": {"bar1": 3, "bar2": 4},
	}
	fmt.Println("map:", n4) // map: map[bar:map[bar1:3 bar2:4] foo:map[foo1:1 foo2:2]]

	fmt.Println("\n========== Iterating over a map ==========")
	for key, value := range n4 {
		fmt.Println(key, value)
	}
	// Output:
	// foo map[foo1:1 foo2:2]
	// bar map[bar1:3 bar2:4]
}
