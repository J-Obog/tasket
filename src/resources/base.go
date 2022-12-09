package resources

import (
	"github.com/J-Obog/tasket/src/rest"
)

type IUserResource interface {
	GetUser(req rest.RestRequest) (rest.RestResponse, int)
	DeleteUser(req rest.RestRequest) (rest.RestResponse, int)
	CreateUser(req rest.RestRequest) (rest.RestResponse, int)
}

type ITaskResource interface {
	GetTask(req rest.RestRequest) (rest.RestResponse, int)
	GetTasks(req rest.RestRequest) (rest.RestResponse, int)
	GetTaskLogs(req rest.RestRequest) (rest.RestResponse, int)
	StopTask(req rest.RestRequest) (rest.RestResponse, int)
	UpdateTask(req rest.RestRequest) (rest.RestResponse, int)
	CreateTask(req rest.RestRequest) (rest.RestResponse, int)
}
