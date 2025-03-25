//  Slice of Structs

package main

import "fmt"

// Define a struct
type Person struct {
    Name string
    Age  int
}

func main() {
    // Create a slice of structs
    people := []Person{
        {Name: "John", Age: 30},
        {Name: "Alice", Age: 25},
        {Name: "Bob", Age: 40},
    }

    // Display the slice of structs
    fmt.Println("People:", people)

    // Accessing individual struct elements
    fmt.Println("First person:", people[0])
    fmt.Println("Second person:", people[1])

    // Modifying a struct element in the slice
    people[2].Age = 42 // Modify Bob's age

    // Display after modification
    fmt.Println("Modified People:", people)
}
