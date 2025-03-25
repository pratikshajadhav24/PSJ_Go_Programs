package main

import "fmt"

func main() {
    a := 10
    b := 20
    sum := a + b

    fmt.Println("The sum of", a, "and", b, "is:", sum)
    fmt.Printf("Using printf: %d + %d = %d\n", a, b, sum)
}
