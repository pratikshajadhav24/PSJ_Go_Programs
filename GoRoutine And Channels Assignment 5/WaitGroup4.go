// Using WaitGroup in a Function Call

package main

import (
	"fmt"
	"sync"
	"time"
)

func processTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Processing task %d\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Task %d completed\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go processTask(i, &wg)
	}

	wg.Wait()
	fmt.Println("All tasks are done")
}
