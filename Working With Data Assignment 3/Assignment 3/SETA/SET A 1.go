//WAP in go language to find the largest and smallest number in an array.

package main

import "fmt"

func findMinMax(arr []int) (int, int) {
	min, max := arr[0], arr[0]
	for _, num := range arr {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min, max
}

func main() {
	arr := []int{10, 5, 20, 8, 25, 3}
	min, max := findMinMax(arr)
	fmt.Println("Smallest Number:", min)
	fmt.Println("Largest Number:", max)
}
