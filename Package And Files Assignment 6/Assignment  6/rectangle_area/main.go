package main

import (
    "fmt"
    "rectangle_area/geometry" // Import the user-defined package
)

func main() {
    var length, width float64

    // Take input from the user
    fmt.Print("Enter the length of the rectangle: ")
    fmt.Scanln(&length)

    fmt.Print("Enter the width of the rectangle: ")
    fmt.Scanln(&width)

    // Calculate and display the area
    area := geometry.Area(length, width)
    fmt.Printf("The area of the rectangle is: %.2f\n", area)
}
