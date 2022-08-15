package channel_buffering

import "fmt"

func ChannelBuffering() {
	fmt.Println("Channel buffering package output:")
	messages := make(chan string, 2)
	messages <- "msg1"
	messages <- "msg2"
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println("")
}
