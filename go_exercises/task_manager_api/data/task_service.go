package data

import (
	"context"
	"errors"
	"task_manager_api/models"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var client *mongo.Client
var coll *mongo.Collection

func ConnectDb(dbUri string) error {
	clnt, err := mongo.Connect(options.Client().ApplyURI(dbUri))
	if err != nil {
		return err
	}
	client = clnt
	coll = client.Database("a2sv").Collection("tasks")
	return nil
}

func AddTask(newTask models.Task) error {
	if newTask.Id == "" || newTask.Title == "" {
		return errors.New("invalid Request")
	}
	if _, err := coll.InsertOne(context.TODO(), newTask); err != nil {
		return err
	}
	return nil
}

func GetTask(id string) (models.Task, error) {
	var task models.Task
	err := coll.FindOne(context.TODO(), bson.D{{Key: "id", Value: id}}).Decode(&task)
	if err != nil {
		return task, errors.New("task Not found")
	}
	return task, nil
}

func GetTasks() []models.Task {
	ans := make([]models.Task, 0)
	cur, err := coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return ans
	}
	for cur.Next(context.TODO()) {
		var task models.Task
		cur.Decode(&task)
		ans = append(ans, task)
	}
	return ans
}

func EditTask(id string, newTask models.Task) error {
	if id != newTask.Id {
		return errors.New("Invalid Request")
	}

	updateResult, _ := coll.ReplaceOne(context.TODO(), bson.D{{Key: "id", Value: id}}, newTask)
	if updateResult.MatchedCount == 0 {
		return errors.New("task with given id not found")
	}

	return nil
}

func DeleteTask(id string) error {
	deleteResult, _ := coll.DeleteOne(context.TODO(), bson.D{{Key: "id", Value: id}})
	if deleteResult.DeletedCount == 0 {
		return errors.New("task with given id not found")
	}
	return nil
}
