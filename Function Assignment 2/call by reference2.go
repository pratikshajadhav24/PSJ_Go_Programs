// Call by Reference with String

package main
import "fmt"

func modifyString(s *string) {
    *s = "Modified" // Dereferencing the pointer to modify the original string
    fmt.Println("Inside function, s:", *s)
}

func main() {
    str := "Original"
    fmt.Println("Before function call, str:", str)
    modifyString(&str) // Passing the address of the variable
    fmt.Println("After function call, str:", str)
}
