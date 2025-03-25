// Print Even Numbers Between 1 and 20

package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
