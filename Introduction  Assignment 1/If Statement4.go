// Check if a Number is Greater Than 100

package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	if num > 100 {
		fmt.Println("The number is greater than 100.")
	}
}
