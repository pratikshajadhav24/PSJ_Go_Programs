

package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch) 

	for val := range ch {
		fmt.Println(val)
	}
}
