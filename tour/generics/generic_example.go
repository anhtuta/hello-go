package main

import "fmt"

// Ref: https://200lab.io/blog/generics-trong-golang/

// Sử dụng Generics cho function trong Go:
// K và V sẽ là 2 kiểu dữ liệu generics cho hàm Map. Chúng ta chưa rõ nó là gì nhưng chúng ta muốn
// mọi kiểu dữ liệu đều có thể dùng được nên cả 2 đều là any. (Hàm map này giống array.map trong JS)
// Ở phần định nghĩa tham số truyền vào chúng ta sẽ cần một mảng K: []K và một hàm transform nhận kiểu K
// và trả về kiểu V. Hàm này là để người gọi muốn làm gì với mỗi item K thì làm, miễn return V là okie.
// Cuối cùng hàm Map phải trả về một mảng mới: []V
func Map[K, V any](s []K, transform func(K) V) []V {
	rs := make([]V, len(s))
	for i, v := range s {
		rs[i] = transform(v)
	}
	return rs
}

// Sử dụng Generics định nghĩa các cấu trúc dữ liệu trong Go:
type Vector[T any] []T

type LinkedList[T any] struct {
	Next *LinkedList[T]
	Val  T
}

type Pair[T1, T2 any] struct {
	V1 T1
	V2 T2
}

type Tuple[T1, T2, T3 any] struct {
	V1 T1
	V2 T2
	V3 T3
}

func genericExample() {

	fmt.Println("\n========== Generic example ==========")

	// Chạy thử với 2 trường hợp mảng int và mảng string
	arr := []int{1, 2, 3}
	resultArr := Map(arr, func(v int) int { return v * 2 })
	fmt.Println(resultArr)

	arr2 := []string{"a", "b", "c"}
	resultArr2 := Map(arr2, func(v string) string { return v + v })
	fmt.Println(resultArr2)
}
