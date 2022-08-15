package channels

import "fmt"

func Channels() {
	fmt.Println("Channels package output:")
	pipe := make(chan string)
	go func() { pipe <- "message1" }()

	msg := <-pipe
	fmt.Println(msg)
	fmt.Println("")
}
