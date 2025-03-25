// 1. WAP to create student struct with student name and marks and sort it based on student marks using sort package

package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name  string
	Marks int
}

type ByMarks []Student

func (s ByMarks) Len() int           { return len(s) }
func (s ByMarks) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByMarks) Less(i, j int) bool { return s[i].Marks < s[j].Marks }

func main() {
	students := []Student{
		{"Alice", 85},
		{"Bob", 92},
		{"Charlie", 78},
		{"Dave", 89},
	}

	sort.Sort(ByMarks(students))
	fmt.Println("Students sorted by marks:")
	for _, s := range students {
		fmt.Printf("Name: %s, Marks: %d\n", s.Name, s.Marks)
	}
}
