// Structures as Function Arguments

package main
import "fmt"
type Books struct {
title string
author string
subject string
book_id int
}
func main() {
var Book1 Books /* Declare Book1 of type Book */
var Book2 Books /* Declare Book2 of type Book */
/* book 1 specification */
Book1.title = "Go Programming"
Book1.author = "ABC"
Book1.subject = "Go Programming Tutorial"
Book1.book_id = 6495407
/* book 2 specification */
Book2.title = "Java Programming"
Book2.author = "Herbert Schildt"
Book2.subject = "Java Programming Tutorial"
Book2.book_id = 6495700
/* print Book1 info */
printBook(Book1)
/* print Book2 info */
printBook(Book2)
}
func printBook( book Books ) {
fmt.Printf( "Book title : %s\n", book.title);
fmt.Printf( "Book author : %s\n", book.author);
fmt.Printf( "Book subject : %s\n", book.subject);
fmt.Printf( "Book book_id : %d\n", book.book_id);
}
