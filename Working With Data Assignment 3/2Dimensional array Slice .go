// Initializing a 2D Array of Slices

package main

import "fmt"

func main() {
    // Declare a 2D array of slices with 3 rows
    var arr [3][]int

    // Initialize each row with a different-sized slice
    arr[0] = []int{1, 2, 3}
    arr[1] = []int{4, 5}
    arr[2] = []int{6, 7, 8, 9}

    // Print the array of slices
    fmt.Println("2D Array of Slices:")
    for i, row := range arr {
        fmt.Printf("Row %d: %v\n", i, row)
    }
}
