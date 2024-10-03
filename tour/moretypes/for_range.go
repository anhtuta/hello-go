package main

import "fmt"

type Student struct {
	Id      int
	Name    string
	Address string
}

func (s Student) String() string {
	return fmt.Sprintf("Id: %d, Name: %s, Address: %s", s.Id, s.Name, s.Address)
}

func forRange() {

	fmt.Println("\n========== For range demo (Copilot) ==========")
	// When using for range, the second value in the loop is a copy of the element, not a reference to the original element.
	// This means that modifying the second value inside the loop does not affect the original element in the array or slice
	arr1 := [5]int{1, 2, 3, 4, 5}
	for i, v := range arr1 {
		v = v * 2
		fmt.Printf("Index: %d, Value: %d, Array: %v\n", i, v, arr1)
	}
	fmt.Println("Final Array:", arr1)
	// If you need to modify the original elements of the array or slice,
	// you should use the index to access and modify the elements directly:
	for i := range arr1 {
		arr1[i] = arr1[i] * 2
	}
	fmt.Println("Final Array after modifying:", arr1)

	// Nếu for range với 1 mảng các phần tử có kiểu struct, thì kết quả sẽ giống như trên.
	var students1 []Student
	students1 = append(students1, Student{Id: 1, Name: "Alice", Address: "123 Main St"})
	students1 = append(students1, Student{Id: 2, Name: "Bob", Address: "456 Elm St"})
	students1 = append(students1, Student{Id: 3, Name: "Charlie", Address: "789 Oak St"})
	for _, s := range students1 {
		// modifying the copy of the struct, not the original struct in the slice.
		s.Address = "New Address"
	}
	fmt.Println("students1:", students1)

	// Nhưng nếu for range với 1 mảng các phần tử có kiểu CON TRỎ của 1 struct, thì kết quả sẽ KHÁC.
	// However, if students is a slice of pointers to structs, s will be a copy of the pointer, and modifying s.Address
	// will affect the original struct because both s and the original element in the slice point to the same struct
	var students2 []*Student
	students2 = append(students2, &Student{Id: 1, Name: "Alice", Address: "123 Main St"})
	students2 = append(students2, &Student{Id: 2, Name: "Bob", Address: "456 Elm St"})
	students2 = append(students2, &Student{Id: 3, Name: "Charlie", Address: "789 Oak St"})
	for _, s := range students2 {
		// modifying the original struct in the slice
		s.Address = "New Address"
	}
	fmt.Println("students2:", students2)
}
