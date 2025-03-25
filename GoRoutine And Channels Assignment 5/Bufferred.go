//  Buffered Channel

package main

import "fmt"

func main() {
	ch := make(chan string, 2) /

	ch <- "Buffered Hello 1" 
	ch <- "Buffered Hello 2"

	fmt.Println(<-ch) 
	fmt.Println(<-ch) 
}
