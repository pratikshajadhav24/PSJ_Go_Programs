// Check if a Number is a Square Number

package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	if num < 0 {
		fmt.Println("The number is not a square number.")
	} else {
		sqrt := int(float64(num) * 0.5)
		if sqrt*sqrt == num {
			fmt.Println("The number is a square number.")
		} else {
			fmt.Println("The number is not a square number.")
		}
	}
}
