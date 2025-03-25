//Direct Factorial (Using Recursion)


package main
import "fmt"

func factorialDirect(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorialDirect(n-1)
}

func main() {
    fmt.Println("Factorial of 5:", factorialDirect(5))
}
