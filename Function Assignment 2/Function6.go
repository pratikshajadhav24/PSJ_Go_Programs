//  Function as a Return Value

package main

import "fmt"

// Function returning another function
func multiplier(factor int) func(int) int {
    return func(num int) int {
        return num * factor
    }
}

func main() {
    double := multiplier(2)
    fmt.Println("Double of 4:", double(4))
}
