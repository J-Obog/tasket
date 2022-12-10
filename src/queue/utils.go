package queue

import "github.com/J-Obog/tasket/src/types"

func MakeTaskStoppedMsg(id string, timestamp int64) types.EventMessage {
	msg := types.TaskStoppedMessage{
		TaskId: id,
	}

	return types.EventMessage{
		Type:      types.EventType_TASK_STOPPED,
		Data:      &msg,
		Timestamp: timestamp,
	}
}

func MakeTaskScheduledMsg(id string, config types.TaskConfig, timestamp int64) types.EventMessage {
	msg := types.TaskScheduledMessage{
		TaskId:     id,
		TaskConfig: config,
	}

	return types.EventMessage{
		Type:      types.EventType_TASK_SCHEDULED,
		Data:      &msg,
		Timestamp: timestamp,
	}
}
