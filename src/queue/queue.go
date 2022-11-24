package queue

import "github.com/J-Obog/tasket/src/models"

type Queue interface {
	Push(message []byte) error
	Pull() (models.Message, error)
}
