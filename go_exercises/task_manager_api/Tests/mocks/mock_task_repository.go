package mocks

import (
	domain "task_manager_api/Domain"

	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) AddTask(newTask domain.Task) error {
	args := m.Called(newTask)
	return args.Error(0)
}

func (m *MockTaskRepository) GetTask(id string) (domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskRepository) GetTasks() []domain.Task {
	args := m.Called()
	return args.Get(0).([]domain.Task)
}

func (m *MockTaskRepository) EditTask(id string, newTask domain.Task) error {
	args := m.Called(id, newTask)
	return args.Error(0)
}

func (m *MockTaskRepository) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
