package repositories_integration_test

import (
	"context"
	domain "task_manager_api/Domain"
	"task_manager_api/Repositories"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type UserRepositorySuite struct {
	suite.Suite
	client *mongo.Client
	repo   Repositories.UserRepository
	coll   *mongo.Collection
}

func (suite *UserRepositorySuite) SetupSuite() {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	suite.Require().NoError(err)
	suite.client = client
	suite.coll = client.Database("task_manager_test").Collection("users")
	suite.repo = Repositories.NewUserRepository(suite.coll)
}

func (suite *UserRepositorySuite) TearDownSuite() {
	suite.client.Database("task_manager_test").Drop(context.TODO())
	suite.client.Disconnect(context.TODO())
}

func (suite *UserRepositorySuite) SetupTest() {
	suite.coll.DeleteMany(context.TODO(), bson.D{{}})
}

func (suite *UserRepositorySuite) TestAddUser() {
	user := domain.User{UserName: "testuser", Password: "hashedpassword"}
	err := suite.repo.AddUser(user)
	suite.NoError(err)

	var retrieved domain.User
	err = suite.coll.FindOne(context.TODO(), bson.D{{Key: "username", Value: "testuser"}}).Decode(&retrieved)
	suite.NoError(err)
	suite.Equal("testuser", retrieved.UserName)
	suite.Equal("admin", retrieved.UserRole) // First user should be admin
}

func (suite *UserRepositorySuite) TestGetUser() {
	user := domain.User{UserName: "testuser", Password: "hashedpassword"}
	suite.repo.AddUser(user)

	retrieved, err := suite.repo.GetUser("testuser")
	suite.NoError(err)
	suite.Equal("testuser", retrieved.UserName)
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}
