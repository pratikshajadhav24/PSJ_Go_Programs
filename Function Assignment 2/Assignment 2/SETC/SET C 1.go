// WAP in go language to illustrate the concept of call by value

.package main
import "fmt"

func modifyValue(x int) {
    x = 100
    fmt.Println("Inside function, x =", x)
}

func main() {
    num := 50
    fmt.Println("Before function call, num =", num)
    modifyValue(num)
    fmt.Println("After function call, num =", num)
}
