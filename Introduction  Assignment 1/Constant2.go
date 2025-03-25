//   Multiple Constants 

package main

import "fmt"

func main() {
    const (
        first  = iota // 0
        second        // 1
        third         // 2
    )

    fmt.Println("First:", first)
    fmt.Println("Second:", second)
    fmt.Println("Third:", third)
}
