// Using cap() with len()

package main

import "fmt"

func main() {
	ch := make(chan int, 4) 

	ch <- 10
	ch <- 20

	fmt.Println("Capacity:", cap(ch))       // 4
	fmt.Println("Current length:", len(ch)) // 2
}
