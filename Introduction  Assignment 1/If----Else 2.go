// Number is Even or Odd

package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Println("The number is even.")
	} else {
		fmt.Println("The number is odd.")
	}
}
