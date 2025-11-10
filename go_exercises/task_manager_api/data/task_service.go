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

var tasks = map[string]models.Task{
	"1": {
		Id:          "1",
		Title:       "Task 1",
		Description: "First task",
		DueDate:     time.Now(),
		Status:      "Pending",
	},
	"2": {
		Id:          "2",
		Title:       "Task 2",
		Description: "Second task",
		DueDate:     time.Now().AddDate(0, 0, 1),
		Status:      "In Progress",
	},
	"3": {
		Id:          "3",
		Title:       "Task 3",
		Description: "Third task",
		DueDate:     time.Now().AddDate(0, 0, 2),
		Status:      "Completed",
	},
}

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
	for _, task := range tasks {
		ans = append(ans, task)
	}
	return ans
}

func EditTask(id string, newTask models.Task) error {
	task, ok := tasks[id]
	if !ok {
		return errors.New("task Not Found")
	}
	newTask.Id = id
	tasks[task.Id] = task
	return nil
}

func DeleteTask(id string) error {
	task, ok := tasks[id]
	if !ok {
		return errors.New("task Not Found")
	}
	delete(tasks, task.Id)
	return nil
}
