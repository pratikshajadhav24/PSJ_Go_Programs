//Declare and Print Variables of All Basic Data Types

package main

import "fmt"

func main() {
    var intVar int = 10
    var floatVar float64 = 20.5
    var stringVar string = "Hello, Go!"
    var boolVar bool = true
    var charVar rune = 'G'

    fmt.Println("Integer Value:", intVar)
    fmt.Println("Float Value:", floatVar)
    fmt.Println("String Value:", stringVar)
    fmt.Println("Boolean Value:", boolVar)
    fmt.Printf("Character Value: %c\n", charVar)
}
