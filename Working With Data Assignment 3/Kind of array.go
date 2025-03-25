package main

import (
    "fmt"
    "reflect"
)

func main() {
    arr := [3]float64{1.1, 2.2, 3.3}

    fmt.Println("Type:", reflect.TypeOf(arr))  // Full type with size
    fmt.Println("Kind:", reflect.TypeOf(arr).Kind()) // Kind of the type
}
