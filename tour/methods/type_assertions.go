package main

import "fmt"

// Type switches
// A type switch is a construct that permits several type assertions in series
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func TypeAssertions() {
	fmt.Println("\n========== TypeAssertions ==========")

	// Type assertions
	// A type assertion provides access to an interface value's underlying concrete value.
	// To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.
	// t, ok := i.(T)
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	// f = i.(float64) // panic: interface conversion: interface {} is string, not float64

	f, ok := i.(float64) // Not panic!
	fmt.Println(f, ok)

	// Type switches
	do(21)
	do("hello")
	do(true)
}
