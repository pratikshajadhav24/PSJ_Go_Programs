// 3. Write a program in go language to copy all elements of one array into another using method.

package main

import "fmt"

type ArrayCopy struct{}

func (ArrayCopy) Copy(src []int) []int {
	dest := make([]int, len(src))
	copy(dest, src)
	return dest
}

func main() {
	source := []int{1, 2, 3, 4, 5}
	ac := ArrayCopy{}
	destination := ac.Copy(source)

	fmt.Println("Source Array:", source)
	fmt.Println("Destination Array:", destination)
}
