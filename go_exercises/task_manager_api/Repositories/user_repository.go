package Repositories

import (
	"context"
	"errors"
	domain "task_manager_api/Domain"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository interface {
	GetUser(username string) (domain.User, error)
	AddUser(newUser domain.User) error
}

type UserRepositoryMongo struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) UserRepository {
	return &UserRepositoryMongo{
		collection: collection,
	}
}

func (ur *UserRepositoryMongo) GetUser(username string) (domain.User, error) {
	var user domain.User
	err := ur.collection.FindOne(context.TODO(), bson.D{{Key: "username", Value: username}}).Decode(&user)
	if err != nil {
		return user, errors.New("user Not found")
	}
	return user, nil
}

func (ur *UserRepositoryMongo) AddUser(newUser domain.User) error {
	if newUser.Password == "" || newUser.UserName == "" {
		return errors.New("invalid Request")
	}
	if ur.collection.FindOne(context.TODO(), bson.D{{Key: "username", Value: newUser.UserName}}).Decode(&domain.User{}) != mongo.ErrNoDocuments {
		return errors.New("username already exists")
	}
	// The first user to register is an admin
	if curr, err := ur.collection.Find(context.TODO(), bson.D{{}}); err == nil && !curr.Next(context.TODO()) {
		newUser.UserRole = "admin"
	} else {
		newUser.UserRole = "regular"
	}

	if _, err := ur.collection.InsertOne(context.TODO(), newUser); err != nil {
		return err
	}
	return nil
}
