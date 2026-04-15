package main

import (
	"demo/array"
	"demo/closure"
	"demo/constants"
	rangeovertypes "demo/range-over-types"
	"demo/slice"
	stringsandrunes "demo/strings-and-runes"
	"fmt"
)

func main() {
	// https://gobyexample.com/constants
	fmt.Println("My nickname is", constants.MY_NICKNAME)
	fmt.Println("My university is", constants.MY_UNIVERSITY)
	fmt.Println("My company is", constants.MY_COMPANY)
	fmt.Println("Another constant is", constants.ANOTHER_CONSTANT)

	// https://gobyexample.com/arrays
	array.ArrayDemo()

	// https://gobyexample.com/slices
	slice.SliceDemo()

	// https://gobyexample.com/closures
	fmt.Print("\n======= Closures =======\n\n")
	nextInt := closure.IntSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	newInts := closure.IntSeq()
	fmt.Println(newInts())

	// https://gobyexample.com/range-over-built-in-types
	rangeovertypes.RangeDemo()

	// https://gobyexample.com/strings-and-runes
	stringsandrunes.RuneDemo()
}
