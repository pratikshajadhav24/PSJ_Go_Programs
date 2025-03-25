// 1. .WAP in Go to create buffered channel, store few values in it and find channel capacity and length. Read values from channel and find modified length of a channel.

package main

import "fmt"

func main() {
	ch := make(chan int, 5) 
	ch <- 10
	ch <- 20
	ch <- 30

	fmt.Println("Channel Capacity:", cap(ch))
	fmt.Println("Channel Length:", len(ch))

	fmt.Println("Reading values:")
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("Modified Length of Channel:", len(ch))
}
