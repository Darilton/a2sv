package models

import "sync"

type Book struct {
	ID int
	Title string
	Author string
	Status string
	Mu	*sync.Mutex
}
