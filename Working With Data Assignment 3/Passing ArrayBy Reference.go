package main

import "fmt"

// Function that takes a pointer to an array
func modifyArray(arr *[5]int) {
    arr[0] = 100 // Modify first element
    fmt.Println("Inside modifyArray:", *arr)
}

func main() {
    nums := [5]int{1, 2, 3, 4, 5}
    
    modifyArray(&nums) // Pass address of array

    fmt.Println("After modifyArray:", nums) // Original array is changed
}
