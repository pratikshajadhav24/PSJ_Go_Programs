// 2. WAP in GO program that executes 5 go routines simultaneously which generates numbers from 0 to 10,waiting between 0 and 250 ms after each go routine.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func generateNumbers(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i++ {
		fmt.Printf("Goroutine %d: %d\n", id, i)
		time.Sleep(time.Duration(rand.Intn(250)) * time.Millisecond)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go generateNumbers(i, &wg)
	}

	wg.Wait()
}
