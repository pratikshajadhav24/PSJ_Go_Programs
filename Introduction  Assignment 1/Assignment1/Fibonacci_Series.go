//  WAP in go language to print Fibonacci series of n terms.

package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number of terms: ")
	fmt.Scanln(&n)

	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
}
