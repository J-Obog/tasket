package queue

import "github.com/J-Obog/tasket/src/models"

type Queue interface {
	Push(message models.Message)
	Pull() models.Message
}
