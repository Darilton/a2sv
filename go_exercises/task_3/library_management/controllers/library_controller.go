package controllers

import "fmt"
import "library_management/models"
import "library_management/services"

var lib services.Library

func Init() {
	lib.Books = make(map[int]models.Book)
	lib.Member = make(map[int]models.Member)
}

func ListAvailableBooks() {
	availableBooks := lib.ListAvailableBooks()
	title := "Available Books:"
	if len(availableBooks) == 0 {
		title = "Sorry, there are no books available."
	}
	fmt.Println(title)
	for _, book := range availableBooks {
		fmt.Println("Book Id: ", book.ID)
		fmt.Println("Book Title: ", book.Title)
		fmt.Println("Book Author: ", book.Author)
		fmt.Println()
	}
}

func AddNewBook() {
	var book models.Book
	
	fmt.Println("*********Add Book Menu*********")
	fmt.Print("Book Id: ")
	fmt.Scanf("%d", &book.ID)
	fmt.Print("Book Title: ")
	fmt.Scanln(&book.Title)
	fmt.Print("Book Author: ")
	fmt.Scanln(&book.Author)
	book.Status = "Available"
	lib.AddBook(book)

	fmt.Println("Book Added Successfuly!")
	fmt.Println()
}