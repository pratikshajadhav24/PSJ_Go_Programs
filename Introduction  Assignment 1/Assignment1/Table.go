//  WAP in go to print table of given number.

package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	for i := 1; i <= 10; i++ {
		fmt.Println(num, "x", i, "=", num*i)
	}
}
