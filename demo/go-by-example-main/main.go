package main

import (
	"demo/array"
	"demo/closure"
	"demo/constants"
	"demo/goroutines"
	rangeovertypes "demo/range-over-types"
	"demo/slice"
	stringsandrunes "demo/strings-and-runes"
	"fmt"
)

func main() {
	fmt.Println("My nickname is", constants.MY_NICKNAME)
	fmt.Println("My university is", constants.MY_UNIVERSITY)
	fmt.Println("My company is", constants.MY_COMPANY)
	fmt.Println("Another constant is", constants.ANOTHER_CONSTANT)

	array.ArrayDemo()

	slice.SliceDemo()

	fmt.Print("\n======= Closures =======\n\n")
	nextInt := closure.IntSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	newInts := closure.IntSeq()
	fmt.Println(newInts())

	rangeovertypes.RangeDemo()

	stringsandrunes.RuneDemo()

	goroutines.GoroutineDemo()
	goroutines.ChannelDemo()
	goroutines.ChannelSyncDemo()
	goroutines.ChannelDirectionDemo()
	goroutines.SelectDemo()
	goroutines.TimeoutDemo()
	goroutines.ChannelNonBlockingOpsDemo()
	goroutines.ChannelCloseDemo()
}
