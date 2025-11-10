package main

import (
	"fmt"
	"task_manager_api/data"
	"task_manager_api/router"
)

func main() {
	if err := data.ConnectDb("mongodb://localhost:27017"); err != nil {
		fmt.Println("Could not connect to database")
		fmt.Println(err)
	}
	app := router.GetRouter()
	app.Run(":8080")
}
