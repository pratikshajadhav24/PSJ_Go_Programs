// Convert String to Float

package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := "78.95"
    num, err := strconv.ParseFloat(str, 64) // String to Float
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("String:", str)
        fmt.Println("Converted to Float:", num)
    }
}
