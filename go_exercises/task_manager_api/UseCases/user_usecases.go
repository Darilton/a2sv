package usecases

import (
	"errors"
	domain "task_manager_api/Domain"
	"task_manager_api/Infrastructure"
	"task_manager_api/Repositories"
)

func RegisterUser(newUser domain.User) error {
	hashedPassword, err := Infrastructure.HashPassword(newUser.Password)
	if err != nil {
		return err
	}

	newUser.Password = string(hashedPassword)
	return Repositories.AddUser(newUser)
}

func LoginUser(loginUser domain.User) (string, error) {
	user, err := Repositories.GetUser(loginUser.UserName)
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
