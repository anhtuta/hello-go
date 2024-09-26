package main

import "fmt"

// Stringers
// One of the most ubiquitous interfaces is Stringer defined by the fmt package.
// type Stringer interface {
// 	String() string
// }
// A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values
// Hàm String() của Stringer giống như hàm toString() trong Java.
// 1 kiểu dữ liệu implement interface Stringer giống 1 class override hàm toString() trong Java.

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// defines a new type named IPAddr which is an array of 4 bytes.
// Note: In Go, when you write `type A B`, it means that you are defining a new type A that has the same underlying type as B.
// Ko giống như Java, it does not mean that A is "extended" from B in the sense of inheritance.
type IPAddr [4]byte

// Exercise: Stringers
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func StringersDemo() {
	fmt.Println("\n========== StringersDemo ==========")
	a := Person{"Arthur Dent", 42}
	z := Person{"Monkey D. Luffy", 20}
	fmt.Println(a)
	fmt.Println(z)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	ip := IPAddr{192, 168, 1, 1}
	fmt.Println(ip)
}
