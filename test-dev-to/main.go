package main

import (
	"fmt"
	"test-dev-to/mock_demo"
)

func main() {
	res := mock_demo.DivByRand(200, mock_demo.StandardRand{})
	fmt.Println("200 after divided by a random number:", res)
}
