// two dimensional slice 

package main

import "fmt"

func main() {
    // Create a 2D slice with 3 rows and 4 columns (3x4 matrix)
    matrix := make([][]int, 3) // 3 rows

    // Create 4 columns for each row
    for i := range matrix {
        matrix[i] = make([]int, 4) // 4 columns per row
    }

    // Initialize values in the matrix
    matrix[0][0] = 1
    matrix[0][1] = 2
    matrix[0][2] = 3
    matrix[0][3] = 4

    matrix[1][0] = 5
    matrix[1][1] = 6
    matrix[1][2] = 7
    matrix[1][3] = 8

    matrix[2][0] = 9
    matrix[2][1] = 10
    matrix[2][2] = 11
    matrix[2][3] = 12

    // Print the 2D slice (matrix)
    fmt.Println("2D Slice:")
    for _, row := range matrix {
        fmt.Println(row)
    }
}
