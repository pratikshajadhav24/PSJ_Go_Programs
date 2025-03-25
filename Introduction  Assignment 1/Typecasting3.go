// Convert Float to Integer

package main

import "fmt"

func main() {
    var floatNum float64 = 45.67
    var intNum int = int(floatNum) // Float to Integer (truncation)

    fmt.Printf("Float: %.2f\n", floatNum)
    fmt.Println("Converted to Integer:", intNum)
}
