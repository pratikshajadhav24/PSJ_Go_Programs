// Check Whether a Number is Even or Odd

package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	switch {
	case num%2 == 0:
		fmt.Println("Even number")
	default:
		fmt.Println("Odd number")
	}
}
'