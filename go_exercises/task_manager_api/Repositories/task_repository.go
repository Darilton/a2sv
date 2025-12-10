package Repositories

import (
	"context"
	"errors"
	domain "task_manager_api/Domain"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type TaskRepository interface {
	AddTask(newTask domain.Task) error
	GetTask(id string) (domain.Task, error)
	GetTasks() []domain.Task
	EditTask(id string, newTask domain.Task) error
	DeleteTask(id string) error
}

type TaskRepositoryMongo struct {
	collection *mongo.Collection
}

func NewTaskRepository(collection *mongo.Collection) TaskRepository {
	return &TaskRepositoryMongo{
		collection: collection,
	}
}

func (tr *TaskRepositoryMongo) AddTask(newTask domain.Task) error {
	if newTask.Id == "" || newTask.Title == "" {
		return errors.New("invalid Request")
	}
	if _, err := tr.collection.InsertOne(context.TODO(), newTask); err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepositoryMongo) GetTask(id string) (domain.Task, error) {
	var task domain.Task
	err := tr.collection.FindOne(context.TODO(), bson.D{{Key: "id", Value: id}}).Decode(&task)
	if err != nil {
		return task, errors.New("task Not found")
	}
	return task, nil
}

func (tr *TaskRepositoryMongo) GetTasks() []domain.Task {
	ans := make([]domain.Task, 0)
	cur, err := tr.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return ans
	}
	for cur.Next(context.TODO()) {
		var task domain.Task
		cur.Decode(&task)
		ans = append(ans, task)
	}
	return ans
}

func (tr *TaskRepositoryMongo) EditTask(id string, newTask domain.Task) error {
	if id != newTask.Id {
		return errors.New("invalid Request")
	}

	updateResult, err := tr.collection.ReplaceOne(context.TODO(), bson.D{{Key: "id", Value: id}}, newTask)
	if err != nil {
		return err
	}
	if updateResult.MatchedCount == 0 {
		return errors.New("task with given id not found")
	}

	return nil
}

func (tr *TaskRepositoryMongo) DeleteTask(id string) error {
	deleteResult, err := tr.collection.DeleteOne(context.TODO(), bson.D{{Key: "id", Value: id}})
	if err != nil {
		return err
	}
	if deleteResult.DeletedCount == 0 {
		return errors.New("task with given id not found")
	}
	return nil
}
