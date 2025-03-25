// Add Two Numbers

package main

import "fmt"

func main() {
    var num1, num2 int
    fmt.Print("Enter the first number: ")
    fmt.Scanln(&num1) // Input for num1
    fmt.Print("Enter the second number: ")
    fmt.Scanln(&num2) // Input for num2

    sum := num1 + num2
    fmt.Println("The sum of", num1, "and", num2, "is:", sum)
}
