// Unbuffered Channel 

package main

import "fmt"

func main() {
	ch := make(chan string) 

	go func() {
		ch <- "Unbuffered Hello"
		fmt.Println("Sent to Unbuffered Channel")
	}()

	msg := <-ch 
	fmt.Println("Received:", msg)
}
