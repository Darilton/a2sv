package usecases

import (
	domain "task_manager_api/Domain"
	"task_manager_api/Repositories"
)

type TaskUseCase struct {
	taskRepo Repositories.TaskRepository
}

func NewTaskUseCase(taskRepo Repositories.TaskRepository) *TaskUseCase {
	return &TaskUseCase{
		taskRepo: taskRepo,
	}
}

func (tu *TaskUseCase) EditTask(id string, newTask domain.Task) error {
	return tu.taskRepo.EditTask(id, newTask)
}

func (tu *TaskUseCase) DeleteTask(id string) error {
	return tu.taskRepo.DeleteTask(id)
}

func (tu *TaskUseCase) AddTask(newTask domain.Task) error {
	return tu.taskRepo.AddTask(newTask)
}

func (tu *TaskUseCase) GetTask(id string) (domain.Task, error) {
	return tu.taskRepo.GetTask(id)
}

func (tu *TaskUseCase) GetTasks() []domain.Task {
	return tu.taskRepo.GetTasks()
}
