package resource

import "github.com/J-Obog/tasket/src/model"

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetTasksRequest struct {
	Name          *string           `json:"name"`
	Status        *model.TaskStatus `json:"status"`
	CreatedAfter  *int64            `json:"createdAfter"`
	CreatedBefore *int64            `json:"createdBefore"`
}

type GetTaskLogsRequest struct {
	Source        *model.LogSource `json:"source"`
	CreatedAfter  *int64           `json:"createdAfter"`
	CreatedBefore *int64           `json:"createdBefore"`
}

type UpdateTaskRequest struct {
	Name string `json:"name"`
}

type CreateTaskRequest struct {
	Name   string           `json:"name"`
	Config model.TaskConfig `json:"config"`
	Token  string           `json:"token"`
}
