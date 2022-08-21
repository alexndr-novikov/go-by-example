package closing_channels

import (
	"fmt"
)

func ClosingChannels() {
	fmt.Println("Closing channels package output:")
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Received job", j)
			} else {
				fmt.Println("Received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 0; j < 3; j++ {
		jobs <- j
		fmt.Println("Sent job", j)
	}
	close(jobs)
	fmt.Println("Sent all jobs")
	<-done
	fmt.Println("")
}
