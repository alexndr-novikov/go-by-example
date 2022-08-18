package select_examples

import (
	"fmt"
	"time"
)

func Select() {
	fmt.Println("Select package output:")
	chan1 := make(chan string, 1)
	chan2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		chan1 <- "Message 1 received (2sec delay)"
	}()

	go func() {
		time.Sleep(time.Second)
		chan2 <- "Message 2 received (1 sec delay)"
	}()

	start := time.Now()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Println(msg1)
		case msg2 := <-chan2:
			fmt.Println(msg2)

		}
	}

	fmt.Println("Execution time:", time.Since(start))
	fmt.Println("")
}
