package data

import (
	"errors"
	"task_manager_api/models"
)

var tasks = map[string]models.Task{
	"1": {
		Id:          "1",
		Title:       "Complete Task 5",
		Description: "Complete Task Manager API and push work to github",
	},
}

func AddTask(newTask models.Task) error {
	if newTask.Id == "" || newTask.Title == "" {
		return errors.New("invalid Request")
	}
	// makes shure to add only if task does not exists
	_, ok := tasks[newTask.Id]
	if ok {
		return errors.New("task Already Exists")
	}

	tasks[newTask.Id] = newTask
	return nil
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

func EditTask(id string, newTask models.Task) error {
	task, ok := tasks[id]
	if !ok {
		return errors.New("task Not Found")
	}
	newTask.Id = id
	tasks[task.Id] = task
	return nil
}

func DeleteTask(id string) error {
	task, ok := tasks[id]
	if !ok {
		return errors.New("task Not Found")
	}
	delete(tasks, task.Id)
	return nil
}
