package slice

import (
	"fmt"
	"slices"
)

// https://gobyexample.com/slices
// In Go, every slice is backed by an underlying array.
// The slice is a descriptor containing a pointer to the array, its length, and its capacity.
// The slice itself does not store the data—only the underlying array does.
// Multiple slices can share the same underlying array.
func SliceDemo() {
	fmt.Print("\n======= Slices =======\n\n")

	// slices are typed only by the elements they contain (not the number of elements).
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// make a slice of strings of length 3 (initially zero-valued)
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	// Array doesn't have built-in append function, but slices do.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s, "len:", len(s), "cap:", cap(s))

	// Slices can also be copy’d. Here we create an empty slice c of the same length as s and copy into c from s.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c, "len:", len(c), "cap:", cap(c))

	// Slices support a “slice” operator with the syntax slice[low:high].
	// For example, this gets a slice of the elements s[2], s[3], and s[4].
	// Because l can be extended up to the end of the underlying array starting from index 2. That’s why cap(l) is 4.
	l := s[2:5]
	fmt.Println("sl1:", l, "len:", len(l), "cap:", cap(l))

	// slice l has len = 3, cap = 4, mean you can add another element to l without allocating a new array, up to its capacity
	l = append(l, "g")
	fmt.Println("sl1:", l, "len:", len(l), "cap:", cap(l)) // len = 4, cap = 4

	// When you append elements to a slice in Go and exceed its current capacity,
	// Go automatically allocates a new, larger underlying array to accommodate the new elements.
	// The new capacity is usually doubled
	l = append(l, "h")
	fmt.Println("sl1:", l, "len:", len(l), "cap:", cap(l)) // len = 5, cap = 8 (the capacity is doubled when the slice needs to grow)

	// Cannot access index 8 because the length of the slice is 5, even though the capacity is 8.
	// fmt.Println(l[cap(l)]) // panic: runtime error: index out of range [8] with length 5

	// This slices from the start of s to s[4] (exclude s[5]).
	l = s[:5]
	fmt.Println("sl2:", l, "len:", len(l), "cap:", cap(l))

	// This slices from s[2] to the end.
	l = s[2:]
	fmt.Println("sl3:", l, "len:", len(l), "cap:", cap(l))

	// declare and initialize a variable for slice in a single line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t, "len:", len(t), "cap:", cap(t))

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// The length of the inner slices can vary, unlike with multi-dimensional arrays
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// https://www.youtube.com/watch?v=RVTfPy_NELc
	var numbers [6]int = [6]int{0, 1, 2, 3, 4, 5}
	var slice []int = numbers[2:4]
	fmt.Println("slice:", slice) // [2 3]

	// re-slice the same slice variable up to its current capacity (which is 4)
	slice = slice[0:4]
	fmt.Println("slice:", slice) // [2 3 4 5]
	slice1 := numbers[0:4]
	fmt.Println("slice1:", slice1) // [0 1 2 3]

	// slice the entire underlying array
	slice2 := numbers[:]
	fmt.Println("slice2:", slice2) // [0 1 2 3 4 5]

	slice3 := slice2[:]
	numbers[2] = 100
	fmt.Println("slice2:", slice2) // [0 1 100 3 4 5]
	fmt.Println("slice3:", slice3) // [0 1 100 3 4 5]
}
