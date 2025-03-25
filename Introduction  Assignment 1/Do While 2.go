// Print Even Numbers Less Than 20

package main

import "fmt"

func main() {
	i := 2
	for {
		fmt.Println(i)
		i += 2
		if i > 20 {
			break
		}
	}
}
