//Using Multiple Goroutines with WaitGroup

package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(name, "started")
	time.Sleep(time.Second)
	fmt.Println(name, "completed")
}

func main() {
	var wg sync.WaitGroup

	tasks := []string{"Task 1", "Task 2", "Task 3", "Task 4"}

	for _, t := range tasks {
		wg.Add(1)
		go task(t, &wg)
	}

	wg.Wait()
	fmt.Println("All tasks are finished")
}
