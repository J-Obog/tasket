package models

type Task struct {
	Id          string
	UserId      string
	Name        string
	Status      string
	MaxTimeout  int64
	MaxMemory   int64
	Commands    []string
	Environ     map[string]string
	SourceMeta  map[string]string
	StartedAt   *int64
	CompletedAt *int64
	CreatedAt   int64
	UpdatedAt   int64
}

type TaskRequest struct {
	Name       string
	MaxTimeout int64
	MaxMemory  int64
	Commands   []string
	Environ    map[string]string
	SourceMeta map[string]string
}

type TaskFilter struct {
	Name          *string
	Status        *string
	CreatedAfter  *int64
	CreatedBefore *int64
}
