//   Check if a Year is a Leap Year

package main

import "fmt"

func main() {
	var year int
	fmt.Print("Enter a year: ")
	fmt.Scanln(&year)

	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Println("The year is a leap year.")
	} else {
		fmt.Println("The year is not a leap year.")
	}
}
