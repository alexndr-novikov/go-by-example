package goroutines

import (
	"fmt"
	"time"
)

func counter(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(fmt.Sprintf("%s - %d", s, i))
	}
}

func RunSample() {
	fmt.Println("Goroutines package output:")
	counter("direct")
	go counter("goroutine")
	go counter("mixed")
	go func(s string) {
		fmt.Println(s)
	}("anonymous")
	time.Sleep(time.Second)
	fmt.Println("")
}
