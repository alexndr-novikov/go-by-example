package waitgroups

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Println("Worker ", id, " launched")
	time.Sleep(time.Second)
	fmt.Println("Worker ", id, " finished")
}

func RunSample() {
	fmt.Println("Wait groups package output:")
	var group sync.WaitGroup
	for i := 1; i <= 10; i++ {
		group.Add(1)

		i := i

		go func() {
			defer group.Done()
			worker(i)
		}()
	}
	group.Wait()
}
