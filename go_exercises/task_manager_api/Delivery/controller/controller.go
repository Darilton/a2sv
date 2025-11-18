package controller

import (
	"net/http"
	domain "task_manager_api/Domain"
	"task_manager_api/Repositories"
	usecases "task_manager_api/UseCases"

	"github.com/gin-gonic/gin"
)

func GetTasks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Repositories.GetTasks())
}

func GetTask(ctx *gin.Context) {
	task, err := usecases.GetTask(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func LoginUser(ctx *gin.Context) {
	var loginData domain.User

	if err := ctx.ShouldBindBodyWithJSON(&loginData); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := usecases.LoginUser(loginData)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func RegisterUser(ctx *gin.Context) {
	var newUser domain.User

	if err := ctx.ShouldBindBodyWithJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := usecases.RegisterUser(newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User Registered Successfully"})
}

func AddTask(ctx *gin.Context) {
	var newTask domain.Task

	if err := ctx.ShouldBindBodyWithJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := usecases.AddTask(newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task Added Successfully"})
}

func PutTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var newTask domain.Task
	if err := ctx.ShouldBindBodyWithJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := usecases.EditTask(id, newTask); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task Updated Successfully!"})
}

func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := usecases.DeleteTask(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task Deleteded Successfully!"})
}
