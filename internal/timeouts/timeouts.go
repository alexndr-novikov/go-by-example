package timeouts

import (
	"fmt"
	"time"
)

func RunSample() {
	fmt.Println("Timeouts package output:")
	chan1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		chan1 <- "Message retrieved in 2 seconds"
	}()
	select {
	case msg := <-chan1:
		fmt.Println(msg)
	case <-time.After(time.Second):
		fmt.Println("Timeout 1 second")
	}

	chan2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		chan2 <- "Message retrieved in 2 seconds"
	}()
	select {
	case msg := <-chan2:
		fmt.Println(msg)
	case <-time.After(time.Second * 3):
		fmt.Println("Timeout 3 seconds")
	}

	fmt.Println("")
}
