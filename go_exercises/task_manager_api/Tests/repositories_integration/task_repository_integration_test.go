package repositories_integration_test

import (
	"context"
	domain "task_manager_api/Domain"
	"task_manager_api/Repositories"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type TaskRepositorySuite struct {
	suite.Suite
	client *mongo.Client
	repo   Repositories.TaskRepository
	coll   *mongo.Collection
}

func (suite *TaskRepositorySuite) SetupSuite() {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	suite.Require().NoError(err)
	suite.client = client
	suite.coll = client.Database("task_manager_test").Collection("tasks")
	suite.repo = Repositories.NewTaskRepository(suite.coll)
}

func (suite *TaskRepositorySuite) TearDownSuite() {
	suite.client.Database("task_manager_test").Drop(context.TODO())
	suite.client.Disconnect(context.TODO())
}

func (suite *TaskRepositorySuite) SetupTest() {
	suite.coll.DeleteMany(context.TODO(), bson.D{{}})
}

func (suite *TaskRepositorySuite) TestAddTask() {
	task := domain.Task{Id: "1", Title: "Test Task", Description: "Desc", DueDate: time.Now(), Status: "Pending"}
	err := suite.repo.AddTask(task)
	suite.NoError(err)

	var retrieved domain.Task
	err = suite.coll.FindOne(context.TODO(), bson.D{{Key: "id", Value: "1"}}).Decode(&retrieved)
	suite.NoError(err)
	suite.Equal("Test Task", retrieved.Title)
}

func (suite *TaskRepositorySuite) TestGetTask() {
	task := domain.Task{Id: "1", Title: "Test Task"}
	suite.repo.AddTask(task)

	retrieved, err := suite.repo.GetTask("1")
	suite.NoError(err)
	suite.Equal("Test Task", retrieved.Title)
}

func (suite *TaskRepositorySuite) TestDeleteTask() {
	task := domain.Task{Id: "1", Title: "Test Task"}
	suite.repo.AddTask(task)

	err := suite.repo.DeleteTask("1")
	suite.NoError(err)

	_, err = suite.repo.GetTask("1")
	suite.Error(err)
}

func TestTaskRepositorySuite(t *testing.T) {
	suite.Run(t, new(TaskRepositorySuite))
}
