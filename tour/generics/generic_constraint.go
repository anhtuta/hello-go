package main

import "fmt"

// Ref: https://200lab.io/blog/generics-trong-golang/#s%E1%BB%AD-d%E1%BB%A5ng-constraint-generics-trong-go

// Sử dụng constraint Generics trong Go:
/*
func Min[T any](s []T) T {
	r := s[0]
	for _, v := range s[1:] {
		// invalid operation: cannot compare v < r (operator < not defined on T)
		// Lỗi này cũng dễ hiểu vì là any nên không phải value nào cũng thực hiện phép toán so sánh được
		if v < r {
			r = v
		}
	}
	return r
}
*/

// Solution 1: truyền vào thêm một hàm compare để so sánh
func Smallest1[T any](s []T, compare func(T, T) int) T {
	r := s[0]
	for _, v := range s[1:] {
		if compare(r, v) == 1 {
			r = v
		}
	}
	return r
}

// Solution 2: sử dụng constraint
type SignedInteger interface {
	int | int8 | int16 | int32 | int64
}

// Bây giờ T bắt buộc phải thuộc một trong những kiểu int | int8 | int16 | int32 | int64 thì mới dùng phép toán so sánh được.
func Smallest2[T SignedInteger](s []T) T {
	r := s[0]
	for _, v := range s[1:] {
		if r > v {
			r = v
		}
	}
	return r
}

func genericConstraint() {
	fmt.Println("\n========== Generic constraint ==========")

	arr := []int{11, 17, 9, 5, 12}

	// Solution 1
	min := Smallest1(arr, func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})
	fmt.Println("min:", min)

	// Solution 2
	min = Smallest2(arr)
	fmt.Println("min:", min)
}
