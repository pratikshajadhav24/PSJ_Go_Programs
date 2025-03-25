// 2. Write a program in go language to store n student information(rollno, name, percentage) and write a method to display student information in descending order of percentage.

package main

import (
	"fmt"
	"sort"
)

type Student struct {
	RollNo     int
	Name       string
	Percentage float64
}

type StudentList []Student

func (s StudentList) DisplayDescending() {
	sort.Slice(s, func(i, j int) bool {
		return s[i].Percentage > s[j].Percentage
	})
	for _, student := range s {
		fmt.Printf("Roll No: %d, Name: %s, Percentage: %.2f\n", student.RollNo, student.Name, student.Percentage)
	}
}

func main() {
	var n int
	fmt.Print("Enter number of students: ")
	fmt.Scan(&n)

	students := make(StudentList, n)
	for i := 0; i < n; i++ {
		fmt.Print("Enter Roll No, Name and Percentage: ")
		fmt.Scan(&students[i].RollNo, &students[i].Name, &students[i].Percentage)
	}

	fmt.Println("\nStudent Information in Descending Order of Percentage:")
	students.DisplayDescending()
}
