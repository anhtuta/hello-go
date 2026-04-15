package array

import "fmt"

// Only exported (public) functions, which start with an uppercase letter,
// can be accessed from other packages.
// In typical Go code, slices are much more common; arrays are useful in some special scenarios
func ArrayDemo() {
	fmt.Print("\n======= Arrays =======\n\n")

	// an array that will hold exactly 5 ints
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	// declare and initialize an array in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// You can also have the compiler count the number of elements for you with ...
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// If you specify the index with :, the elements in between will be zeroed
	b = [...]int{100, 3: 400, 500} // [100 0 0 400 500]
	fmt.Println("idx:", b)

	// build multi-dimensional data structures
	var twoD [2][3]int
	for i := range twoD {
		for j := range twoD[i] {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
