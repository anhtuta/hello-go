package mock_demo

import "math/rand"

const MAX_RANDOM_NUMBER = 10

type RandNumberGenerator interface {
	// Should return a random integer between 0 and max
	randomInt(max int) int
}

// our standard-library implementation is an empty struct whose randomInt method calls math/rand.Intn
type StandardRand struct{}

func (s StandardRand) randomInt(max int) int {
	return rand.Intn(max)
}

// 1. Take the nondeterministic function calls and wrap them in an interface.
// For DivByRand, the nondeterministic code we're working with is math/rand's Intn function.
// We will create RandNumberGenerator interface (above) that wraps the rand.Intn function.
// Here is the code before:
//
//	func DivByRand(numerator int) int {
//		return numerator / int(rand.Intn(10))
//	}
//
// numerator = tử số, denominator = mẫu số
// Here is the code after:
func DivByRand(numerator int, r RandNumberGenerator) int {
	// return numerator / r.randomInt(10)
	// A panic from dividing by zero. Let's fix divByRand method to prevent this:
	denominator := r.randomInt(10)
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

// We can call this in production code like this: mock_demo.DivByRand(200, mock_demo.StandardRand{}).
// In tests, though, we will instead use a mock implementation of our randNumberGenerator interface
