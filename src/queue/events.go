package queue

import "github.com/J-Obog/tasket/src/models"

type EventType string

const (
	EventType_UNKNOWN        EventType = ""
	EventType_TASK_SCHEDULED EventType = "task.scheduled"
	EventType_TASK_STOPPED   EventType = "task.stopped"
)

type TaskScheduledEvent struct {
	Type       EventType         `json:"type"`
	TaskId     string            `json:"taskId"`
	TaskConfig models.TaskConfig `json:"taskConfig"`
}

type TaskStoppedEvent struct {
	Type   EventType `json:"type"`
	TaskId string    `json:"taskId"`
}
