package queue

import (
	"github.com/J-Obog/tasket/src/model"
)

func MakeTaskStoppedMsg(id string, timestamp int64) EventMessage {
	msg := TaskStoppedMessage{
		TaskId: id,
	}

	return EventMessage{
		Type:      EventType_TASK_STOPPED,
		Data:      &msg,
		Timestamp: timestamp,
	}
}

func MakeTaskScheduledMsg(id string, config model.TaskConfig, timestamp int64) EventMessage {
	msg := TaskScheduledMessage{
		TaskId:     id,
		TaskConfig: config,
	}

	return EventMessage{
		Type:      EventType_TASK_SCHEDULED,
		Data:      &msg,
		Timestamp: timestamp,
	}
}
