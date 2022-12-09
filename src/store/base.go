package store

import "github.com/J-Obog/tasket/src/models"

type IUserStore interface {
	GetById(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Insert(user models.User) error
	Delete(id string) error
}

type ITaskStore interface {
	GetById(id string) (*models.Task, error)
	GetByFilter(userId string, options models.TaskOptions) ([]models.Task, error)
	Update(id string, updatedTask models.UpdatedTask) error
	UpdateStatus(id string, status models.TaskStatus) error
	Insert(task models.Task) error
}

type ILogStore interface {
	GetByTask(taskId string) ([]models.Log, error)
	GetByFilter(taskId string, options models.LogOptions) ([]models.Log, error)
	Insert(log models.Log) error
}
