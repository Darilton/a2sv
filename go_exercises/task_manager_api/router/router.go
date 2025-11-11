package router

import (
	"github.com/gin-gonic/gin"
	"task_manager_api/controller"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", controller.GetTasks)
	router.GET("/tasks/:id", controller.GetTask)
	router.PUT("/tasks/:id", controller.PutTask)
	router.DELETE("/tasks/:id", controller.DeleteTask)
	router.POST("/tasks", controller.AddTask)

	router.POST("/register", controller.RegisterUser)
	return router
}
