// . Using a for Loop with an Indexpackage main

import "fmt"

func main() {
    // Define a slice of integers
    numbers := []int{10, 20, 30, 40, 50}

    // Iterate using a traditional for loop
    fmt.Println("Using for loop with index:")
    for i := 0; i < len(numbers); i++ {
        fmt.Println("Index:", i, "Value:", numbers[i])
    }
}


