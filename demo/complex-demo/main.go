package main

import (
	// demo is the module name, complex is the package name,
	// complex-demo/complex is the location of the complex package in the demo module
	"demo/complex-demo/complex"
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
