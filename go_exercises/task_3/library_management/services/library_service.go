package services

import "library_management/models"

type LibraryManager interface {
	AddBook(book models.Book)
}

type Library struct {
	Books map[int]models.Book
}

func (l Library) AddBook(book models.Book) {
	l.Books[book.ID] = book
}
