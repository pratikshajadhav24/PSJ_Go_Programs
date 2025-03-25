
// Appending a Single Element
package main

import "fmt"

func main() {
    // Initial slice
    nums := []int{10, 20, 30}

    // Append a single element to the slice
    nums = append(nums, 40)

    // Display the updated slice
    fmt.Println("Updated Slice:", nums) // Output: [10 20 30 40]
}
