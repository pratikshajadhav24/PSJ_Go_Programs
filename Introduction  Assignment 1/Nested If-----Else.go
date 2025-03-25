// Check if a Number is Positive, Negative, or Zero

package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	if num > 0 {
		fmt.Println("The number is positive.")
	} else if num < 0 {
		fmt.Println("The number is negative.")
	} else {
		fmt.Println("The number is zero.")
	}
}
