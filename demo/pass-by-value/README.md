# Golang is pass-by-value

Yes, Go always passes by value. In Go, every time you pass a variable to a function, the language **makes a copy of that variable** and **gives the copy to the function**.

While the "always pass-by-value" rule is absolute, the behavior you see depends on the type of data you are passing.

## 1. Simple Values (Integers, Booleans, Structs, Arrays)

When you pass these, Go copies the entire content. Any changes made inside the function only affect the local copy, not the original variable.

- Integers/Booleans: Copies the small value.
- Structs/Arrays: Copies every field or element in the structure. Large structs can be expensive to copy.

## 2. Pointers

When you pass a pointer (e.g., `*int`), Go still passes it by value. However, the "value" being copied is the memory address.

- Since **both the original and the copied pointer point to the same address**, the function can **modify the data at that address**.
- If you change the pointer itself inside the function (e.g., point it to a new address), that change won't be seen by the caller.

## 3. "Reference-like" Types (Slices, Maps, Channels)

These types often confuse beginners because they feel like they are passed by reference, but they are still technically passed by value.

- Slices: A slice is **actually a small struct** (header) containing a pointer to an underlying array, a length, and a capacity. Passing a slice copies this header, but the internal pointer still points to the same array.
- Maps and Channels: These are internally pointers to more complex header structures. Passing them copies the pointer to that structure.

## Summary of Behavior

| Type                | What is Copied?                      | Effect of Mutation in Function             |
| ------------------- | ------------------------------------ | ------------------------------------------ |
| int, string, struct | The entire data                      | No effect on caller.                       |
| `*pointer (T)`      | The memory address                   | Affects the original data.                 |
| slice               | The slice header (pointer, len, cap) | Affects existing elements; append may not. |
| map                 | The pointer to the map header        | Affects the original map.                  |

Tóm lại

- Nếu 1 function nhận kiểu value-type (int, struct, array), nó sẽ nhận một bản sao của biến đó. Thay đổi trong function sẽ không ảnh hưởng đến biến gốc.
- Nếu function nhận kiểu reference-like (slice, map, channel), nó sẽ nhận một bản sao của header chứa con trỏ đến dữ liệu thực sự. Thay đổi nội dung dữ liệu sẽ ảnh hưởng đến biến gốc, nhưng thay đổi header (như gán slice mới) sẽ không ảnh hưởng đến biến gốc.
- Nếu function nhận một pointer, nó sẽ nhận một bản sao của con trỏ (địa chỉ). Thay đổi dữ liệu tại địa chỉ đó sẽ ảnh hưởng đến biến gốc, nhưng thay đổi con trỏ (gán pointer mới) sẽ không ảnh hưởng đến biến gốc.

Ref: Google AI

# Example code

To prove a language is truly pass-by-reference, you have to try to swap the underlying address of the variable itself.

In C++, you can write a function that takes a reference and points the original variable to a totally different memory location. In Go, you simply cannot do this because the pointer itself is a copy.

## The C++ Code (True Pass-by-Reference)

In this example, the function reassign changes what ptr points to in main.

```cpp
#include <iostream>
int globalNum = 500;
// The '&' makes 'p' an alias for 'ptr' in main.// Changing 'p' changes 'ptr' itself.void reassign(int* &p) {
    p = &globalNum;
}
int main() {
    int localNum = 10;
    int* ptr = &localNum;

    std::cout << "Before: " << *ptr << std::endl; // Prints 10

    reassign(ptr);

    std::cout << "After: " << *ptr << std::endl;  // Prints 500
    return 0;
}
```

## Why Go Cannot Do This

In Go, you can change the data at an address, but you cannot change the original variable to point somewhere else by passing it into a function.
If we try to write the equivalent in Go, it fails to update the original:

```go
package main
import "fmt"
var globalNum = 500
func reassign(p *int) {
	// 'p' is a COPY of the pointer.
	// We are only changing the local copy to point to globalNum.
	p = &globalNum
}
func main() {
	localNum := 10
	ptr := &localNum

	fmt.Println("Before:", *ptr) // Prints 10

	reassign(ptr)

	fmt.Println("After:", *ptr)  // Still prints 10!
}
```

## The Difference

- C++: The function parameter `p` is `ptr`. If you move `p`, you move `ptr`.
- Go: The function parameter `p` is a new box containing the same address as `ptr`. If you move `p`, `ptr` stays exactly where it was.

To achieve the C++ result in Go, you would have to pass a pointer-to-a-pointer (`**int`), which is just passing another value (a memory address of a memory address).
Would you like to see how to use a pointer-to-a-pointer in Go to simulate that C++ behavior?

Ref: Google AI
