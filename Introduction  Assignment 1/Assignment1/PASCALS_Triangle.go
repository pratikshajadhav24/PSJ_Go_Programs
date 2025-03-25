// WAP in go language to print PASCALS triangle.

package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number of rows: ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("1 ")
		}
		fmt.Println()
	}
}

