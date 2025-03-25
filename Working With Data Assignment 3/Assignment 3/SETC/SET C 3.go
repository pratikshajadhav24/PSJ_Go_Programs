// WAP in go language to demonstrate working of slices (like append, remove, copy etc.)

package main

import "fmt"

func main() {
    // Initial Slice
    nums := []int{10, 20, 30, 40}
    fmt.Println("Initial Slice:", nums)

    // Append elements
    nums = append(nums, 50, 60)
    fmt.Println("After Append:", nums)

    // Remove an element (index 1)
    nums = append(nums[:1], nums[2:]...)
    fmt.Println("After Removing index 1:", nums)

    // Copy elements to a new slice
    newSlice := make([]int, len(nums))
    copy(newSlice, nums)
    fmt.Println("Copied Slice:", newSlice)
}
