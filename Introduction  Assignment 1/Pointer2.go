

package main

import "fmt"

func main() {
	var a int = 10
	var p *int = &a
	*p = 20
	fmt.Println("Modified value of a:", a)
}
