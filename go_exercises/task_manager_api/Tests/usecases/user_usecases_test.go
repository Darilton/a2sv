package usecases_test

import (
	"errors"
	domain "task_manager_api/Domain"
	"task_manager_api/Infrastructure"
	"task_manager_api/Tests/mocks"
	usecases "task_manager_api/UseCases"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUser(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userUseCase := usecases.NewUserUseCase(mockRepo)

	user := domain.User{UserName: "test", Password: "password"}

	mockRepo.On("AddUser", mock.Anything).Return(nil)

	err := userUseCase.RegisterUser(user)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestLoginUser_Success(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userUseCase := usecases.NewUserUseCase(mockRepo)

	password := "password"
	hashedPassword, _ := Infrastructure.HashPassword(password)
	user := domain.User{UserName: "test", Password: hashedPassword, UserRole: "regular"}

	mockRepo.On("GetUser", "test").Return(user, nil)

	token, err := userUseCase.LoginUser(domain.User{UserName: "test", Password: password})

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestLoginUser_InvalidPassword(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userUseCase := usecases.NewUserUseCase(mockRepo)

	password := "password"
	hashedPassword, _ := Infrastructure.HashPassword(password)
	user := domain.User{UserName: "test", Password: hashedPassword}

	mockRepo.On("GetUser", "test").Return(user, nil)

	token, err := userUseCase.LoginUser(domain.User{UserName: "test", Password: "wrongpassword"})

	assert.Error(t, err)
	assert.Empty(t, token)
	assert.Equal(t, "invalid Username or Password", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestLoginUser_UserNotFound(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userUseCase := usecases.NewUserUseCase(mockRepo)

	mockRepo.On("GetUser", "unknown").Return(domain.User{}, errors.New("user not found"))

	token, err := userUseCase.LoginUser(domain.User{UserName: "unknown", Password: "password"})

	assert.Error(t, err)
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
}
