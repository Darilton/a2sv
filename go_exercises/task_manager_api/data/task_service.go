package data

import (
	"context"
	"errors"
	"task_manager_api/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var taskColl *mongo.Collection

func SetTaskCollection(collection *mongo.Collection) {
	taskColl = collection
}

func AddTask(newTask models.Task) error {
	if newTask.Id == "" || newTask.Title == "" {
		return errors.New("invalid Request")
	}
	if _, err := taskColl.InsertOne(context.TODO(), newTask); err != nil {
		return err
	}
	return nil
}

func GetTask(id string) (models.Task, error) {
	var task models.Task
	err := taskColl.FindOne(context.TODO(), bson.D{{Key: "id", Value: id}}).Decode(&task)
	if err != nil {
		return task, errors.New("task Not found")
	}
	return task, nil
}

func GetTasks() []models.Task {
	ans := make([]models.Task, 0)
	cur, err := taskColl.Find(context.TODO(), bson.D{{}})
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

	updateResult, _ := taskColl.ReplaceOne(context.TODO(), bson.D{{Key: "id", Value: id}}, newTask)
	if updateResult.MatchedCount == 0 {
		return errors.New("task with given id not found")
	}

	return nil
}

func DeleteTask(id string) error {
	deleteResult, _ := taskColl.DeleteOne(context.TODO(), bson.D{{Key: "id", Value: id}})
	if deleteResult.DeletedCount == 0 {
		return errors.New("task with given id not found")
	}
	return nil
}
