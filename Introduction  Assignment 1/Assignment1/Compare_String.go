// WAP in go language to accept two strings and compare them.

package main

import "fmt"

func main() {
	var str1, str2 string
	fmt.Print("Enter first string: ")
	fmt.Scanln(&str1)
	fmt.Print("Enter second string: ")
	fmt.Scanln(&str2)

	if str1 == str2 {
		fmt.Println("The strings are equal.")
	} else {
		fmt.Println("The strings are not equal.")
	}
}
