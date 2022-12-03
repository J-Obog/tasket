package queue

import "github.com/J-Obog/tasket/src/models"

type TaskScheduledEvent struct {
	TaskId     string
	TaskConfig models.TaskConfig
}
