package main

import "fmt"

func main() {
    var matrix [3][3]int // 3x3 integer matrix

    // Assigning values
    matrix[0][0] = 1
    matrix[0][1] = 2
    matrix[0][2] = 3
    matrix[1][0] = 4
    matrix[1][1] = 5
    matrix[1][2] = 6
    matrix[2][0] = 7
    matrix[2][1] = 8
    matrix[2][2] = 9

    fmt.Println("2D Array:", matrix)
}
