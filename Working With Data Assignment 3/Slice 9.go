//  Slices Refer to an Underlying Array

package main

import "fmt"

func main() {
    // Create an array
    arr := [5]int{1, 2, 3, 4, 5}

    // Create a slice referring to part of the array
    slice := arr[1:4]  // This refers to elements 2, 3, 4

    // Display the slice and the array
    fmt.Println("Array:", arr)   // [1 2 3 4 5]
    fmt.Println("Slice:", slice) // [2 3 4]

    // Modify the slice (modifies the underlying array)
    slice[0] = 99

    // Display the modified array and slice
    fmt.Println("Modified Array:", arr)   // [1 99 3 4 5]
    fmt.Println("Modified Slice:", slice) // [99 3 4]
}
