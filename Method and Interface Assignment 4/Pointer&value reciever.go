// Pointer and Value Receivers

package main

import "fmt"

type Author struct {
	name  string
	books int
}

func (a Author) Display() {
	fmt.Println("Author:", a.name)
	fmt.Println("Number of Books:", a.books)
}

func (a *Author) UpdateBooks(count int) {
	a.books = count
}

func main() {
	author := Author{name: "Mark Twain", books: 12}

	author.Display()

	author.UpdateBooks(15)

	author.Display()
}
