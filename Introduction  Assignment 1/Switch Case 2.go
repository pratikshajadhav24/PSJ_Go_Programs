//  Check Grade

package main

import "fmt"

func main() {
	var score int
	fmt.Print("Enter your score: ")
	fmt.Scanln(&score)

	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	case score >= 60:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: F")
	}
}
