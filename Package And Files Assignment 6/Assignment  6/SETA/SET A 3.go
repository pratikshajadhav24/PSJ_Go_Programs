// 3. WAP in Go language to create an user defined package to find out the area of a rectangle.

Create a new folder for your project (e.g., rectangle_area).

Inside this folder, create another folder named geometry to store the package.

Inside the geometry folder, create a file named rectangle.go.

Code for rectangle.go (inside geometry folder)

package geometry

// Area calculates the area of a rectangle.
func Area(length, width float64) float64 {
    return length * width
}
Step 2: Create the Main Program
Go back to the rectangle_area folder.

Create a file named main.go.

Code for main.go

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
Step 3: Initialize the Go Module
Open a terminal in the rectangle_area folder.

Run the following command to initialize the Go module:

go mod init rectangle_area
Step 4: Run the Program
Execute the program using:

go run main.go
