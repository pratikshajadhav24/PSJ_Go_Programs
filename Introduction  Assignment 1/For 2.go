
//Calculate the Sum of the First 10 Natural Numbers

package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum of first 10 natural numbers:", sum)
}
