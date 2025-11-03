package services

import (
	"errors"
	"fmt"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	ListAvailableBooks() []models.Book
	BorrowBook(bookID int, memberID int) error
}

type Library struct {
	Books  map[int]models.Book
	Member map[int]models.Member
}

func (l Library) BorrowBook(bookID int, memberID int) error {
	book, ok := l.Books[bookID]
	if !ok {
		return errors.New("book not found in library")
	}

	if book.Status == "Borrowed" {
		return errors.New("book is currently not available")
	}

	member, ok := l.Member[memberID]
	if !ok {
		return errors.New("member not found")
	}

	book.Status = "Borrowed"
	l.Books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)

	l.Member[memberID] = member
	return nil
}

func (l Library) AddBook(book models.Book) {
	l.Books[book.ID] = book
}

func (l Library) AddMember(member models.Member) {
	l.Member[member.ID] = member
}

func (l Library) ListAvailableBooks() []models.Book {
	ans := make([]models.Book, 0)
	for _, book := range l.Books {
		if book.Status == "Available" {
			ans = append(ans, book)
		}
	}
	return ans
}

func (l Library) ListMembers() []models.Member {
	ans := make([]models.Member, 0)
	fmt.Println(l.Member)
	for _, member := range l.Member {
		ans = append(ans, member)
	}
	return ans
}
