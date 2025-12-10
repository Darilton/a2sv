package usecases_test

import (
	"errors"
	domain "task_manager_api/Domain"
	"task_manager_api/Tests/mocks"
	usecases "task_manager_api/UseCases"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddTask(t *testing.T) {
	mockRepo := new(mocks.MockTaskRepository)
	taskUseCase := usecases.NewTaskUseCase(mockRepo)

	task := domain.Task{Id: "1", Title: "Test Task", Description: "Desc", DueDate: time.Now(), Status: "Pending"}

	mockRepo.On("AddTask", task).Return(nil)

	err := taskUseCase.AddTask(task)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetTasks(t *testing.T) {
	mockRepo := new(mocks.MockTaskRepository)
	taskUseCase := usecases.NewTaskUseCase(mockRepo)

	tasks := []domain.Task{
		{Id: "1", Title: "Task 1"},
		{Id: "2", Title: "Task 2"},
	}

	mockRepo.On("GetTasks").Return(tasks)

	result := taskUseCase.GetTasks()

	assert.Equal(t, tasks, result)
	mockRepo.AssertExpectations(t)
}

func TestGetTask_Success(t *testing.T) {
	mockRepo := new(mocks.MockTaskRepository)
	taskUseCase := usecases.NewTaskUseCase(mockRepo)

	task := domain.Task{Id: "1", Title: "Task 1"}

	mockRepo.On("GetTask", "1").Return(task, nil)

	result, err := taskUseCase.GetTask("1")

	assert.NoError(t, err)
	assert.Equal(t, task, result)
	mockRepo.AssertExpectations(t)
}

func TestGetTask_NotFound(t *testing.T) {
	mockRepo := new(mocks.MockTaskRepository)
	taskUseCase := usecases.NewTaskUseCase(mockRepo)

	mockRepo.On("GetTask", "999").Return(domain.Task{}, errors.New("task not found"))

	_, err := taskUseCase.GetTask("999")

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestEditTask(t *testing.T) {
	mockRepo := new(mocks.MockTaskRepository)
	taskUseCase := usecases.NewTaskUseCase(mockRepo)

	task := domain.Task{Id: "1", Title: "Updated Task"}

	mockRepo.On("EditTask", "1", task).Return(nil)

	err := taskUseCase.EditTask("1", task)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTask(t *testing.T) {
	mockRepo := new(mocks.MockTaskRepository)
	taskUseCase := usecases.NewTaskUseCase(mockRepo)

	mockRepo.On("DeleteTask", "1").Return(nil)

	err := taskUseCase.DeleteTask("1")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
