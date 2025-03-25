// Convert Integer to String


package main

import (
    "fmt"
    "strconv"
)

func main() {
    num := 456
    str := strconv.Itoa(num) // Integer to String

    fmt.Println("Integer:", num)
    fmt.Println("Converted to String:", str)
}
