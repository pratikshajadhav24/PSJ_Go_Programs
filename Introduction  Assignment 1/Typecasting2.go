// Convert String to Integer

package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := "123"
    num, err := strconv.Atoi(str) // String to Integer
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("String:", str)
        fmt.Println("Converted to Integer:", num)
    }
}
