package router

import (
	"task_manager_api/Delivery/controller"
	middleware "task_manager_api/Infrastructure"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", middleware.AuthMiddleware(), controller.GetTasks)
	router.GET("/tasks/:id", middleware.AuthMiddleware(), controller.GetTask)
	router.PUT("/tasks/:id", middleware.AuthMiddleware(), middleware.CheckAdminMiddleware(), controller.PutTask)
	router.DELETE("/tasks/:id", middleware.AuthMiddleware(), middleware.CheckAdminMiddleware(), controller.DeleteTask)
	router.POST("/tasks", middleware.AuthMiddleware(), middleware.CheckAdminMiddleware(), controller.AddTask)

	router.POST("/register", controller.RegisterUser)
	router.POST("/login", controller.LoginUser)
	return router
}
