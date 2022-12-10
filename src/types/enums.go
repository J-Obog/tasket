package types

type EventType string

const (
	EventType_UNKNOWN        EventType = ""
	EventType_TASK_SCHEDULED EventType = "task.scheduled"
	EventType_TASK_STOPPED   EventType = "task.stopped"
)

type LogSource uint

const (
	LogSource_UNKNOWN LogSource = 0
	LogSource_TASK    LogSource = 1
	LogSource_SERVER  LogSource = 2
)

type TaskStatus uint

const (
	TaskStatus_UNKNOWN   TaskStatus = 0
	TaskStatus_COMPLETED TaskStatus = 1
	TaskStatus_STOPPED   TaskStatus = 2
	TaskStatus_RUNNING   TaskStatus = 3
	TaskStatus_PENDING   TaskStatus = 4
)
