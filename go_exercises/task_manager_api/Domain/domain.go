package Domain

import "time"

type Task struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	UserRole string `json:"-"`
}
