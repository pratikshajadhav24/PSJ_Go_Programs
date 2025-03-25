//Call by Value with String

package main
import "fmt"

func modifyString(s string) {
    s = "Modified"
    fmt.Println("Inside function, s:", s)
}

func main() {
    str := "Original"
    fmt.Println("Before function call, str:", str)
    modifyString(str)
    fmt.Println("After function call, str:", str)
}
