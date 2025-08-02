package main

import (
	"fmt"
	"sync"
	"time"
)

// Ta có thể thấy hai hàm produce và consume có thể trao đổi dữ liệu qua lại bằng channel,
// do đó việc gửi nhận dữ liệu vô cùng dễ dàng và an toàn, khỏi phải dùng lock,
// wait hay notify trong Java, khỏe re luôn
// Ref: https://batnamv.medium.com/go-concurrency-d%C3%A0nh-cho-java-developers-c7709f1f8752
func produce(value int, in chan<- int, wg sync.WaitGroup) {
	wg.Add(1)
	in <- value
	defer wg.Done()
}

func consume(out <-chan int) {
	for {
		val := <-out
		fmt.Println(val)
	}
}

func main() {
	var wg sync.WaitGroup

	ch := make(chan int, 100)
	go consume(ch)

	for i := 0; i < 50; i++ {
		go produce(i, ch, wg)
		time.Sleep(time.Second / 2)
	}

	defer wg.Wait()
}
