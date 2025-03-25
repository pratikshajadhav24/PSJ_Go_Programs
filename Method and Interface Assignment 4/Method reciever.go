// method with  reciever of author type

package main

import "fmt"

type Author struct {
	name  string
	books int
}

func (a Author) Show() {
	fmt.Println("Author:", a.name)
	fmt.Println("Number of Books:", a.books)
}

func main() {
	author := Author{name: "J.K. Rowling", books: 7}
	author.Show()
}
