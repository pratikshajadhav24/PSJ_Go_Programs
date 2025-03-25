// 1. Write a go program using go routine and channel that will print the sum of the squares and cubes of the individual digits of a number.

package main

import (
	"fmt"
)

func sumOfSquares(n int, ch chan int) {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	ch <- sum
}

func sumOfCubes(n int, ch chan int) {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit * digit
		n /= 10
	}
	ch <- sum
}

func main() {
	number := 123
	squaresCh := make(chan int)
	cubesCh := make(chan int)

	go sumOfSquares(number, squaresCh)
	go sumOfCubes(number, cubesCh)

	squares := <-squaresCh
	cubes := <-cubesCh
	finalSum := squares + cubes

	fmt.Println("Sum of squares =", squares)
	fmt.Println("Sum of cubes =", cubes)
	fmt.Println("Final sum of squares and cubes =", finalSum)
}
