// 3. WAP in Go how to create channel and illustrate how to close a channel using for range loop and close function.

package main

import "fmt"

func sendValues(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch) // Closing channel
}

func main() {
	ch := make(chan int)

	go sendValues(ch)

	fmt.Println("Values received from channel:")
	for val := range ch {
		fmt.Println(val)
	}
}
