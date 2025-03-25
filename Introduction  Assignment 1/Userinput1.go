// Userinput1.go

package main

import "fmt"

func main() {
    var name string
    fmt.Print("Enter your name: ")
    fmt.Scanln(&name) // Taking user input
    fmt.Println("Hello,", name)
}
