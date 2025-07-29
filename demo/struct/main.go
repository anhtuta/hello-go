package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson1(name string) *Person {
	p := Person{name: name}
	p.age = 42
	return &p
}

// Go doesn't support constructor overloading (Polymorphism)
func newPerson2(name string, age int) *Person {
	p := Person{name: name, age: age}
	return &p
}

// Ref: https://gobyexample.com/structs
func main() {
	fmt.Println(Person{"Bob", 20})              // struct literal
	fmt.Println(Person{name: "Alice", age: 30}) // struct fields
	fmt.Println(Person{name: "Fred"})           // Omitted fields will be zero-valued
	fmt.Println(&Person{name: "Ann", age: 40})  // An & prefix yields a pointer to the struct

	// It’s idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson1("Jon"))
	fmt.Println(newPerson2("Jon", 30))

	// Access struct fields with a dot
	s := Person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// You can also use dots with struct pointers - the pointers are automatically dereferenced
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable
	sp.age = 51
	fmt.Println(sp) // &{Sean 51} --> sp is a pointer to s, so changing sp will change s

	sClone := s
	sClone.age = 52
	fmt.Println(s)      // {Sean 51} --> s is not changed, because sClone is a copy of s
	fmt.Println(sClone) // {Sean 52}

	// anonymous struct type
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog) // {Rex true}
}
