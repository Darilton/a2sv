package router

import (
	"github.com/gin-gonic/gin"
	"task_manager_api/controller"
	"task_manager_api/middleware"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", middleware.AuthMiddleware(), controller.GetTasks)
	router.GET("/tasks/:id", middleware.AuthMiddleware(), controller.GetTask)
	router.PUT("/tasks/:id", middleware.AuthMiddleware(), controller.PutTask)
	router.DELETE("/tasks/:id", middleware.AuthMiddleware(), controller.DeleteTask)
	router.POST("/tasks", middleware.AuthMiddleware(), controller.AddTask)

	router.POST("/register", controller.RegisterUser)
	router.POST("/login", controller.LoginUser)
	return router
}
