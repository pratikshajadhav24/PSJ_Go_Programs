// Sum of Numbers Until It Exceeds 50

package main

import "fmt"

func main() {
	sum := 0
	i := 1
	for {
		sum += i
		if sum > 50 {
			break
		}
		i++
	}
	fmt.Println("Sum exceeds 50:", sum)
}
