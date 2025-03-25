

package main

import "fmt"

func newInt(value int) *int {
	p := new(int)
	*p = value
	return p
}

func main() {
	p := newInt(10)
	fmt.Println("Value at pointer p:", *p)
}
