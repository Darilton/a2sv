package main

import (
	"fmt"
	"library_management/controllers"
)

func main() {
	controllers.Init()
	opt := 1
	for opt != 7 {
		fmt.Println("*********Lib, Library Management System*********")
		fmt.Println("Please choose an option")
		fmt.Println("1 Add a new book")
		fmt.Println("2 Remove an existing book")
		fmt.Println("3 Borrow a book")
		fmt.Println("4 Return a book")
		fmt.Println("5 List all available books")
		fmt.Println("6 List all borrowed books by a member")
		fmt.Println("7 Exit")
		fmt.Print("opt: ")
		fmt.Scanf("%d", &opt)
		switch opt {
		case 1: controllers.AddNewBook()
		case 5: controllers.ListAvailableBooks()
		}
	}
}
