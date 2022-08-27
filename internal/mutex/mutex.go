package mutex

import (
	"fmt"
	"sync"
)

type Container struct {
	mutex    sync.Mutex
	counters map[string]int
}

func (container *Container) increment(key string) {
	container.mutex.Lock()
	defer container.mutex.Unlock()
	container.counters[key]++
}

func RunSample() {
	var wg sync.WaitGroup
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}
	fmt.Println("Mutex package output:")
	doIncrement := func(key string, times int) {
		for i := 0; i < times; i++ {
			c.increment(key)
		}
		wg.Done()
	}
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)
	go doIncrement("a", 10000)
	wg.Wait()
	fmt.Println(c.counters)
	fmt.Println("")
}
