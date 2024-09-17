package main

// Tên package là phần tử cuối cùng của import path, ex: với package "math/rand",
// ta có thể truy cập các file bắt đầu bằng package rand.
// Có thể dùng nhiều lệnh import cho từng package, hoặc nhóm chúng lại bằng cặp (),
// cách nhóm này gọi là "factored" import statement.
import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

// Shorten: func add(x, y int) int { ... }
func add(x int, y int) int {
	return x + y
}

// Multiple results
func swap(x, y string) (string, string) {
	return y, x
}

// Named return values
// Go's return values may be named. If so, they are treated as variables defined at the top of the function.
// These names should be used to document the meaning of the return values.
// A return statement without arguments returns the named return values. This is known as a "naked" return.
// Naked return chỉ nên dùng với hàm ngắn, vì khó hiểu!
// Hàm sau tách 1 số thành 2 số. x và ý là 2 biến được khai báo ở đầu hàm và được trả về khi hàm kết thúc.
// Do đó lệnh return không cần ghi rõ return x, y.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// Variables
var c, python, java bool
var x1, x2 int = 1995, 11

// Không thể dùng := ở ngoài hàm
// x3 := 2

// Go's basic types are
// bool
// string
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte // alias for uint8
// rune // alias for int32, represents a Unicode code point
// float32 float64
// complex64 complex128

// variable declarations may be "factored" into blocks, as with import statements.
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// Constants cannot be declared using the := syntax. We must use the const keyword.
const Pi = 3.14

// Numeric constants are high-precision values
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))

	// 1 tên nếu được export thì nó thường được viết hoa.
	fmt.Println("Pi is", math.Pi)

	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
	fmt.Println(split(100))

	var i int
	fmt.Println(i, c, python, java)
	fmt.Println(x1, x2, c, python, java)

	// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type
	// Variable x3 is declared and initialized with value 2, and its type is inferred from the right-hand side.
	x3 := 5
	fmt.Println(x3)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Zero values
	// Biến không được khởi tạo sẽ có giá trị zero value của kiểu dữ liệu đó.
	// Zero values có thể là: 0 cho số, false cho bool, "" cho string.
	var x4 int
	var x5 float64
	var x6 bool
	var x7 string
	fmt.Printf("%v %v %v %q\n", x4, x5, x6, x7)

	// Type conversions
	// The expression T(v) converts the value v to the type T.
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	// Unlike in C, in Go assignment between items of different type requires an explicit conversion
	// this is wrong:
	// var f1 float64 = x

	// Constants
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
