package main

import "fmt"

func referenceTypes() {
	fmt.Println("\n========== Reference types ==========")

	fmt.Println("1. Slice is a reference type, array is a value type")
	slice := []int{1, 2, 3}
	modifySlice(slice)
	fmt.Println(slice) // Output: [42, 2, 3]

	arr := [3]int{1, 2, 3}
	modifyArray(arr)
	fmt.Println(arr) // Output: [1, 2, 3]

	fmt.Println("2. Map")
	m := map[string]int{"key": 1}
	modifyMap(m)
	fmt.Println(m) // Output: map[key:42]

	fmt.Println("3. Channel")
	ch := make(chan string)
	go sendMessage(ch)
	msg := <-ch
	fmt.Println(msg) // Output: Hello, World!

	fmt.Println("4. Pointer")
	pointer := 1
	modifyPointer(&pointer)
	fmt.Println(pointer) // Output: 42

	fmt.Println("5. Function")
	f := func() int { return 1 }
	f = modifyFunction(f)
	fmt.Println(f()) // Output: 2
}

func modifySlice(s []int) {
	s[0] = 42
}

// Array is a value type, so it will not be modified
func modifyArray(a [3]int) {
	a[0] = 42
}

func modifyMap(m map[string]int) {
	m["key"] = 42
}

func sendMessage(ch chan string) {
	ch <- "Hello, World!"
}

func modifyPointer(p *int) {
	*p = 42
}

func modifyFunction(f func() int) func() int {
	return func() int {
		return f() + 1
	}
}
