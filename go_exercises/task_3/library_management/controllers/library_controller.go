package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strings"
)

var lib services.Library

func Init() {
	lib.Books = make(map[int]models.Book)
	lib.Member = make(map[int]models.Member)
}

func ListAvailableBooks() {
	availableBooks := lib.ListAvailableBooks()

	fmt.Println("*********Available Books Listing Menu*********")

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

func BorrowBook() {
	var bookId, memberId int

	fmt.Println("*********Borrow Book Menu*********")

	fmt.Print("Book Id: ")
	fmt.Scanf("%d", &bookId)

	fmt.Print("Member Id: ")
	fmt.Scanf("%d", &memberId)

	err := lib.BorrowBook(bookId, memberId)
	if err != nil {
		fmt.Println("Failure: ", err)
	} else {
		fmt.Println("Book Borrowed Successfuly!")
	}
}

func ListMembers() {
	members := lib.ListMembers()
	title := "Available Members:"

	fmt.Println("*********Member Listing Menu*********")

	if len(members) == 0 {
		title = "Sorry, there are no registered members."
	}
	fmt.Println(title)
	for _, book := range members {
		fmt.Println("Member Id: ", book.ID)
		fmt.Println("Member Name: ", book.Name)
		fmt.Println()
	}
}

func AddNewBook() {
	var book models.Book
	buffer := bufio.NewReader(os.Stdin)

	fmt.Println("*********Add Book Menu*********")
	fmt.Print("Book Id: ")
	fmt.Scanf("%d", &book.ID)
	fmt.Print("Book Title: ")
	book.Title, _ = buffer.ReadString('\n')
	book.Title = strings.TrimSpace(book.Title)
	fmt.Print("Book Author: ")
	book.Author, _ = buffer.ReadString('\n')
	book.Author = strings.TrimSpace(book.Author)
	book.Status = "Available"
	lib.AddBook(book)

	fmt.Println("Book Added Successfuly!")
	fmt.Println()
}

func AddNewMember() {
	var member models.Member
	buffer := bufio.NewReader(os.Stdin)
	fmt.Println("*********Add Member Menu*********")
	fmt.Print("Member Id: ")
	fmt.Scanf("%d", &member.ID)
	fmt.Print("Member Name: ")
	member.Name, _ = buffer.ReadString('\n')
	member.Name = strings.TrimSpace(member.Name)

	lib.AddMember(member)
	fmt.Println("Member Added Successfuly!")
	fmt.Println()
}

func ListBorrowedBooks() {
	var memberID int
	fmt.Println("*********Borrowed Books Listing Menu*********")
	fmt.Print("Member Id: ")
	fmt.Scanf("%d", &memberID)

	borrowedBooks := lib.ListBorrowedBooks(memberID)
	if borrowedBooks == nil {
		fmt.Println("No books borowed by given member")
	} else {
		fmt.Println("*********Borrowed Books by Given Member*********")
	}
	for _, book := range borrowedBooks {
		fmt.Println("Book Id: ", book.ID)
		fmt.Println("Book Title: ", book.Title)
		fmt.Println("Book Author: ", book.Author)
		fmt.Println()
	}
}
