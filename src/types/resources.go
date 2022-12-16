package types

type IUserResource interface {
	GetUser(req RestRequest) RestResponse
	DeleteUser(req RestRequest) RestResponse
	CreateUser(req RestRequest) RestResponse
}

type ITaskResource interface {
	GetTask(req RestRequest) RestResponse
	GetTasks(req RestRequest) RestResponse
	GetTaskLogs(req RestRequest) RestResponse
	StopTask(req RestRequest) RestResponse
	UpdateTask(req RestRequest) RestResponse
	CreateTask(req RestRequest) RestResponse
}
