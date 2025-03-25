package main
import "fmt"

func multiply(a, b int) (product int) {
    product = a * b // 'product' is a named return variable
    return          // Implicitly returns 'product'
}

func main() {
    result := multiply(3, 4)
    fmt.Println("Product:", result)
}
