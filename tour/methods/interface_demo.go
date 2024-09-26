package main

import "fmt"

// Interfaces
// An interface type is defined as a set of method signatures.
// A value of interface type can hold any value that implements those methods.

// Go khác Java, C++ ở chỗ: nó không có class, do đó không thể tạo class để implement 1 interface.
// Muốn implement 1 interface, phải implement các method của interface đó cho 1 kiểu dữ liệu nào đó.

// Interfaces are implemented implicitly
// A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.
// Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement

// Geometric = hình học (adj)
type Geometric interface {
	Perimeter() float64
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

// Implement methods of Geometric interface for Rectangle type.
// This method means type Rectangle implements the interface Geometric, but we don't need to explicitly declare that it does so.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

// Implement methods of Geometric interface for Circle type.
func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func (c *Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Interface values
// Under the hood, interface values can be thought of as a tuple of a value and a concrete type:
// (value, type)
// An interface value holds a value of a specific underlying concrete type.
// Calling a method on an interface value executes the method of the same name on its underlying type.
// Check value và type của 1 interface: fmt.Printf("(%v, %T)\n", i, i)

// Nil interface values
// A nil interface value holds neither value nor concrete type.
// Calling a method on a nil interface is a run-time error because there is no type inside the interface
// tuple to indicate which concrete method to call.

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func InterfaceDemo() {
	fmt.Println("\n========== InterfaceDemo ==========")
	rect := Rectangle{Width: 3, Height: 4}
	fmt.Println("Rectangle perimeter = ", rect.Perimeter())
	fmt.Println("Rectangle area = ", rect.Area())

	circle := Circle{Radius: 5}
	fmt.Println("Circle perimeter = ", circle.Perimeter())
	fmt.Println("Circle area = ", circle.Area())

	var g Geometric = rect
	fmt.Println("Rectangle perimeter = ", g.Perimeter())
	fmt.Println("Rectangle area = ", g.Area())
	fmt.Printf("(value, type) of g: (%v, %T)\n", g, g)

	g = &circle
	fmt.Println("Circle perimeter = ", g.Perimeter())
	fmt.Println("Circle area = ", g.Area())
	fmt.Printf("(value, type) of g: (%v, %T)\n", g, g)

	// The empty interface
	// The interface type that specifies zero methods is known as the empty interface:
	// interface{}
	// An empty interface may hold values of any type. (Every type implements at least zero methods.)
	// Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}

	fmt.Println("\n========== Empty interface ==========")
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}
