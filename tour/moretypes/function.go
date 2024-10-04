package main

import (
	"fmt"
	"math"
)

func functionsAreValues() {
	fmt.Println("\n========== Functions are values ==========")

	// Functions are values too. They can be passed around just like other values.
	// Function values may be used as function arguments and return values.
	// Function values can be used as arguments to other functions or assigned to variables.
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12)) // 13

	fmt.Println(compute(hypot))    // sqrt(3^2 + 4^2) -> 5
	fmt.Println(compute(math.Pow)) // 3^4 -> 81

	// Function closures
	// Go functions may be closures. A closure is a function value that references variables from outside its body.
	// The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
	// Giống closure trong JavaScript.
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
	// The adder function returns a closure. Each closure is bound to its own sum variable.

	fmt.Println("\n========== Exercise: Fibonacci closure ==========")

	// Lần gọi đầu tiên: f() = f0, các lần gọi tiếp theo: f() = f1, f() = f2...
	fmt.Println("My implementation:")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("f%d: %d\n", i, f())
	}

	fmt.Println("My implementation with enhancement:")
	f1 := fibonacci1()
	for i := 0; i < 10; i++ {
		fmt.Printf("f%d: %d\n", i, f1())
	}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// (My implementation) fibonacci is a function that returns a function that returns an int.
// Note: cannot calculate f0 (=0), only can calculate f1, f2, f3, f4, ...
func fibonacci() func() int {
	f0 := 0
	f1 := 1
	return func() int {
		f2 := f0 + f1
		f0 = f1
		f1 = f2
		// return f2 --> Sai, vì f0 = 0, f1 = 1, f2 = 1
		return f0
	}
}

// Enhancement: can calculate f0
func fibonacci1() func() int {
	a, b := 0, 1
	return func() int {
		res := a
		a, b = b, a+b
		return res
	}
}
