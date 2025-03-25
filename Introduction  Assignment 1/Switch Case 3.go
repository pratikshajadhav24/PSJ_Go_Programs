// Check Whether a Character is a Vowel or Consonant

package main

import "fmt"

func main() {
	var char string
	fmt.Print("Enter a character: ")
	fmt.Scanln(&char)

	switch char {
	case "a", "e", "i", "o", "u":
		fmt.Println("Vowel")
	default:
		fmt.Println("Consonant")
	}
}

