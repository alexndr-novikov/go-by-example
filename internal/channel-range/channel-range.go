package channel_range

import (
	"fmt"
	"time"
)

func ChannelRange() {
	fmt.Println("Channel range package output:")
	ch := make(chan string)

	go func() {
		for item := range ch {
			fmt.Println(item)
			//if item == "world" {
			//	close(ch)
			//}
		}
	}()

	ch <- "Hello"
	ch <- "world"
	time.Sleep(time.Millisecond * 4)

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}

	fmt.Println("")

}
