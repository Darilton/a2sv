package data

import (
	"context"
	"errors"
	"task_manager_api/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var userColl *mongo.Collection

func SetUserCollection(collection *mongo.Collection) {
	userColl = collection
}

func AddUser(newUser models.User) error {
	if newUser.Password == "" || newUser.UserName == "" {
		return errors.New("invalid Request")
	}
	if userColl.FindOne(context.TODO(), bson.D{{Key: "username", Value: newUser.UserName}}).Decode(&models.User{}) != mongo.ErrNoDocuments {
		return errors.New("username already exists")
	}
	// The first user to register is an admin
	if curr, err := userColl.Find(context.TODO(), bson.D{{}}); err == nil && !curr.Next(context.TODO()) {
		newUser.UserRole = "admin"
	}else {
		newUser.UserRole = "regular"
	}

	if _, err := userColl.InsertOne(context.TODO(), newUser); err != nil {
		return err
	}
	return nil
}