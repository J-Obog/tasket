package managers

import "github.com/J-Obog/tasket/src/model"

type LogManager interface {
	GetBy(taskId string, options model.LogOptions) ([]model.Log, error)
	Create(taskId string, source model.LogSource, content []byte) error
}
