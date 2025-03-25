// Use of sleep ()

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")
	time.Sleep(2 * time.Second) // Pauses execution for 2 seconds
	fmt.Println("End after 2 seconds")
}
