// 1. WAP in Go language to add two integers and write code for unit test to test this code.

package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Sum:", Add(2, 3))
}
