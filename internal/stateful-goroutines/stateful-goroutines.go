package stateful_goroutines

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type ReadOperation struct {
	key  int
	resp chan int
}

type WriteOperaion struct {
	key   int
	value int
	resp  chan bool
}

func RunSample() {
	fmt.Println("Stateful goroutines package output:")

	var readOps uint64
	var writeOps uint64

	reads := make(chan ReadOperation)
	writes := make(chan WriteOperaion)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.value
				write.resp <- true
			}
		}
	}()

	for i := 0; i < 100; i++ {
		go func() {
			for {
				readOp := ReadOperation{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- readOp
				<-readOp.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for j := 0; j < 10; j++ {
		go func() {
			for {
				write := WriteOperaion{
					key:   rand.Intn(5),
					value: rand.Intn(100),
					resp:  make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	fmt.Println("Read ops: ", atomic.LoadUint64(&readOps))
	fmt.Println("Write ops: ", atomic.LoadUint64(&writeOps))
	fmt.Println("")
}
