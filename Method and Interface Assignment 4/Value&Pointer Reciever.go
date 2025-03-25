// value Receiver and Pointer receiver Value

package main
import (
"fmt"
)
type Person struct {
Name string
Age int
}
func ValueReceiver(p Person) {
p.Name = "ABC"
fmt.Println("Inside ValueReceiver : ", p.Name)
}
func PointerReceiver(p *Person) {
p.Age = 24
fmt.Println("Inside PointerReceiver model: ", p.Age)
}
func main() {
	p := Person{"XYZ", 28}
p1:= &Person{"LMN", 68}
ValueReceiver(p)
fmt.Println("Inside Main after value receiver : ", p.Name)
PointerReceiver(p1)
fmt.Println("Inside Main after value receiver : ", p1.Age)
}

