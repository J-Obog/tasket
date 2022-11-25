package store

import "github.com/J-Obog/tasket/src/models"

type UserStore interface {
	GetById(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Insert(user models.User) error
	Delete(id string) error
}

type TaskStore interface {
	GetById(id string) (*models.Task, error)
	GetByUser(userId string) ([]models.Task, error)
	GetByFilter(userId string, filter models.TaskFilter) ([]models.Task, error)
	UpdateName(id string, name string) error
	UpdateStatus(id string, status models.TaskStatus) error
	Insert(task models.Task) error
}

type LogStore interface {
	GetByTask(taskId string) ([]models.Log, error)
	GetByFilter(taskId string, filter models.LogFilter) ([]models.Log, error)
	Insert(log models.Log) error
}
