package data

import (
	"errors"
	"task_manager_api/models"
)

var tasks = map[string]models.Task{
	"1": {
		Id: "1",
		Title: "Complete Task 5",
		Description: "Complete Task Manager API and push work to github",
	},
}

func GetTask(id string) (models.Task, error) {
	task, ok := tasks[id]
	if !ok {
		return task, errors.New("task Not Found")
	}
	return task, nil
}

func GetTasks() []models.Task {
	ans := make([]models.Task, 0)
	for _, task := range tasks {
		ans = append(ans, task)
	}
	return ans
}
