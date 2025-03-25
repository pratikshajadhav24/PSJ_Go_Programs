//  WAP in go language to demonstrate use of names returns variables.

package main
import "fmt"

func sumAndProduct(a, b int) (sum int, product int) {
    sum = a + b
    product = a * b
    return
}

func main() {
    num1, num2 := 5, 3
    sum, product := sumAndProduct(num1, num2)
    fmt.Println("Sum:", sum)
    fmt.Println("Product:", product)
}
