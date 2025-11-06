package router

import (
	"github.com/gin-gonic/gin"
	"task_manager_api/controller"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", controller.GetTasks)
	router.GET("/tasks/:id", controller.GetTask)
	return router
}
