// Creating and Initializing a Multidimensional Slice

package main

import "fmt"

func main() {
    // Create a 2D slice with 3 rows
    rows := 3
    cols := 4
    matrix := make([][]int, rows)

    // Initialize each row with a slice of length 'cols'
    for i := range matrix {
        matrix[i] = make([]int, cols)
    }

    // Assign values
    matrix[0][0] = 1
    matrix[1][1] = 2
    matrix[2][2] = 3

    // Print the 2D slice
    fmt.Println("2D Slice:")
    for _, row := range matrix {
        fmt.Println(row)
    }
}
