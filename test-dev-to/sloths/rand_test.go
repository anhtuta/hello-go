package sloths

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

// Nondeterministic code: code that you don't have total control over the code's logic and output.
// Example:
// Code that uses pseudorandom number generators, like math/rand
// Web API calls and their response payloads, or errors
// The current time returned by time.Now
// One tool that helps with testing nondeterministic code, is the mock package of the popular Testify testing framework

// Here's the steps to making it testable with Testify Mock:
// 1. Take the nondeterministic piece of functionality and wrap it in a Go interface type.
// 2. Write an implementation of the interface that uses Testify Mock.
// 3. In the tests, use the mock implementation to select deterministic results of the functions your interface calls.

// Let's try it out on divByRand, step by step!
// 1. Take the nondeterministic function calls and wrap them in an interface: check in rand.go
// 2. Write an implementation of the interface that uses Testify Mock
// First, set up our mock implementation by embedding a Mock into a struct.
// By embedding a Mock, now the mockRand type has the methods for registering an API call that you expect to happen in the tests
type mockRand struct {
	mock.Mock
}

func newMockRand() *mockRand { return &mockRand{} }

// Let mockRand implement the RandNumberGenerator interface (which contains method that we want to mock)
func (m *mockRand) randomInt(max int) int {
	// we call Mock.Called to record that it was called with the value passed in for max.
	args := m.Called(max)

	// m.Called returns an Arguments object, which contains the return value(s) that we specified in the test to be
	// returned if the function gets the given value of max. For example, if we called m.On("randomInt", 20).Return(7)
	// in the test, m.Called(20) would return an Arguments object holding the return value 7.
	// Nếu setup m.On("randomInt", 20).Return(7, 8) thì args.Int(0) sẽ trả về 7, args.Int(1) sẽ trả về 8.
	// if you have a return value of another type, you would retrieve it with the Arguments.Get method, which returns
	// an interface{} you would then convert to the type you expect to get back
	return args.Int(0)
}

// 3. In the tests, use the mock implementation and feed in results the function returns.
func TestDivByRand(t *testing.T) {
	// Get our mockRand:
	// m là 1 biến có kiểu mockRand, biến này implement RandNumberGenerator interface và embed mock.Mock,
	// do đó m có thể gọi stub method của RandNumberGenerator, và có thể gọi các method của mock.Mock
	// để assert và verify
	m := newMockRand()

	// Stub:
	// specify our return value. Since the code in divByRand passes MAX_RANDOM_NUMBER into randomInt, we pass
	// MAX_RANDOM_NUMBER in as the argument to go with randomInt, and specify that we want the method to return 6.
	m.On("randomInt", MAX_RANDOM_NUMBER).Return(6)

	// Test:
	// now run divByRand and assert that we got back the
	// return value we expected, just like in a Go test that
	// doesn't use Testify Mock.
	// quotient = thương số, = numerator/denominator =  = tử số/mấu số = 30/6 = 5
	quotient := DivByRand(30, m)
	if quotient != 5 {
		t.Errorf("expected quotient to be 5, got %d", quotient)
	}

	// Verify:
	// check that randomInt was called with the number MAX_RANDOM_NUMBER;
	// if not then the test fails
	m.AssertCalled(t, "randomInt", MAX_RANDOM_NUMBER)
}

func TestDivByRand_DivideByZero(t *testing.T) {
	m := newMockRand()
	m.On("randomInt", MAX_RANDOM_NUMBER).Return(0)

	quotient := DivByRand(30, m)
	if quotient != 0 {
		t.Errorf("expected quotient to be 0, got %d", quotient)
	}
}
