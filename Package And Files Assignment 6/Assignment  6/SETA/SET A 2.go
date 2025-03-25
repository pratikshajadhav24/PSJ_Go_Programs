// 2. WAP in Go language using user defined package calculator that performs one calculator operation as per the0user's choice.



// Create the package: Create a folder named calculator and inside it create a file calculator.go with the following code:
// calculator/calculator.go
package calculator

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) int {
	if b != 0 {
		return a / b
	}
	return 0 // Return 0 for division by zero
}



package main

import (
	"calculator"
	"fmt"
)

func main() {
	var choice, a, b int
	fmt.Println("Calculator Menu:")
	fmt.Println("1. Add\n2. Subtract\n3. Multiply\n4. Divide")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)

	switch choice {
	case 1:
		fmt.Println("Result:", calculator.Add(a, b))
	case 2:
		fmt.Println("Result:", calculator.Subtract(a, b))
	case 3:
		fmt.Println("Result:", calculator.Multiply(a, b))
	case 4:
		fmt.Println("Result:", calculator.Divide(a, b))
	default:
		fmt.Println("Invalid choice")
	}
}

