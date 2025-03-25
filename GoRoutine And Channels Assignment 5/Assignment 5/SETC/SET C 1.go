/* 1. Write a go program to implement the checkpoint synchronization problem which is a problem of
synchronizing multiple tasks. Consider a workshop where several workers assembling details of some
mechanism. When each of them completes his work, they put the details together. There is no store, so a
worker who finished its part first must wait for others before starting another one. Putting details
together is the checkpoint at which tasks synchronize themselves before going their paths apart.*/


package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(id int, checkpoint *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(time.Second * time.Duration(id))
	fmt.Printf("Worker %d reached checkpoint\n", id)

	checkpoint.Done()
	checkpoint.Wait() // Wait for all workers to reach checkpoint

	fmt.Printf("Worker %d continued after checkpoint\n", id)
}

func main() {
	totalWorkers := 3
	var checkpoint sync.WaitGroup

	checkpoint.Add(totalWorkers)
	for i := 1; i <= totalWorkers; i++ {
		wg.Add(1)
		go worker(i, &checkpoint)
	}

	wg.Wait()
	fmt.Println("All workers completed.")
}
