//  Creating a New Goroutine
package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from Goroutine!")
}

func main() {
	go sayHello() // Launch a new Goroutine
	time.Sleep(time.Second) // Give time for Goroutine to execute
	fmt.Println("Main function finished")
}
