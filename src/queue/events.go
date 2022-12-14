package queue

import "github.com/J-Obog/tasket/src/model"

type EventType string

const (
	EventType_UNKNOWN        EventType = ""
	EventType_TASK_SCHEDULED EventType = "task.scheduled"
	EventType_TASK_STOPPED   EventType = "task.stopped"
)

type EventMessage struct {
	Type      EventType   `json:"type"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

type TaskScheduledMessage struct {
	TaskId     string           `json:"taskId"`
	TaskConfig model.TaskConfig `json:"taskConfig"`
}

type TaskStoppedMessage struct {
	TaskId string `json:"taskId"`
}
