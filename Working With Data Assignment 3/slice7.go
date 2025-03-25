//  Display a Specific Element from the Slice

package main

import "fmt"

func main() {
    // Initialize a slice
    nums := []int{10, 20, 30, 40, 50}

    // Specify the index
    index := 2

    if index >= 0 && index < len(nums) {
        fmt.Printf("Element at index %d: %d\n", index, nums[index])
    } else {
        fmt.Println("Index out of range")
    }
}
