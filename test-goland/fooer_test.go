package main

// Unit tests in Go are located in the same package (that is, the same folder) as the tested function.
// By convention, if your function is in the file fooer.go file, then the unit test for that function
// is in the file fooer_test.go.

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFooer(t *testing.T) {
	result := Fooer(3)
	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
	}
}

// Use a table-driven approach to help reduce repetition. As the name suggests, this involves
// organizing a test case as a table that contains the inputs and the desired outputs
func TestFooerTableDriven(t *testing.T) {
	// Defining the columns of the table
	var tests = []struct {
		name  string
		input int
		want  string
	}{
		// The table itself. Each row of the table lists a test case to execute
		{"9 should be Foo", 9, "Foo"},
		{"3 should be Foo", 3, "Foo"},
		{"1 is not Foo", 1, "1"},
		{"0 should be Foo", 0, "Foo"},
	}

	// The execution loop calls t.Run(), which defines a sub-test. As a result,
	// each row of the table defines a sub-test named [NameOfTheFunction]/[NameOfTheSubTest].
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := Fooer(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}

	// This way of writing tests is very popular, and considered the canonical way to write unit tests in Go
	// (Đây được coi là cách viết unit test chuẩn trong Go)
}

// The Testing Package
// The testing package plays a pivotal role in Go testing. It enables developers to create unit tests with
// different types of test functions. The testing.T type offers methods to control test execution, such as running
// tests in parallel with Parallel(), skipping tests with Skip(), and calling a test teardown function with Cleanup().

// The testing.T type provides various practical tools to interact with the test workflow, including t.Errorf(),
// which prints out an error message and sets the test as failed.
// Note: t.Error* does not stop the execution of the test. Instead, all encountered errors will be reported once
// the test is completed. Sometimes it makes more sense to fail the execution; in that case, you should use t.Fatal*
// t.Errorf() sẽ in ra lỗi và đánh dấu test case đó là failed, nhưng test case vẫn tiếp tục chạy,
// nếu muốn dừng thì dùng t.Fatalf()

func TestFooer2(t *testing.T) {
	input := 3
	result := Fooer(3)
	t.Logf("The input was %d, and result was %s", input, result)
	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
	}
	//t.Fatalf("Stop the test now, we have seen enough")
	// This will be skipped, because t.Fatalf has already terminated this test.
	//t.Error("This won't be executed")
}

// Running Parallel Tests
// By default, tests are run sequentially; the method Parallel() signals that a test should be run in parallel.
// All tests calling this function will be executed in parallel. go test handles parallel tests by pausing each test
// that calls t.Parallel(), and then resuming them in parallel when all non-parallel tests have been completed.
// The GOMAXPROCS environment defines how many tests can run in parallel at one time, and by default this number
// is equal to the number of CPUs.
func TestFooerParallel(t *testing.T) {
	t.Run("Test 3 in Parallel", func(t *testing.T) {
		t.Parallel()
		result := Fooer(3)
		if result != "Foo" {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
		}
	})
	t.Run("Test 7 in Parallel", func(t *testing.T) {
		t.Parallel()
		result := Fooer(7)
		if result != "7" {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result, "7")
		}
	})
}

// Skipping Tests
// go test accepts a flag called -test.short that is intended to run a “fast” test. You need to use a combination
// of testing.Short(), which is set to true when -short is used, and t.Skip(), as illustrated below.
// This test will be executed if you run "go test -v", but will be skipped if you run "go test -v -test.short"
func TestFooerSkipped(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	result := Fooer(3)
	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
	}
}

// Writing Fuzz Tests
// Fuzz testing is an exciting testing technique in which random input is used to discover bugs or edge cases.
// Go’s fuzzing algorithm is smart because it will try to cover as many statements in your code as possible
// by generating many new input combinations.
// The goal of the fuzz test is not to validate the output of the function, but instead to use unexpected inputs
// to find potential edge cases. By default, fuzzing will run indefinitely, as long as there isn’t a failure or
// user doesn't stop it (by ctrl+C).
// To run a fuzz test, you need to use the -fuzz flag with the go test command: go test -fuzz FuzzFooer
// To stop it: go test -fuzz FuzzFooer -fuzztime=20s, or press ctrl+C
func FuzzFooer(f *testing.F) {
	f.Add(3)
	f.Fuzz(func(t *testing.T, a int) {
		Fooer(a)
	})
}

// The Testify Package
// Testify is a testing framework providing many tools for testing Go code. There is a considerable debate
// in the Go community about whether you should use Testify or just the standard library.
// Testify provides assert functions and mocks, which are similar to traditional testing frameworks,
// like JUnit for Java or Jasmine for NodeJS.
func TestFooerWithTestify(t *testing.T) {
	// assert equality
	assert.Equal(t, "Foo", Fooer(0), "0 is divisible by 3, should return Foo")
	assert.Equal(t, "1", Fooer(1), "1 is not divisible by 3, should not return 1")

	// assert inequality
	assert.NotEqual(t, "Foo", Fooer(1), "1 is not divisible by 3, should not return Foo")
}

// Testify provides two packages, require and assert. The require package will stop execution if
// there is a test failure, which helps you fail fast. assert lets you collect information,
// but accumulate the results of assertions.
func TestMapWithTestify(t *testing.T) {
	// require equality
	// require.Equal(t, map[int]string{1: "1", 2: "2"}, map[int]string{1: "1", 2: "333"})
	// all the lines below the first require assertion will be skipped. Uncomment it to see the result

	// assert equality
	assert.Equal(t, map[int]string{1: "1", 2: "2"}, map[int]string{1: "1", 2: "2"})
}

// Wrapping Up
// Go offers numerous tools out of the box to test your application. You can write any test with the standard library,
// and Testify and its “better” assertion function and mock capability offer optional additional functionality.

// The testing package offers three testing modes: regular tests (testing.T), benchmark tests (testing.B),
// and fuzz tests (testing.F). Setting any type of test is very simple. The testing package also offers many
// helper functions that will help you write better and cleaner tests. Spend some time exploring the library
// before jumping into testing.
