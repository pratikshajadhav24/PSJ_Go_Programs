// . Write a program in go language to print multiplication of two numbers using method.

package main

import "fmt"

func multiply(a, b int) int {
    return a * b
}

func main() {
    var num1, num2 int

    fmt.Print("Enter first number: ")
    fmt.Scan(&num1)

    fmt.Print("Enter second number: ")
    fmt.Scan(&num2)

    result := multiply(num1, num2)
    fmt.Printf("Multiplication of %d and %d is: %d\n", num1, num2, result)
}
