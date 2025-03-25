SET A Q2 
// 2. WAP in Go language using user defined package calculator that performs one calculator operation as per the0user's choice.
package main

import (
    "calculator_project/calculator"
    "fmt"
)

func main() {
    var a, b float64
    var choice int

    fmt.Println("Enter two numbers:")
    fmt.Scan(&a, &b)

    fmt.Println("Choose an operation:")
    fmt.Println("1. Addition")
    fmt.Println("2. Subtraction")
    fmt.Println("3. Multiplication")
    fmt.Println("4. Division")
    fmt.Scan(&choice)

    switch choice {
    case 1:
        fmt.Println("Result:", calculator.Add(a, b))
    case 2:
        fmt.Println("Result:", calculator.Subtract(a, b))
    case 3:
        fmt.Println("Result:", calculator.Multiply(a, b))
    case 4:
        result, err := calculator.Divide(a, b)
        if err != nil {
            fmt.Println("Error:", err)
        } else {
            fmt.Println("Result:", result)
        }
    default:
        fmt.Println("Invalid choice")
    }
}
