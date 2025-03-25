// Check if a Number is Positive

package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	if num > 0 {
		fmt.Println("The number is positive.")
	}
}
