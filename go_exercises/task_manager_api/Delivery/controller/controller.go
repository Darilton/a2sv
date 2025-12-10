package controller

import (
	"net/http"
	domain "task_manager_api/Domain"
	usecases "task_manager_api/UseCases"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskUseCase *usecases.TaskUseCase
	userUseCase *usecases.UserUseCase
}

func NewTaskController(taskUseCase *usecases.TaskUseCase, userUseCase *usecases.UserUseCase) *TaskController {
	return &TaskController{
		taskUseCase: taskUseCase,
		userUseCase: userUseCase,
	}
}

func (tc *TaskController) GetTasks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, tc.taskUseCase.GetTasks())
}

func (tc *TaskController) GetTask(ctx *gin.Context) {
	task, err := tc.taskUseCase.GetTask(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (tc *TaskController) LoginUser(ctx *gin.Context) {
	var loginData domain.User

	if err := ctx.ShouldBindBodyWithJSON(&loginData); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := tc.userUseCase.LoginUser(loginData)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (tc *TaskController) RegisterUser(ctx *gin.Context) {
	var newUser domain.User

	if err := ctx.ShouldBindBodyWithJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := tc.userUseCase.RegisterUser(newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User Registered Successfully"})
}

func (tc *TaskController) AddTask(ctx *gin.Context) {
	var newTask domain.Task

	if err := ctx.ShouldBindBodyWithJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := tc.taskUseCase.AddTask(newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task Added Successfully"})
}

func (tc *TaskController) PutTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var newTask domain.Task
	if err := ctx.ShouldBindBodyWithJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := tc.taskUseCase.EditTask(id, newTask); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task Updated Successfully!"})
}

func (tc *TaskController) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := tc.taskUseCase.DeleteTask(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task Deleteded Successfully!"})
}
