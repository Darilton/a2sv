package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task_manager_api/data"
	"task_manager_api/models"
)

func GetTasks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, data.GetTasks())
}

func GetTask(ctx *gin.Context) {
	task, err := data.GetTask(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func AddTask(ctx *gin.Context) {
	var newTask models.Task

	if err := ctx.ShouldBindBodyWithJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := data.AddTask(newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task Added Successfully"})
}

func PutTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var newTask models.Task
	if err := ctx.ShouldBindBodyWithJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := data.EditTask(id, newTask); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task Updated Successfully!"})
}

func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := data.DeleteTask(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task Deleteded Successfully!"})
}
