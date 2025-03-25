// WAP in go language to concatenate two strings using pointers.

package main

import "fmt"

func concatenate(str1, str2 *string) string {
	return *str1 + *str2
}

func main() {
	str1 := "Hello "
	str2 := "World"
	result := concatenate(&str1, &str2)
	fmt.Println(result)
}
