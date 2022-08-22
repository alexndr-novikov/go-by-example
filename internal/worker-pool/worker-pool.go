package worker_pool

import (
	"fmt"
	"time"
)

func worker(id int, jobs chan int, results chan int) {
	for j := range jobs {
		fmt.Println("Worker ", id, " started job", j)
		time.Sleep(time.Second * 2)
		results <- j * j
		fmt.Println("Worker ", id, " finished job", j)
	}
}

func WorkerPool() {
	fmt.Println("Worker pool package output:")
	startTime := time.Now().Second()
	numJobs := 10
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	for i := 0; i < 5; i++ {
		go worker(i, jobs, results)
	}
	for i := 0; i < numJobs; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 0; i < numJobs; i++ {
		<-results
	}
	fmt.Println("Seconds taken in parallel: ", time.Now().Second()-startTime, "secs")
	fmt.Println("If not in parallel it would be: ", numJobs*2, "secs")
	fmt.Println("")
}
