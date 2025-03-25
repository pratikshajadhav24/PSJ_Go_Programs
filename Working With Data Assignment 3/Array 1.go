package main

import "fmt"

func main() {
    numbers := [...]int{1, 2, 3, 4, 5} // Compiler infers length (5)
    fmt.Println(numbers)
}
