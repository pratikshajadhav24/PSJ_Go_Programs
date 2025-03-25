// User Input Until "exit" is Entered

package main

import "fmt"

func main() {
	var input string
	for {
		fmt.Print("Enter something (type 'exit' to quit): ")
		fmt.Scanln(&input)
		if input == "exit" {
			break
		}
		fmt.Println("You entered:", input)
	}
}
