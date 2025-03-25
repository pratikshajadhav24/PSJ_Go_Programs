//Find Factorial of a Number

package main

import "fmt"

func main() {
	num := 5
	factorial := 1
	for num > 0 {
		factorial *= num
		num--
	}
	fmt.Println("Factorial:", factorial)
}
