//Working with Complex Numbers

package main

import "fmt"

func main() {
    var c1 complex64 = 3 + 4i
    var c2 complex64 = 1 + 2i
    var sum = c1 + c2

    fmt.Println("Complex Number 1:", c1)
    fmt.Println("Complex Number 2:", c2)
    fmt.Println("Sum of Complex Numbers:", sum)
    fmt.Printf("Real Part: %.2f, Imaginary Part: %.2f\n", real(sum), imag(sum))
}
