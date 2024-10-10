package main

import "strconv"

// If the number is divisible by 3, return "Foo", otherwise return the number as a string
func Fooer(input int) string {
	isFoo := (input % 3) == 0
	if isFoo {
		return "Foo"
	}
	return strconv.Itoa(input)
}
