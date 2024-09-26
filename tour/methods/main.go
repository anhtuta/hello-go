package main

import (
	"fmt"
	"math"
)

// Methods
// Go does not have classes. However, you can define methods on types.
// Note: in Go, you cannot define methods inside a struct. Methods are defined separately from the struct definition,
// you can associate methods with a struct type by defining them with a receiver argument
type Vertex struct {
	X, Y float64
}

// A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// In this example, the Abs method has a receiver of type Vertex named v.
// Tên method: Abs, receiver: v, kiểu của receiver: Vertex, kiểu trả về: float64.
// Abs is a method of the Vertex type. This code below defines a method for the Vertex type.
func (v Vertex) Abs() float64 {
	v.X = 0 // This line does not modify the field X of the receiver v, because v is passed by value.
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// You can declare a method on non-struct types, too.
// MyFloat is a numeric type.
// Note: chỉ có thể define method cho kiểu ở cùng package với method đó.
// You cannot declare a method with a receiver whose type is defined in another package
// (which includes the built-in types such as int, float64).
// Do không thể define 1 method cho kiểu dữ liệu built-in như float64 (vì chúng ở package khác),
// nên phải tạo 1 kiểu mới từ built-in type đó.
type MyFloat float64

// Define method Abs cho kiểu MyFloat, cũng có tên là Abs, nhưng được associate với kiểu MyFloat.
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Pointer receivers
// You can declare methods with pointer receivers.
// This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)
// For example, the Scale method here is defined on *Vertex.
// Methods with pointer receivers can modify the value to which the receiver points.
// Thường thì method sẽ muốn modify receiver, nên pointer receiver sẽ phổ biến hơn value receiver.
// With a value receiver, the Scale method operates on a copy of the original Vertex value.
// With a pointer receiver, the Scale method operates on the original value.
// Note: methods with pointer receivers có thể được call từ 1 biến kiểu Vertex hoặc 1 biến kiểu *Vertex:
// var v Vertex
// v.Scale(5)  // OK
// p := &v
// p.Scale(10) // OK
func (v *Vertex) Scale(f float64) {
	// Receiver v là pointer, nên method này có thể modify giá trị của v.
	v.X = v.X * f
	v.Y = v.Y * f
}

// Tóm lại:
// Có 1 kiểu dữ liệu T, được định nghĩa bằng type T, nó có thể là struct
// Nếu muốn define method cho kiểu T nhưng không muốn modify field nào bên trong T, thì phải define method với receiver là T
// Nếu muốn define method cho kiểu T và muốn modify field bên trong T, thì phải define method với receiver là *T

// Choosing a value or pointer receiver
// There are two reasons to use a pointer receiver.
// The first is so that the method can modify the value that its receiver points to.
// The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example
// In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both

// Hàm Scale sau có logic giống với method Scale ở trên, tức là nếu gọi 1 trong 2 hàm này với 1 biến kiểu Vertex,
// thì đều có kết quả như nhau. Nhưng khi gọi hàm Scale, ta phải truyền vào 1 con trỏ tới biến kiểu Vertex:
// Scale(&v, 10) // OK
// Scale(v, 10)  // Error: cannot use v (type Vertex) as type *Vertex in argument to Scale
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println("v = ", v) // v = {3 4}
	fmt.Println(v.Abs())
	fmt.Println("v = ", v) // v = {3 4}

	fmt.Println(Vertex{X: 12, Y: 17}.Abs()) // 20.80865...

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v.Scale(10)
	fmt.Println("v after scale = ", v) // v = {30 40}
	fmt.Println(v.Abs())               // 50

	Scale(&v, 10)
	fmt.Println("v after scale = ", v) // v = {300 400}
	fmt.Println(v.Abs())               // 500

	v1 := Vertex{3, 4}
	p1 := &v1

	// Hàm Scale bắt buộc phải gọi từ 1 biến kiểu *Vertex
	(&v1).Scale(10)
	fmt.Println("v1 after scale x10 = ", v1) // v1 = {30 40}

	// &v1 hay p1 đều là con trỏ, có kiểu *Vertex, nên cả 2 đều có thể gọi method Scale
	p1.Scale(10)
	fmt.Println("v1 after scale x100 = ", v1) // v1 = {300 400}

	// Nhưng Go cho phép viết ngắn gọn hơn: thay vì dùng (&v1), có thể dùng v1
	// Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver
	v1.Scale(10)                               // ngầm hiểu là (&v1).Scale(10)
	fmt.Println("v1 after scale x1000 = ", v1) // v1 = {3000 4000}

	v2 := Vertex{3, 4}
	p2 := &v2

	// Tương tự, hàm AbsFunc nhận input kiểu Vertex, thì KHÔNG thể truyền biến kiểu *Vertex
	fmt.Println(AbsFunc(v2)) // 5
	// fmt.Println(AbsFunc(&p2)) // --> SAI, compile error!

	// v2 hay *p2 đều có kiểu Vertex, nên cả 2 đều có thể truyền vào hàm AbsFunc
	fmt.Println(AbsFunc(*p2)) // 5
	// fmt.Println(AbsFunc(p2)) // --> SAI, compile error!

	InterfaceDemo()

	EmptyInterfaceDemo()
	// doSomething() // should not be called because it is unexported (private), although we can call it literally

	TypeAssertions()

	StringersDemo()
}
