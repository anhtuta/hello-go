package goroutines

import "fmt"

func ChannelDemo() {
	fmt.Print("\n======= Channels =======\n\n")

	// Channels are typed by the values they convey
	messages := make(chan string)

	// Send a value into a channel
	go func() { messages <- "ping" }()

	// receives a value from the channel
	msg := <-messages
	fmt.Println("msg:", msg)

	// messages <- "ping2" // fatal error: all goroutines are asleep - deadlock!
	// fmt.Println("msg:", <-messages)

	fmt.Print("\n======= Channel Buffering =======\n\n")

	// Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive
	messages1 := make(chan string, 2)

	messages1 <- "buffered"
	messages1 <- "channel"

	// This will cause a deadlock because the buffer is full and there is no concurrent receive to free up space in the buffer
	// fatal error: all goroutines are asleep - deadlock!
	// messages1 <- "test"

	fmt.Println(<-messages1)
	fmt.Println(<-messages1)
}
