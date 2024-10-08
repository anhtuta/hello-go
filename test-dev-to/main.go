package main

import (
	"fmt"
	"net/http"
	"test-dev-to/httptest_demo"
	"test-dev-to/mock_demo"
)

func main() {
	// Mock demo
	res := mock_demo.DivByRand(200, mock_demo.StandardRand{})
	fmt.Println("200 after divided by a random number:", res)

	// byte and string demo.
	// []byte() and string() are not functions, but rather type conversions
	bytes := []byte(`{"message": "Stay slothful!"}`)
	fmt.Println("Slothful message in []byte:", bytes)
	fmt.Println("Slothful message:", string(bytes))

	// httptest demo
	http.ListenAndServe(":1123", httptest_demo.AppRouter())
}
