package tickers

import (
	"fmt"
	"time"
)

func RunSample() {
	fmt.Println("tickers package output:")
	ticker := time.NewTicker(time.Millisecond * 500)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick", t)
			}
		}
	}()
	time.Sleep(time.Millisecond * 2100)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
	fmt.Println("")
}
