package signals

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func RunSample() {
	fmt.Println("Signals package output:")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)
	go func() {
		sig := <-sigs
		fmt.Println("\r\nSignal: ", sig)
		done <- true
	}()

	fmt.Println("Waiting signal")
	<-done
	fmt.Println("Exiting custom")
	fmt.Println("")
}
