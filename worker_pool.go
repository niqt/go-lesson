package main

import (
	"fmt"
	"sync"
)

func workerPool() {
	jobs := make(chan int, 10)
	var wg sync.WaitGroup

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for job := range jobs {
				fmt.Printf("Worker %d processing job %d\n", id, job)
			}
		}(w)
	}

	// Send 10 jobs
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	fmt.Println("All jobs done!")
}
