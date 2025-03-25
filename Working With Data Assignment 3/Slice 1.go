package main

import "fmt"

func main() {
    // Declare and initialize a slice with size 5
    nums := []int{10, 20, 30, 40, 50}

    fmt.Println("Slice:", nums)       // [10 20 30 40 50]
    fmt.Println("Length:", len(nums)) // 5
    fmt.Println("Capacity:", cap(nums)) // 5
}
