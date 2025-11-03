package main

import (
	"fmt"
	"library_management/controllers"
)

func main() {
	controllers.Init()
	opt := 1
	for opt != 8 {
		fmt.Println("*********Lib, Library Management System*********")
		fmt.Println("Please choose an option")
		fmt.Println("1 Add a new book")
		fmt.Println("2 Add a new Member")
		fmt.Println("3 Remove an existing book")
		fmt.Println("4 Borrow a book")
		fmt.Println("5 Return a book")
		fmt.Println("6 List all available books")
		fmt.Println("7 List all borrowed books by a member")
		fmt.Println("8 Exit")
		fmt.Print("opt: ")
		fmt.Scanf("%d", &opt)
		switch opt {
		case 1: controllers.AddNewBook()
		case 2: controllers.AddNewMember()
		case 5: controllers.ListAvailableBooks()
		}
	}
}
