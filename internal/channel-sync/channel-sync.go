package channel_sync

import (
	"fmt"
	"time"
)

func worker(c chan bool) {
	fmt.Println("Worker started")
	time.Sleep(time.Second)
	fmt.Println("Done")
	c <- true
}

func ChannelSync() {
	fmt.Println("Channel sync package output:")
	channel := make(chan bool, 1)
	go worker(channel)
	<-channel
	go worker(channel)
	fmt.Println("Second worker skipped of half-skipped")
}
