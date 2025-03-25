// WAP in go Language to print address of a variable.

package main

import "fmt"

func main() {
	var num int = 42
	fmt.Printf("Address of num: %p\n", &num)
}
