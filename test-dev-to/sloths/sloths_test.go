package sloths

// All test files in Go have a name that ends with the _test.go suffix,
// which tells the go test command line tool where to find the test code for your Go project

import (
	"testing"
)

type isSlothfulTestCase struct {
	testName string
	str      string
	expected bool
}

func TestIsSlothful(t *testing.T) {
	// In JavaScript, a test that we expect IsSlothful to return false would look something like this
	// expect(IsSlothful('hello, world!')).toBeFalse();

	// In Go, a test fails if its `testing.T` is made to fail. We use the t.Error function to report a test failure.
	// There's three other ways to make a testing.T fail: t.Fail, t.FailNow, and t.Fatal.
	if IsSlothful("hello, world!") {
		t.Error("hello, world! is not supposed to be slothful")
	}

	if !IsSlothful("hello, slothful world!") {
		t.Error("hello, slothful world! is supposed to be slothful")
	}

	if !IsSlothful("Sloths rule!") {
		t.Error("Sloths rule! is supposed to be slothful")
	}
}

func TestIsSlothful_TableTesting(t *testing.T) {
	var isSlothfulTestCases = []isSlothfulTestCase{{
		testName: "string with nothing slothful isn't slothful",
		str:      "hello, world!",
		expected: false,
	}, {
		testName: `string with the substring "sloth" is slothful`,
		str:      "hello, slothful world!",
		expected: true,
	}, {
		testName: `checking for the word "sloth" is case-insensitive`,
		str:      "Sloths rule!",
		expected: true,
	}, {
		testName: "strings with the 🌺 emoji are normally slothful",
		str:      "Nothing like an iced hibiscus tea! 🧊🌺",
		expected: true,
	}, {
		testName: "the 🏎️ emoji negates the 🌺 emoji's slothfulness",
		str:      "Get your 🌺 flowers! They're going fast! 🏎️",
		expected: false,
	}}
	for _, c := range isSlothfulTestCases {
		// Since Go 1.7, if inside your Go test, you run the t.Run(string, func(*testing.T)) method,
		// your testing.T will make a sub-test within your test.
		t.Run(c.testName, func(t *testing.T) {
			assertIsSlothful(t, c.str, c.expected)
		})
	}
}

func assertIsSlothful(t *testing.T, s string, expected bool) {
	if IsSlothful(s) != expected {
		if expected {
			t.Errorf("%s is supposed to be slothful", s)
		} else {
			t.Errorf("%s is not supposed to be slothful", s)
		}
	}
}

// Outside the standard library, a lot of assertions, like whether two values are equal to each other,
// or whether a given value is nil, or whether a function errored, are so common...
// --> we can use Testify
