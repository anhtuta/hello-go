package main

import (
	"fmt"
	"strings"
)

func arrayAndSlice() {
	fmt.Println("\n========== Array and slice demo ==========")

	// Arrays
	// The type [n]T is an array of n values of type T.
	// An array's length is part of its type, so arrays cannot be resized
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Slices
	// Hiểu đơn giản nó là 1 mảng con của 1 mảng khác, hoặc 1 mảng động có thể resize.
	var s []int = primes[1:4] // lấy từ index = 1 đến index = 3 (không lấy index = 4)
	fmt.Println("s = ", s)    // [3 5 7]

	// Cú pháp: a[low : high] (a low and high bound, low is included, high is excluded)
	// Có thể bỏ qua low hoặc high bound để sử dụng default value.
	// low bound mặc định = 0, high bound mặc định = len(a) (length of the slice)

	// Slices are like references to arrays
	// Nó không lưu data, mà tham chiếu đến 1 phần của 1 mảng khác.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a1 := names[0:2]
	b1 := names[1:4]
	fmt.Println("a1, b1 = ", a1, b1)

	b1[0] = "XXX" // changing b1[0] will change a1[1] and names[1]
	fmt.Println("After changing b1[0], now a1, b1 = ", a1, b1)
	fmt.Println("After changing b1[0], now names = ", names)

	// Slice literals
	// A slice literal is like an array literal without the length.
	// Tạo mới 1 slice mà không cần phải chỉ rõ length, nhưng phải list các giá trị.
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	q = append(q, 17, 19)
	fmt.Println("q after append:", q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	// Tương đương với: var s1 []struct{i int; b bool}; s1 = ...
	s1 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println("s1:", s1) // [{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]

	// Slice defaults
	// When slicing, you may omit the high or low bounds to use their defaults instead
	s2 := []int{2, 3, 5, 7, 11, 13}

	s = s2[1:4]
	fmt.Println(s) // [3 5 7]

	s = s[:2]
	fmt.Println(s) // [3 5]

	s = s[1:]
	fmt.Println(s) // [5]

	fmt.Println("\n========== Slice length and capacity ==========")
	// Slice length and capacity
	// A slice has both a length and a capacity.
	// The length of a slice is the number of elements it contains.
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	// (capacity = số lượng phần tử ở trong mảng gốc)
	s3 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s3) // len=6 cap=6 [2 3 5 7 11 13]

	// Slice the slice to give it zero length.
	// Lệnh sau sẽ tạo mới 1 slice = cách bắt đầu từ phần tử 0 của slice cũ, và không lấy phần tử nào.
	// Underlying array không bị thay đổi: [2, 3, 5, 7, 11, 13]
	// Do đó vẫn có thể mở rộng slice sau này và lấy các phần tử cũ.
	s3 = s3[:0]
	printSlice(s3) // len=0 cap=6 []

	// Extend its length.
	// Do underlying array không thay đổi, nên ta có thể mở rộng length của slice.
	s3 = s3[:4]
	printSlice(s3) // len=4 cap=6 [2 3 5 7]

	// Drop its first two values: creating a new slice that starts from the third element of the original slice
	// and includes all subsequent elements.
	// Lệnh sau tạo mới 1 slice = cách bắt đầu từ phần tử thứ 2 của slice cũ, và lấy tất cả các phần tử sau đó.
	// Underlying array của s3 lúc này khác với underlying array của s3 ban đầu, length và capacity cũng khác nhau.
	// KHÔNG thể lấy lại các phần tử cũ của s3 nữa.
	s3 = s3[2:]
	printSlice(s3) // len=2 cap=4 [5 7]

	// Nil slices
	// The zero value of a slice is nil.
	// A nil slice has a length and capacity of 0 and has no underlying array.
	var s4 []int
	fmt.Println(s4, len(s4), cap(s4)) // [] 0 0
	if s4 == nil {
		fmt.Println("nil!")
	}

	fmt.Println("\n========== Creating a slice with make ==========")
	// Creating a slice with make
	// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
	// The make function allocates a zeroed array and returns a slice that refers to that array.
	// make([]T, len): creates a slice of type T with a length of len.
	// make([]T, len, cap): creates a slice of type T with a length of len and a capacity of cap.
	a2 := make([]int, 5)
	printSlice(a2) // len=5 cap=5 [0 0 0 0 0]

	// Slices of slices: giống như mảng 2 chiều
	// Slices can contain any type, including other slices
	// Create a tic-tac-toe board
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	fmt.Println("\n========== So sánh slice với array (Copilot) ==========")
	// The primary difference between creating an array and a slice in Go is that with a slice,
	// you do not need to specify the size. Here are the key points:
	// Array: Fixed Size: The size of the array is specified at the time of creation and cannot be changed.
	arr := [5]int{1, 2, 3, 4, 5}
	// arr = append(arr, 6) // Error: first argument to append must be slice; have [5]int
	fmt.Println("Array:", arr)
	// Slice: Dynamic Size: The size of the slice can grow and shrink dynamically.
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice) // [1 2 3 4 5]
	slice = append(slice, 6, 7, 8)
	fmt.Println("Slice after append:", slice) // [1 2 3 4 5 6 7 8]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
