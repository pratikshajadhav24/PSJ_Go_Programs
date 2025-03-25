

package main

import "fmt"

func increment(ptr *int) {
	*ptr += 1
}

func main() {
	var a int = 5
	increment(&a)
	fmt.Println("Value after increment:", a)
}
