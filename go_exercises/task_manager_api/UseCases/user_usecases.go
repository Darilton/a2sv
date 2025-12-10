package usecases

import (
	"errors"
	domain "task_manager_api/Domain"
	"task_manager_api/Infrastructure"
	"task_manager_api/Repositories"
)

type UserUseCase struct {
	userRepo Repositories.UserRepository
}

func NewUserUseCase(userRepo Repositories.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (uu *UserUseCase) RegisterUser(newUser domain.User) error {
	hashedPassword, err := Infrastructure.HashPassword(newUser.Password)
	if err != nil {
		return err
	}

	newUser.Password = string(hashedPassword)
	return uu.userRepo.AddUser(newUser)
}

func (uu *UserUseCase) LoginUser(loginUser domain.User) (string, error) {
	user, err := uu.userRepo.GetUser(loginUser.UserName)
	if err != nil {
		return "", err
	}

	if !Infrastructure.CheckPasswordHash(loginUser.Password, user.Password) {
		return "", errors.New("invalid Username or Password")
	}

	token, err := Infrastructure.GenerateJWT(user.UserName, user.UserRole)
	if err != nil {
		return "", err
	}
	return token, nil
}
