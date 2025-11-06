package controller

import (
	"net/http"
	"task_manager_api/data"

	"github.com/gin-gonic/gin"
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
