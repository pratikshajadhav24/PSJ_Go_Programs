// Goroutine

package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(500 * time.Millisecond) 
	}
}

func main() {
	go printNumbers()

	fmt.Println("Main function execution")

	time.Sleep(3 * time.Second)
	fmt.Println("Main function finished")
}
