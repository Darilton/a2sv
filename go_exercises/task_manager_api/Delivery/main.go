package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"task_manager_api/Delivery/router"
	"task_manager_api/data"
)

var dbUri = "mongodb://localhost:27017"
var db = "a2sv"

func main() {
	clnt, err := mongo.Connect(options.Client().ApplyURI(dbUri))
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	// Set collections to use for each service class
	data.SetTaskCollection(clnt.Database(db).Collection("tasks"))
	data.SetUserCollection(clnt.Database(db).Collection("users"))

	app := router.GetRouter()
	app.Run(":8080")
}
