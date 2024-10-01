package main

import "fmt"

// Structs
// A struct is a collection of fields.
type Vertex struct {
	X int
	Y int
}

func main() {
	i, j := 42, 2701

	// Pointers
	// Go has pointers. A pointer holds the memory address of a value.
	// Giống như C/C++, ta dùng dấu * để khai báo con trỏ, và dùng dấu & để lấy địa chỉ của biến.
	// var p *int
	// p = &i       // point to i: p sẽ trỏ tới địa chỉ của biến i, tức là p lưu trữ địa chỉ của i.
	p := &i         // point to i, but shorter
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer: *p = 21 tương đương với i = 21
	fmt.Println(i)  // see the new value of i

	p = &j                                     // point to j
	*p = *p / 37                               // divide j through the pointer
	fmt.Println("j after modifying via p:", j) // see the new value of j: 2701 / 37 = 73

	fmt.Println(Vertex{1, 2})

	// Struct fields are accessed using a dot
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	fmt.Println(v)

	// Struct fields can be accessed through a struct pointer.
	// To access the field X of a struct when we have the struct pointer p we could write (*p).X.
	// However, that notation is cumbersome, so the language permits us instead to write just p.X,
	// without the explicit dereference
	// (Dereferencing is the process of accessing the value that a pointer points to.
	// In Go, you use the * operator to dereference a pointer)
	v1 := Vertex{1, 2}
	p1 := &v1  // var p1 *Vertex: p1 is a pointer to v1 (a struct)
	p1.X = 1e9 // a shorten form of: (*p1).X = 1e9
	fmt.Println(v1)

	// Struct Literals: tạo mới một struct = cách list các giá trị của các fields.
	v2 := Vertex{1, 2}  // has type Vertex
	v3 := Vertex{X: 1}  // Y:0 is implicit
	v4 := Vertex{}      // X:0 and Y:0
	p2 := &Vertex{1, 2} // has type *Vertex
	p3 := Vertex{1, 2}  // has type Vertex, NOT a pointer to Vertex
	fmt.Println(v2, v3, v4)

	// p2 points to a Vertex struct allocated in memory. Modifying p2 modifies the struct it points to.
	// p3 is a value type, holding the struct data directly. Modifying p3 only affects this specific instance.
	fmt.Println("p2:", *p2) // Dereference p2 to get the Vertex value
	fmt.Println("p3:", p3)

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
	// Hiểu đơn giản nó là 1 mảng con của 1 mảng khác.
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
	fmt.Println(s1)

	// Slice defaults
	// When slicing, you may omit the high or low bounds to use their defaults instead
	s2 := []int{2, 3, 5, 7, 11, 13}

	s = s2[1:4]
	fmt.Println(s) // [3 5 7]

	s = s[:2]
	fmt.Println(s) // [3 5]

	s = s[1:]
	fmt.Println(s) // [5]

	//////// So sánh slice với array (Copilot)
	fmt.Println("\n========== So sánh slice với array ==========")
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

	// Not done yet, continue at https://go.dev/tour/moretypes/11
}
