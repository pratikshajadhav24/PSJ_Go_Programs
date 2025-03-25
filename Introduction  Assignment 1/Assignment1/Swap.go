//   WAP in go language to swap the number without temporary variable.

package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Enter two numbers: ")
	fmt.Scanln(&a, &b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Println("After swapping: a =", a, "b =", b)
}
