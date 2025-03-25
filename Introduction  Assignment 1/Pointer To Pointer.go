

package main

import "fmt"

func main() {
	var a int = 10
	var p *int = &a
	var pp **int = &p
	fmt.Println("Value of a:", a)
	fmt.Println("Value at pointer p:", *p)
	fmt.Println("Value at pointer to pointer pp:", **pp)
}
