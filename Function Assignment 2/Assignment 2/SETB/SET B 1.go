//  WAP in go language to swap two numbers using call by reference concept.

package main
import "fmt"

func swap(a, b *int) {
    *a, *b = *b, *a
}

func main() {
    var num1, num2 int
    fmt.Print("Enter first number: ")
    fmt.Scan(&num1)
    fmt.Print("Enter second number: ")
    fmt.Scan(&num2)
    
    fmt.Println("Before swapping: num1 =", num1, "num2 =", num2)
    
    swap(&num1, &num2)
    
    fmt.Println("After swapping: num1 =", num1, "num2 =", num2)
}
