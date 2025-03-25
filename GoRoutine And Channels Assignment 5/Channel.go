// Creating a New Channel

package main

import "fmt"

func main() {
	ch := make(chan string) 

	go func() {
		ch <- "Hello, Channel!" 
	}()

	msg := <-ch 
	fmt.Println(msg)
}
