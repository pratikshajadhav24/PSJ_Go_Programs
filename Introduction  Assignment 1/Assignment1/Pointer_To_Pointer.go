// WAP in go language to illustrate pointer to pointer

package main

import "fmt"

func main() {
	var a int = 10
	var ptr *int = &a
	var ptr2 **int = &ptr

	fmt.Println("Value of a:", a)
	fmt.Println("Value using ptr:", *ptr)
	fmt.Println("Value using ptr2:", **ptr2)
}
