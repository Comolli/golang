package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}

func changeBook1 (book * Books){
   book.title = "title_change1"
}

func changeBook2 (book *Books) {
   book.title = "title_change2"
}

func main() {
   var book1 Books;
   book1.title = "book1"
   book1.author = "hoult"
   book1.book_id = 1
   changeBook1(&book1)
   fmt.Println(book1)
   changeBook2(&book1)
   fmt.Println(book1)
}

//out:
//{title_change1 hoult  1}
//{title_change2 hoult  1}
