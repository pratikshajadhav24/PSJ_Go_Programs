// 3. Write a go program that creates a slice of integers, checks numbers from slice are even or odd and further sent to respective go routines through channel and display values received by go routines.

package main

import "fmt"

func checkEven(nums []int, ch chan int) {
	for _, num := range nums {
		if num%2 == 0 {
			ch <- num
		}
	}
	close(ch)
}

func checkOdd(nums []int, ch chan int) {
	for _, num := range nums {
		if num%2 != 0 {
			ch <- num
		}
	}
	close(ch)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenCh := make(chan int)
	oddCh := make(chan int)

	go checkEven(numbers, evenCh)
	go checkOdd(numbers, oddCh)

	fmt.Println("Even Numbers:")
	for num := range evenCh {
		fmt.Println(num)
	}

	fmt.Println("Odd Numbers:")
	for num := range oddCh {
		fmt.Println(num)
	}
}
