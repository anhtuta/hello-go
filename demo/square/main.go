package main

import (
	"fmt"
	"time"
)

func square(numbers []int) []int {
	var res = make([]int, len(numbers))
	for i := 0; i < len(numbers); i++ {
		res[i] = numbers[i] * numbers[i]
	}
	return res
}

type Num struct {
	idx int
	val int
}

func square1(numbers []int) []int {
	n := len(numbers)
	var res = make([]int, n)
	queue := make(chan Num, n)
	var tempIdx int
	for i := 0; i < n; i++ {
		tempIdx = i
		go func() {
			// Race condition: tempIdx always = n-1
			time.Sleep(2 * time.Second)
			temp := numbers[tempIdx] * numbers[tempIdx]
			queue <- Num{idx: tempIdx, val: temp}
		}()
	}

	for i := 0; i < n; i++ {
		num := <-queue
		res[num.idx] = num.val
	}
	return res
}

func square2(numbers []int) []int {
	n := len(numbers)
	var res = make([]int, n)
	queue := make(chan Num, n)
	var tempIdx int
	for i := 0; i < n; i++ {
		tempIdx = i
		// Most explicit/safe pattern (works clearly in all versions and styles):
		// pass the value directly into the goroutine, so no accidental sharing.
		go func(tempIdx int) {
			time.Sleep(2 * time.Second)
			temp := numbers[tempIdx] * numbers[tempIdx]
			queue <- Num{idx: tempIdx, val: temp}
		}(tempIdx)
	}

	for i := 0; i < n; i++ {
		num := <-queue
		res[num.idx] = num.val
	}
	return res
}

func square3(numbers []int) []int {
	n := len(numbers)
	var res = make([]int, n)
	queue := make(chan Num, n)

	// In modern Go (Go 1.22+), loop variables created in the loop header are per-iteration variables.
	// So each iteration gets its own i value, and each goroutine captures that iteration’s i.
	for i := 0; i < n; i++ {
		go func() {
			time.Sleep(2 * time.Second)
			temp := numbers[i] * numbers[i]
			queue <- Num{idx: i, val: temp}
		}()
	}

	for i := 0; i < n; i++ {
		num := <-queue
		res[num.idx] = num.val
	}
	return res
}

func main() {
	nums := []int{1, 9, 3, 4, 5}
	res := square2(nums)
	fmt.Println("res: ", res)
}
