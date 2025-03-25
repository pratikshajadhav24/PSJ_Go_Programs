package main
import "fmt"

func reverseAndLength(s string) (string, int) {
    r := ""
    for i := len(s) - 1; i >= 0; i-- {
        r += string(s[i])
    }
    return r, len(s)
}

func main() {
    reversed, length := reverseAndLength("hello")
    fmt.Println("Reversed:", reversed)
    fmt.Println("Length:", length)
}
