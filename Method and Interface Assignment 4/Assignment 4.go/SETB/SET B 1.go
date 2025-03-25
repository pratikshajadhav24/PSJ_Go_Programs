// 1. Write a program in go language to create structure student. Write a method show() whose receiver is a pointer of struct student.

package main

import "fmt"

type Student struct {
	RollNo int
	Name   string
	Marks  float64
}

func (s *Student) Show() {
	fmt.Printf("Roll No: %d\nName: %s\nMarks: %.2f\n", s.RollNo, s.Name, s.Marks)
}

func main() {
	s := Student{RollNo: 101, Name: "John Doe", Marks: 89.5}
	s.Show()
}
