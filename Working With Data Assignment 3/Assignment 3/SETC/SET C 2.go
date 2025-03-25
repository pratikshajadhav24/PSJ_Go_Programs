//  WAP in go language to accept n records of employee information (eno,ename,salary) and display record of employees having maximum salary.

package main

import "fmt"

type Employee struct {
    ENo    int
    EName  string
    Salary float64
}

func main() {
    var n int
    fmt.Print("Enter number of employees: ")
    fmt.Scan(&n)

    employees := make([]Employee, n)
    var maxSalary float64

    for i := 0; i < n; i++ {
        fmt.Printf("Enter details for Employee %d\n", i+1)
        fmt.Print("Employee No: ")
        fmt.Scan(&employees[i].ENo)
        fmt.Print("Name: ")
        fmt.Scan(&employees[i].EName)
        fmt.Print("Salary: ")
        fmt.Scan(&employees[i].Salary)

        if employees[i].Salary > maxSalary {
            maxSalary = employees[i].Salary
        }
    }

    fmt.Println("\nEmployees with Maximum Salary:")
    for _, emp := range employees {
        if emp.Salary == maxSalary {
            fmt.Printf("ENo: %d, Name: %s, Salary: %.2f\n", emp.ENo, emp.EName, emp.Salary)
        }
    }
}

