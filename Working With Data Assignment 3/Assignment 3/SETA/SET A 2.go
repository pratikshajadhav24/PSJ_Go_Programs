// WAP in go language to accept the book details such as BookID, Title, Author, Price. Read and display the details of n number of books.

package main

import "fmt"

type Book struct {
    BookID int
    Title  string
    Author string
    Price  float64
}

func main() {
    var n int
    fmt.Print("Enter number of books: ")
    fmt.Scan(&n)

    books := make([]Book, n)

    for i := 0; i < n; i++ {
        fmt.Printf("Enter details for Book %d\n", i+1)
        fmt.Print("BookID: ")
        fmt.Scan(&books[i].BookID)
        fmt.Print("Title: ")
        fmt.Scan(&books[i].Title)
        fmt.Print("Author: ")
        fmt.Scan(&books[i].Author)
        fmt.Print("Price: ")
        fmt.Scan(&books[i].Price)
    }

    fmt.Println("\nBook Details:")
    for _, book := range books {
        fmt.Printf("ID: %d, Title: %s, Author: %s, Price: %.2f\n", book.BookID, book.Title, book.Author, book.Price)
    }
}
