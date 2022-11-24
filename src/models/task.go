package models

type TaskStatus uint

const (
	TaskStatus_UNKNOWN   TaskStatus = 0
	TaskStatus_COMPLETED            = 1
	TaskStatus_STOPPED              = 2
	TaskStatus_RUNNING              = 3
	TaskStatus_PENDING              = 4
)

type TaskConfig struct {
	MaxTimeout int64
	MaxMemory  int64
	Commands   []string
	Environ    map[string]string
	SourceMeta map[string]string
}

type Task struct {
	Id          string
	UserId      string
	Name        string
	Status      TaskStatus
	Config      TaskConfig
	StartedAt   *int64
	CompletedAt *int64
	CreatedAt   int64
	UpdatedAt   int64
}

type TaskRequest struct {
	Name   string
	Config TaskConfig
}

type TaskFilter struct {
	Name          *string
	Status        *TaskStatus
	CreatedAfter  *int64
	CreatedBefore *int64
}
