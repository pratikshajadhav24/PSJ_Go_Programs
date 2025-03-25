//Copying Elements from One Slice to Another

package main

import "fmt"

func main() {
    // Create the source slice
    source := []int{10, 20, 30, 40, 50}

    // Create a destination slice with enough space to hold the elements
    destination := make([]int, len(source))

    // Copy elements from source to destination
    n := copy(destination, source)

    // Print the destination slice
    fmt.Println("Destination Slice:", destination) // Output: [10 20 30 40 50]
    fmt.Println("Number of elements copied:", n)   // Output: 5
}
