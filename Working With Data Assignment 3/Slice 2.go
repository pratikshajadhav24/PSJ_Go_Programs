//Assigning a array element in an array  in slicepackage main

import "fmt"

func main() {
    // Declare an array
    arr := [5]int{1, 2, 3, 4, 5}

    // Create a slice from the array
    slice := arr[1:4] // Slice contains elements [2, 3, 4]

    fmt.Println("Original array:", arr)  // [1 2 3 4 5]
    fmt.Println("Slice:", slice)         // [2 3 4]

    // Assign a new value to an element in the slice
    slice[0] = 10

    fmt.Println("Modified array:", arr)  // [1 10 3 4 5]
    fmt.Println("Modified slice:", slice) // [10 3 4]
}


