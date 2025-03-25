//Search for an Element in an Array

package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	var key int
	fmt.Print("Enter element to search: ")
	fmt.Scanln(&key)
	found := false
	for i, value := range arr {
		if value == key {
			fmt.Printf("Element found at index %d\n", i)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Element not found")
	}
}
