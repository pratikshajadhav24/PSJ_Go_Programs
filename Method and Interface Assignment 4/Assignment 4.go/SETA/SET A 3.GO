// 3. Write a program in go language to create structure author. Write a method show() whose receiver is structauthor.

package main

import "fmt"

type Author struct {
	Name    string
	Country string
}

func (a Author) show() {
	fmt.Println("Author Name:", a.Name)
	fmt.Println("Country:", a.Country)
}

func main() {

	author1 := Author{Name: "J.K. Rowling", Country: "United Kingdom"}

	author1.show()
}
