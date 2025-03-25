//  Function as a Parameter

package main

import "fmt"

// Function that accepts another function as a parameter
func operate(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}

func main() {
    sum := func(a, b int) int { return a + b }
    fmt.Println("Result:", operate(3, 4, sum))
}
