package mocks

import (
	domain "task_manager_api/Domain"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUser(username string) (domain.User, error) {
	args := m.Called(username)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserRepository) AddUser(newUser domain.User) error {
	args := m.Called(newUser)
	return args.Error(0)
}
