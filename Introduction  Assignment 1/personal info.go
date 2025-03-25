package main

import "fmt"

func main() {
    name := "John Doe"
    age := 25
    city := "New York"

    fmt.Println("Personal Information:")
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("City:", city)
    fmt.Printf("Name: %s, Age: %d, City: %s\n", name, age, city)
}
