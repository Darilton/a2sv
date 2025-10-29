package services

import "library_management/models"

type LibraryManager interface {
	AddBook(book models.Book)
	ListAvailableBooks() []models.Book
}

type Library struct {
	Books map[int]models.Book
}

func (l Library) AddBook(book models.Book) {
	l.Books[book.ID] = book
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