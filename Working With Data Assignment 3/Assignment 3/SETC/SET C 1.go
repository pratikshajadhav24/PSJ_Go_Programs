// WAP in go language to accept two matrices and display it's multiplication.

package main

import "fmt"

func main() {
    var r1, c1, r2, c2 int
    fmt.Print("Enter rows and columns for first matrix: ")
    fmt.Scan(&r1, &c1)
    fmt.Print("Enter rows and columns for second matrix: ")
    fmt.Scan(&r2, &c2)

    if c1 != r2 {
        fmt.Println("Matrix multiplication not possible!")
        return
    }

    A := make([][]int, r1)
    B := make([][]int, r2)
    C := make([][]int, r1)

    fmt.Println("Enter first matrix elements:")
    for i := 0; i < r1; i++ {
        A[i] = make([]int, c1)
        for j := 0; j < c1; j++ {
            fmt.Scan(&A[i][j])
        }
    }

    fmt.Println("Enter second matrix elements:")
    for i := 0; i < r2; i++ {
        B[i] = make([]int, c2)
        for j := 0; j < c2; j++ {
            fmt.Scan(&B[i][j])
        }
    }

    for i := 0; i < r1; i++ {
        C[i] = make([]int, c2)
        for j := 0; j < c2; j++ {
            for k := 0; k < c1; k++ {
                C[i][j] += A[i][k] * B[k][j]
            }
        }
    }

    fmt.Println("Resultant Matrix:")
    for _, row := range C {
        fmt.Println(row)
    }
}
