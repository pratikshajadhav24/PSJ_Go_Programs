//Access Elements in an Array

package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	fmt.Println("First element:", arr[0])
	fmt.Println("Third element:", arr[2])
	fmt.Println("Last element:", arr[len(arr)-1])
}
