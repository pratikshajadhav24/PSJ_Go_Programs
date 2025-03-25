// Assigning a Value to a Specific Element in a Slice

package main

import "fmt"

func main() {
    // Initialize a slice
    nums := []int{1, 2, 3, 4, 5}

    fmt.Println("Original slice:", nums) // [1 2 3 4 5]

    // Assign a new value to the second element (index 1)
    nums[1] = 20

    // Assign a new value to the fourth element (index 3)
    nums[3] = 40

    fmt.Println("Modified slice:", nums) // [1 20 3 40 5]
}
