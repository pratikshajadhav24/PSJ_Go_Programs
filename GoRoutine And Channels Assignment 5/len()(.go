// Basic Example of len()

package main

import "fmt"

func main() {
	ch := make(chan int, 5) 
	ch <- 10
	ch <- 20

	fmt.Println("Current length of channel:", len(ch)) 
}
