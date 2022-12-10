package models

type TaskStatus uint

const (
	TaskStatus_UNKNOWN   TaskStatus = 0
	TaskStatus_COMPLETED TaskStatus = 1
	TaskStatus_STOPPED   TaskStatus = 2
	TaskStatus_RUNNING   TaskStatus = 3
	TaskStatus_PENDING   TaskStatus = 4
)

type TaskConfig struct {
	MaxTimeout int64             `json:"maxTimeout"`
	MaxMemory  int64             `json:"maxMemory"`
	Commands   []string          `json:"commands"`
	Environ    map[string]string `json:"environ"`
	SourceMeta map[string]string `json:"sourceMeta"`
}

type Task struct {
	Id          string     `json:"id"`
	UserId      string     `json:"userId"`
	Name        string     `json:"name"`
	Status      TaskStatus `json:"status"`
	Config      TaskConfig `json:"config"`
	StartedAt   *int64     `json:"startedAt"`
	CompletedAt *int64     `json:"completedAt"`
	CreatedAt   int64      `json:"createdAt"`
	UpdatedAt   int64      `json:"updatedAt"`
}

type NewTask struct {
	Name   string     `json:"name"`
	Config TaskConfig `json:"config"`
}

type UpdatedTask struct {
	Name string `json:"name"`
}

type TaskOptions struct {
	Name          *string     `json:"name"`
	Status        *TaskStatus `json:"status"`
	CreatedAfter  *int64      `json:"createdAfter"`
	CreatedBefore *int64      `json:"createdBefore"`
}
