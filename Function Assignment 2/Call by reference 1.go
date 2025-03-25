package main
import "fmt"

// Function to modify an integer using a pointer
func modifyValue(num *int) {
    *num = 42 // Dereference the pointer to modify the original value
    fmt.Println("Inside function, num:", *num)
}

func main() {
    number := 10
    fmt.Println("Before function call, number:", number)
    modifyValue(&number) // Pass the address of the variable
    fmt.Println("After function call, number:", number)
}
