// Convert Integer to Float 

package main

import "fmt"

func main() {
    var intNum int = 10
    var floatNum float64 = float64(intNum) // Integer to Float
    var intNumAgain int = int(floatNum)   // Float to Integer

    fmt.Println("Integer:", intNum)
    fmt.Println("Converted to Float:", floatNum)
    fmt.Println("Converted back to Integer:", intNumAgain)
}
