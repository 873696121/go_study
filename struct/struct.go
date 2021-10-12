package main

import "fmt"

type Book struct {
	title  string
	author string
}

func changeBook1(book Book) {
	book.author = "lisi"
}

func changeBook2(book *Book) {
	book.author = "lisi"
}

func main() {
	var book1 Book
	book1.title = "Golang"
	book1.author = "zhangsan"
	fmt.Println(book1)

	changeBook1(book1)
	fmt.Println(book1)

	changeBook2(&book1)
	fmt.Println(book1)
}
