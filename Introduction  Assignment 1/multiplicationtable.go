package main

import "fmt"

func main() {
    number := 5
    fmt.Printf("Multiplication Table for %d:\n", number)

    for i := 1; i <= 10; i++ {
        fmt.Printf("%d x %d = %d\n", number, i, number*i)
    }
}
