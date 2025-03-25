//Basic Call by Value
go
Copy
Edit


package main
import "fmt"

func modifyValue(x int) {
    x = 20
    fmt.Println("Inside function, x:", x)
}

func main() {
    a := 10
    fmt.Println("Before function call, a:", a)
    modifyValue(a)
    fmt.Println("After function call, a:", a)
}
