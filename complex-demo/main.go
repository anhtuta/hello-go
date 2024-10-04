package main

import (
	"complex-demo/complex"
	"fmt"
)

func main() {
	z := complex.New(1, 2)
	fmt.Println(z.Add(z))

	z1 := complex.New(3, 4)
	z2 := complex.New(5, 6)
	z3 := z1.Add(z2)
	fmt.Println(z3)
}
