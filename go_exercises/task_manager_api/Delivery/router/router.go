package router

import (
	"task_manager_api/Delivery/controller"
	middleware "task_manager_api/Infrastructure"

	"github.com/gin-gonic/gin"
)

func GetRouter(tc *controller.TaskController) *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", middleware.AuthMiddleware(), tc.GetTasks)
	router.GET("/tasks/:id", middleware.AuthMiddleware(), tc.GetTask)
	router.PUT("/tasks/:id", middleware.AuthMiddleware(), middleware.CheckAdminMiddleware(), tc.PutTask)
	router.DELETE("/tasks/:id", middleware.AuthMiddleware(), middleware.CheckAdminMiddleware(), tc.DeleteTask)
	router.POST("/tasks", middleware.AuthMiddleware(), middleware.CheckAdminMiddleware(), tc.AddTask)

	router.POST("/register", tc.RegisterUser)
	router.POST("/login", tc.LoginUser)
	return router
}
