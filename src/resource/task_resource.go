package resource

type TaskResource interface {
	GetTask(req RestRequest) RestResponse
	GetTasks(req RestRequest) RestResponse
	GetTaskLogs(req RestRequest) RestResponse
	StopTask(req RestRequest) RestResponse
	UpdateTask(req RestRequest) RestResponse
	CreateTask(req RestRequest) RestResponse
}
