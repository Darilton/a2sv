package main

import "task_manager_api/router"

func main() {
	app := router.GetRouter()
	app.Run(":8080")
}
