package main
import "fmt"

func factorialIndirect(n int) int {
    return calculateFactorial(n, 1)
}

func calculateFactorial(n, result int) int {
    if n == 0 {
        return result
    }
    return calculateFactorial(n-1, result*n)
}

func main() {
    fmt.Println("Factorial of 5:", factorialIndirect(5))
}
