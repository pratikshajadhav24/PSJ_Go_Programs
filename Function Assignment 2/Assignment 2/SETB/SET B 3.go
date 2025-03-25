// WAP in go language to show the compiler throws an error if a variable is declared but not used.

package main
import "fmt"

func main() {
    var unusedVar int
    fmt.Println("This will throw a compile error because unusedVar is not used.")
}
