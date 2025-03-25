// WAP in go language to sort array elements in ascending order.

package main

import (
    "fmt"
    "sort"
)

func main() {
    arr := []int{42, 12, 5, 78, 23}
    sort.Ints(arr)
    fmt.Println("Sorted Array:", arr)
}
