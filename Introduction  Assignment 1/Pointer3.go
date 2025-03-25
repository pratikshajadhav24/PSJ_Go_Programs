

package main

import "fmt"

func main() {
	var p *int
	fmt.Println("Pointer p:", p)
	if p == nil {
		fmt.Println("p is a nil pointer")
	}
}
