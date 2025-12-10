package main

import (
	"fmt"
	"task_manager_api/Delivery/controller"
	"task_manager_api/Delivery/router"
	"task_manager_api/Repositories"
	usecases "task_manager_api/UseCases"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var dbUri = "mongodb://localhost:27017"
var db = "a2sv"

func main() {
	clnt, err := mongo.Connect(options.Client().ApplyURI(dbUri))
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	// Initialize Repositories
	taskRepo := Repositories.NewTaskRepository(clnt.Database(db).Collection("tasks"))
	userRepo := Repositories.NewUserRepository(clnt.Database(db).Collection("users"))

	// Initialize UseCases
	taskUseCase := usecases.NewTaskUseCase(taskRepo)
	userUseCase := usecases.NewUserUseCase(userRepo)

	// Initialize Controller
	taskController := controller.NewTaskController(taskUseCase, userUseCase)

	app := router.GetRouter(taskController)
	app.Run(":8080")
}
