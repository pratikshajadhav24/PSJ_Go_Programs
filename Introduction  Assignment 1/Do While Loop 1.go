// Print Numbers from 1 to 5

package main

import "fmt"

func main() {
	i := 1
	for {
		fmt.Println(i)
		i++
		if i > 5 {
			break
		}
	}
}
