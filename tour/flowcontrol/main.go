package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// If with a short statement
// Like for, the if statement can start with a short statement to execute before the condition.
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	sum := 0
	// Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding
	// the three components of the for statement and the braces { } are always required
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum1 := 1
	// The init and post statements are optional.
	// Có thể bỏ qua ; nếu không có init và post statements.
	// for ; sum < 1000; {
	for sum1 < 1000 {
		sum1 += sum1
	}
	fmt.Println(sum1)

	// Go's if statements are like its for loops; the expression need not be surrounded by
	// parentheses ( ) but the braces { } are required.
	if sum1 < 1000 {
		fmt.Println("sum1 < 1000")
	} else {
		fmt.Println("sum1 >= 1000")
	}

	// If with a short statement
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// Switch
	// It runs the first case whose value is equal to the condition expression.
	// Go only runs the selected case, not all the cases that follow. In effect, the break statement
	// that is needed at the end of each case in those languages is provided automatically in Go.
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// Switch evaluation order
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// Switch with no condition
	// Switch without a condition is the same as switch true.
	// This construct can be a clean way to write long if-then-else chains.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// Defer
	// A defer statement defers the execution of a function until the surrounding function returns.
	// The deferred call's arguments are evaluated immediately, but the function call is not executed
	// until the surrounding function returns
	defer fmt.Println("...world. This line will be printed last.")
	fmt.Println("hello...")
	fmt.Println("Today is a good day.")

	// Stacking defers
	// Deferred function calls are pushed onto a stack. When a function returns,
	// its deferred calls are executed in last-in-first-out order.
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
	fmt.Println("deferred function calls are pushed onto a stack.")
	fmt.Println("When a function returns, its deferred calls are executed in last-in-first-out order.")

}
