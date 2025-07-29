package main

import (
	"fmt"
	"sync"
)

var a = 0
var mtx = sync.Mutex{}

func Add1() {
	a++
}

func Add2() {
	mtx.Lock()
	defer mtx.Unlock()
	a++
}

// Ref: https://viblo.asia/p/go-sync-package-6-khai-niem-chinh-ban-can-biet-Ny0VGj2pLPA
func main() {
	// tạo ra 500 goroutines cùng tăng biến a lên 1, hấu hết trường hợp, kết quả sẽ nhỏ hơn 500.
	for i := 0; i < 500; i++ {
		go Add1()
	}

	// time.Sleep(5 * time.Second)
	fmt.Println(a)

	// cải thiện ví dụ trên với sync.Mutex

}
