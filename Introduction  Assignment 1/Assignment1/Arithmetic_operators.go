//  WAP in go language to accept user choice and print answer of using arithmetical operators.

package main

import "fmt"

func main() {
	var num1, num2, choice int
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	fmt.Println("Choose operation: 1.Addition 2.Subtraction 3.Multiplication 4.Division")
	fmt.Print("Enter choice: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Println("Result: ", num1+num2)
	case 2:
		fmt.Println("Result: ", num1-num2)
	case 3:
		fmt.Println("Result: ", num1*num2)
	case 4:
		fmt.Println("Result: ", num1/num2)
	default:
		fmt.Println("Invalid choice.")
	}
}
