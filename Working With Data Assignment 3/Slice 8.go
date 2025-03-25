// Creating a Slice with Low and High Value Using Colon Syntax

package main

import "fmt"

func main() {
    // Initialize a slice
    nums := []int{10, 20, 30, 40, 50, 60, 70}

    // Create a sub-slice using colon syntax (from index 2 to 5)
    subSlice := nums[2:6] // Includes elements at indices 2, 3, 4, and 5

    fmt.Println("Original Slice:", nums)       // [10 20 30 40 50 60 70]
    fmt.Println("Sub-slice from index 2 to 5:", subSlice) // [30 40 50 60]
}
