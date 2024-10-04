package main

import "fmt"

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

	fmt.Println("\n========== So sánh slice với array (Copilot) ==========")
	// The primary difference between creating an array and a slice in Go is that with a slice,
	// you do not need to specify the size. Here are the key points:
	// Array: Fixed Size: The size of the array is specified at the time of creation and cannot be changed.
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)
	// Slice: Dynamic Size: The size of the slice can grow and shrink dynamically.
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice)
	slice = append(slice, 6, 7, 8)
	fmt.Println("Slice after append:", slice)
}
