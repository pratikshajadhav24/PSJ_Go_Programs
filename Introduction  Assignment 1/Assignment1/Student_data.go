//   WAP in go language to print Student name, rollno,division and college name

package main

import "fmt"

func main() {
	var name, college string
	var rollNo int
	var division string
	
	name = "John Doe"
	rollNo = 12345
	division = "A"
	college = "XYZ College"
	
	fmt.Println("Student Name:", name)
	fmt.Println("Roll Number:", rollNo)
	fmt.Println("Division:", division)
	fmt.Println("College Name:", college)
}
