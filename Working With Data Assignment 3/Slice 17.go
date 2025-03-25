// Looping Through a Two-Dimensional Slice

package main

import "fmt"

func main() {
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }

    fmt.Println("Looping through a 2D slice:")
    for rowIndex, row := range matrix {
        for colIndex, value := range row {
            fmt.Printf("Row: %d, Col: %d, Value: %d\n", rowIndex, colIndex, value)
        }
    }
}

