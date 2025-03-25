//  WAP in go language to check whether accepted number is single digit or not.

package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	if num >= 0 && num <= 9 {
		fmt.Println("The number is a single digit.")
	} else {
		fmt.Println("The number is not a single digit.")
	}
}
