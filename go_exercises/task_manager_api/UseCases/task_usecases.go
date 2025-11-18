package usecases

import (
	domain "task_manager_api/Domain"
	"task_manager_api/Repositories"
)

func EditTask(id string, newTask domain.Task) error {
	return Repositories.EditTask(id, newTask)
}

func DeleteTask(id string) error {
	return Repositories.DeleteTask(id)
}

func AddTask(newTask domain.Task) error {
	return Repositories.AddTask(newTask)
}

func GetTask(id string) (domain.Task, error) {
	return Repositories.GetTask(id)
}

func GetTasks() []domain.Task {
	return Repositories.GetTasks()
}
