package closure

// https://gobyexample.com/closures
func IntSeq() func() int {
	i := 0
	return func() int {
		// what does it capture? why i doesn't need to be final?
		i++
		return i
	}
}
