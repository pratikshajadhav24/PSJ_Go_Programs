// Print Fibonacci Sequence Up to 100

package main

import "fmt"

func main() {
	a, b := 0, 1
	for a <= 100 {
		fmt.Println(a)
		a, b = b, a+b
	}
}
