// Function with Return Type

package main

import "fmt"

// Function with a return type
func multiply(a int, b int) int {
    return a * b
}

func main() {
    result := multiply(4, 5)
    fmt.Println("Product:", result)
}
