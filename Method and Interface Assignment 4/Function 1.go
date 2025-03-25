// Go program to illustrate the use of function

package main
import "fmt"
func area(length, width int)int{
Ar := length* width
return Ar
}
func main() {
fmt.Printf("Area of rectangle is : %d", area(15, 10))
}
