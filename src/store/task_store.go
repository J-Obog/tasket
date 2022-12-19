package store

import "github.com/J-Obog/tasket/src/model"

type TaskStore interface {
	Get(id string) (*model.Task, error)
	GetBy(userId string, options model.TaskOptions) ([]model.Task, error)
	Update(id string, taskUpdate model.TaskUpdate) error
	Insert(task model.Task) error
}
