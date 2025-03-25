//  Print Elements Using a for Loop

package main

import "fmt"

func main() {
    numbers := [5]int{10, 20, 30, 40, 50}

    for i := 0; i < len(numbers); i++ {
        fmt.Println("Element at index", i, ":", numbers[i])
    }
}
