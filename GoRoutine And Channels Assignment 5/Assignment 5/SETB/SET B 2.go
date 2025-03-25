// 2. WAP in Go main go routine to read and write Fibonacci series to the channel.

package main

import "fmt"

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go fibonacci(10, ch)

	fmt.Println("Fibonacci Series:")
	for num := range ch {
		fmt.Println(num)
	}
}
