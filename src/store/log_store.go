package store

import "github.com/J-Obog/tasket/src/model"

type LogStore interface {
	GetBy(taskId string, options model.LogOptions) ([]model.Log, error)
	Insert(log model.Log) error
}
