package store

import "github.com/J-Obog/tasket/src/models"

type UserStore interface {
	GetById(id string) *models.User
	Insert(user models.User)
	Delete(id string)
}

type TaskStore interface {
	GetById(id string) *models.Task
	GetByUser(userId string) []models.Task
	GetByFilter(userId string, filter models.TaskFilter) []models.Task
	UpdateName(id string, name string)
	UpdateStatus(id string, status models.TaskStatus)
	Insert(task models.Task)
}

type LogStore interface {
	GetByTask(taskId string) []models.Log
	GetByFilter(taskId string, filter models.LogFilter) []models.Log
	Insert(log models.Log)
}
