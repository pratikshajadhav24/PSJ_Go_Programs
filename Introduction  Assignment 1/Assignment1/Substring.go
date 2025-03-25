//  WAP in go language to check whether first string is substring of another string or not.

package main

import "fmt"

func main() {
	var str1, str2 string
	fmt.Print("Enter first string: ")
	fmt.Scanln(&str1)
	fmt.Print("Enter second string: ")
	fmt.Scanln(&str2)

	if len(str1) <= len(str2) && (str2[:len(str1)] == str1 || str2[1:] == str1) {
		fmt.Println("The first string is a substring of the second string.")
	} else {
		fmt.Println("The first string is not a substring of the second string.")
	}
}
