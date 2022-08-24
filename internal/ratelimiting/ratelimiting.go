package ratelimiting

import (
	"fmt"
	"time"
)

func RunSample() {
	fmt.Println("Rate limiting package output:")
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for request := range requests {
		<-limiter
		fmt.Println("Handle request", request, "time", time.Now())
	}

	burstyLimiter := make(chan time.Time, 5)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("Request ", req, " handled in ", time.Now())
	}
	fmt.Println("")
}
