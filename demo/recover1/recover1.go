package recover1

import "fmt"

func mayPanic() {
	panic("a problem")
}

// Go makes it possible to recover from a panic, by using the recover built-in function.
// A recover can stop a panic from aborting the program and let it continue with execution instead.
func RecoverDemo() {
	fmt.Print("\n======= Recover =======\n\n")

	// recover must be called within a deferred function. When the enclosing function panics,
	// the defer will activate and a recover call within it will catch the panic.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	// This code will not run, because mayPanic panics.
	// The execution of RecoverDemo stops at the point of the panic and resumes in the deferred closure.
	fmt.Println("After mayPanic()")
}
