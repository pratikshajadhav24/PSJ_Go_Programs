// WAP in go language to create a file and write hello world in it and close the file by using defer statement.

package main
import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("hello.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }

    // Ensure the file gets closed at the end of the function
    defer file.Close()

    _, err = file.WriteString("Hello, World!")
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }

    fmt.Println("File created and written to successfully.")
}
