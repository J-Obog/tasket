package managers

import "github.com/J-Obog/tasket/src/model"

type TaskManager interface {
	Get(id string) (*model.Task, error)
	GetBy(userId string, options model.TaskOptions) ([]model.Task, error)
	Update(id string, taskUpdate model.TaskUpdate) error
	Stop(id string) error
	Create(userId string, name string, config model.TaskConfig) error
}
