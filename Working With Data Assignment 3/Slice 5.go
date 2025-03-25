//  Accessing a Particular Element in a Slice
package main

import "fmt"

func main() {
    // Initialize a slice
    nums := []int{10, 20, 30, 40, 50}

    // Access and print specific elements
    fmt.Println("First element:", nums[0])  // 10
    fmt.Println("Third element:", nums[2])  // 30
    fmt.Println("Last element:", nums[4])   // 50
}
