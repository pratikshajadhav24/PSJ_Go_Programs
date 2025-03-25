
package main

import "fmt"

func main() {
    // Create a slice with an initial length and capacity
    nums := []int{10, 20, 30, 40, 50}

    // Length of the slice
    fmt.Println("Length of slice:", len(nums)) // Output: 5

    // Capacity of the slice
    fmt.Println("Capacity of slice:", cap(nums)) // Output: 5
}
