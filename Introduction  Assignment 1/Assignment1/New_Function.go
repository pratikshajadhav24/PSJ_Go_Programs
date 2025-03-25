// WAP in go language to explain new function

package main

import "fmt"

func main() {
	ptr := new(int)
	*ptr = 42
	fmt.Println("Value using new():", *ptr)
}
