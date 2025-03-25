//  WAP in go language to accept n student details like roll_no, stud_name, mark1,mark2, mark3. Calculate the total and average of marks using structure.

package main

import "fmt"

type Student struct {
    RollNo   int
    Name     string
    Marks    [3]float64
    Total    float64
    Average  float64
}

func main() {
    var n int
    fmt.Print("Enter number of students: ")
    fmt.Scan(&n)

    students := make([]Student, n)

    for i := 0; i < n; i++ {
        fmt.Printf("Enter details for Student %d\n", i+1)
        fmt.Print("Roll No: ")
        fmt.Scan(&students[i].RollNo)
        fmt.Print("Name: ")
        fmt.Scan(&students[i].Name)
        fmt.Print("Enter 3 marks: ")
        fmt.Scan(&students[i].Marks[0], &students[i].Marks[1], &students[i].Marks[2])

        students[i].Total = students[i].Marks[0] + students[i].Marks[1] + students[i].Marks[2]
        students[i].Average = students[i].Total / 3
    }

    fmt.Println("\nStudent Details:")
    for _, stud := range students {
        fmt.Printf("RollNo: %d, Name: %s, Total: %.2f, Average: %.2f\n", stud.RollNo, stud.Name, stud.Total, stud.Average)
    }
}
