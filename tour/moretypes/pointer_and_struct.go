package main

import "fmt"

func pointerAndStruct() {
	fmt.Println("\n========== Pointer and struct demo ==========")

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

	p0 := 0x14000110018      // hex
	fmt.Println("p0 = ", p0) // p0 và p đều là 1 con số, nhưng p0 là số bình thường, còn p là địa chỉ ô nhớ
	fmt.Println("p = ", p)   // p lưu trữ địa chỉ của i: 0x14000110018
	fmt.Println("*p = ", *p) // *p lấy giá trị của biến mà p trỏ tới: 21

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
}
