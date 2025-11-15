package services

import (
	"errors"
	"fmt"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	AddMember(member models.Member)
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
	ListReservedBooks(memberID int) []models.Book
	BorrowBook(bookID int, memberID int) error
	ReturnBook(memberID, bookID int) error
	ReserveBook(bookID int, memberID int) error
	UnReserveBook(booID int, memberID int)
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
	fmt.Println(member.BorrowedBooks)
	l.Member[memberID] = member
	return nil
}

func (l Library) AddBook(book models.Book) {
	l.Books[book.ID] = book
}

func (l Library) AddMember(member models.Member) {
	l.Member[member.ID] = member
}

func (l Library) RemoveBook(bookID int) {
	delete(l.Books, bookID)
}

func (l Library) ReserveBook(bookID int, memberID int) error {
	book, ok := l.Books[bookID]
	if !ok {
		return errors.New("book not found in library")
	}

	if book.Status == "Borrowed" || book.Status == "Reserved" {
		return errors.New("book is currently not available")
	}

	member, ok := l.Member[memberID]
	if !ok {
		return errors.New("member not found")
	}
	book.Status = "Reserved"
	l.Books[bookID] = book
	member.ReservedBooks = append(member.ReservedBooks, book)
	fmt.Println(member.ReservedBooks)
	l.Member[memberID] = member
	return nil
}

func (l Library) UnReserveBook(bookID int, memberID int) error {
	member, ok := l.Member[memberID]
	if !ok {
		return errors.New("member not found")
	}

	book_idx := -1 // flag as not found
	for i := range len(member.ReservedBooks) {
		book := member.ReservedBooks[i]
		if book.ID == bookID {
			book_idx = i
		}
	}

	if book_idx < 0 {
		return errors.New("book not Reserved by given member")
	}

	if member.ReservedBooks[book_idx].Status != "Reserved" {
		return errors.New("book not in Reserved Status")
	}
	book := member.ReservedBooks[book_idx]
	book.Status = "Available"
	l.Books[bookID] = book
	member.ReservedBooks = append(member.ReservedBooks[:book_idx], member.ReservedBooks[book_idx+1:]...)
	l.Member[memberID] = member
	return nil
}

func (l Library) ReturnBook(memberID, bookID int) error {
	member, ok := l.Member[memberID]
	if !ok {
		return errors.New("member not found")
	}

	book_idx := -1 // flag as not found
	for i := range len(member.BorrowedBooks) {
		book := member.BorrowedBooks[i]
		if book.ID == bookID {
			book_idx = i
		}
	}
	if book_idx < 0 {
		return errors.New("book not borowed by given member")
	}

	book := member.BorrowedBooks[book_idx]
	book.Status = "Available"
	l.Books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks[:book_idx], member.BorrowedBooks[book_idx+1:]...)
	l.Member[memberID] = member
	return nil
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
	for _, member := range l.Member {
		ans = append(ans, member)
	}
	return ans
}

func (l Library) ListReservedBooks(memberID int) []models.Book {
	member, ok := l.Member[memberID]
	if !ok {
		return nil
	}
	return member.ReservedBooks
}

func (l Library) ListBorrowedBooks(memberID int) []models.Book {
	member, ok := l.Member[memberID]
	if !ok {
		return nil
	}
	return member.BorrowedBooks
}
