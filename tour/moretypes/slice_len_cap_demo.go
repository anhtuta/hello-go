package main

import "fmt"

func sliceLenCapDemo() {
	fmt.Println("\n========== Slice length and capacity demo ==========")

	s3 := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("Original slice:", s3)
	fmt.Println("Length:", len(s3), "Capacity:", cap(s3))

	// Slice to zero length
	s3ZeroLength := s3[:0]
	fmt.Println("After slicing to zero length:", s3ZeroLength)
	fmt.Println("Length:", len(s3ZeroLength), "Capacity:", cap(s3ZeroLength))
	fmt.Println("Underlying array after s3[:0]:", s3)

	// Reset to original slice for next operation
	// s3 = []int{2, 3, 5, 7, 11, 13}
	s3 = s3[:cap(s3)]

	// Slice to drop the first two elements
	s3DropFirstTwo := s3[2:]
	fmt.Println("After slicing to drop first two elements:", s3DropFirstTwo)
	fmt.Println("Length:", len(s3DropFirstTwo), "Capacity:", cap(s3DropFirstTwo))
	fmt.Println("Underlying array after s3[2:]:", s3)
}

// So sánh s3[:0] và s3[2:]
func sliceLenCapDemo1() {
	fmt.Println("\n========== Slice length and capacity demo 1 ==========")

	s3 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s3) // len=6 cap=6 [2 3 5 7 11 13]

	// s3[:0]:
	// Creates a new slice with length 0 but retains the original capacity.
	// You can restore the slice to its full capacity because the underlying array remains unchanged
	// Nếu như không thay đổi low bound, thì slice vẫn giữ nguyên capacity.
	// Underlying array của s3 và s3[:0] vẫn giống nhau.
	s3 = s3[:0]
	printSlice(s3) // len=0 cap=6 []

	s3 = s3[:cap(s3)]
	printSlice(s3) // len=6 cap=6 [2 3 5 7 11 13]

	// s3[2:]:
	// Creates a new slice starting from the third element, effectively dropping the first two elements.
	// The capacity of the new slice is reduced, and you cannot restore the slice to include the first two elements
	// Nếu thay đổi low bound, thì slice sẽ giảm capacity, các phần tử trước đó không thể khôi phục.
	// Underlying array của s3 và s3[2:] không giống nhau.??? KHÔNG rõ đoạn này lắm, vì cap của 2 bọn nó khác nhau.
	// Update: có vẻ như Underlying array của chúng là giống nhau đó! s3 và s3[2:] chỉ khác nhau length và capacity.
	s3 = s3[2:]
	printSlice(s3) // len=4 cap=4 [5 7 11 13]

	s3 = s3[0:]
	printSlice(s3) // len=4 cap=4 [5 7 11 13]
}

func sliceLenCapDemo2() {
	fmt.Println("\n========== Slice length and capacity demo 2 ==========")

	s3 := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("Original slice:", s3)
	fmt.Println("Length:", len(s3), "Capacity:", cap(s3))
	fmt.Printf("Underlying array: %p\n", &s3[0])

	// Slice to zero length
	s3 = s3[:0]
	fmt.Println("After slicing to zero length:", s3)
	fmt.Println("Length:", len(s3), "Capacity:", cap(s3))
	fmt.Printf("Underlying array after s3[:0]: %p\n", &s3[:cap(s3)][0])

	// Reset to original slice for next operation
	// s3 = []int{2, 3, 5, 7, 11, 13}
	s3 = s3[:cap(s3)]

	// Slice to drop the first two elements
	s3 = s3[2:]
	fmt.Println("After slicing to drop first two elements:", s3)
	fmt.Println("Length:", len(s3), "Capacity:", cap(s3))
	fmt.Printf("Underlying array after s3[2:]: %p\n", &s3[0])

	// Now, cannot restore s3 to its full capacity because the underlying array has been changed???
}
