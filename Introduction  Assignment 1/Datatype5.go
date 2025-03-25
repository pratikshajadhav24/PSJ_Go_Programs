Print Size of Data Types

package main

import (
    "fmt"
    "unsafe"
)

func main() {
    var intVar int
    var floatVar float64
    var boolVar bool
    var stringVar string

    fmt.Println("Size of int:", unsafe.Sizeof(intVar), "bytes")
    fmt.Println("Size of float64:", unsafe.Sizeof(floatVar), "bytes")
    fmt.Println("Size of bool:", unsafe.Sizeof(boolVar), "bytes")
    fmt.Println("Size of string:", unsafe.Sizeof(stringVar), "bytes")
}
