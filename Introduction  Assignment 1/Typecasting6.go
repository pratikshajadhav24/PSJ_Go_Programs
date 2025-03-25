// Convert Float to String

package main

import (
    "fmt"
    "strconv"
)

func main() {
    floatNum := 123.456
    str := strconv.FormatFloat(floatNum, 'f', 2, 64) // Float to String

    fmt.Printf("Float: %.2f\n", floatNum)
    fmt.Println("Converted to String:", str)
}
