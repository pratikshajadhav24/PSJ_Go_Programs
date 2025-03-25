// Print Odd Numbers Between 1 and 15

package main

import "fmt"

func main() {
	i := 1
	for {
		if i%2 != 0 {
			fmt.Println(i)
		}
		i++
		if i > 15 {
			break
		}
	}
}
