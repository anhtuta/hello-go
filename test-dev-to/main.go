package main

import (
	"fmt"
	"test-dev-to/sloths"
)

func main() {
	res := sloths.DivByRand(200, sloths.StandardRand{})
	fmt.Println("200 after divided by a random number:", res)
}
