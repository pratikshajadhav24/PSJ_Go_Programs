


package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Element at index %d: %d\n", i, arr[i])
	}
}
