package Domain

import "time"

type Task struct {
	Id          string    `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

type User struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	UserRole string `json:"-"`
}
