package managers

import "github.com/J-Obog/tasket/src/models"

type ILogManager interface {
	GetLogsByTask(taskId string) ([]models.Log, error)
	GetLogsByFilter(taskId string, filter models.LogOptions) ([]models.Log, error)
	CreateLog(taskId string, logReq models.NewLog)
}

type ITaskManager interface {
	GetTaskById(id string) (*models.Task, error)
	GetTasksByFilter(userId string, options models.TaskOptions) ([]models.Task, error)
	UpdateTask(id string, updatedTask models.UpdatedTask) error
	UpdateTaskStatus(id string, status models.TaskStatus) error
	StopTask(id string) error
	CreateTask(userId string, newTask models.NewTask) error
}

type IUserManager interface {
	GetUserById(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(userReq models.NewUser) error
	DeleteUser(id string) error
}
