package resources

import (
	"github.com/J-Obog/tasket/src/managers"
	"github.com/J-Obog/tasket/src/models"
)

type TaskResource struct {
	taskManager managers.TaskManager
	logManager  managers.LogManager
}

func NewTaskResource(taskManager managers.TaskManager, logManager managers.LogManager) *TaskResource {
	return &TaskResource{
		taskManager: taskManager,
		logManager:  logManager,
	}
}

func (this *TaskResource) GetTask(req models.RestRequest) models.RestResponse {
	return models.RestResponse{}
}

func (this *TaskResource) GetTasks(req models.RestRequest) models.RestResponse {
	return models.RestResponse{}
}

func (this *TaskResource) GetTaskLogs(req models.RestRequest) models.RestResponse {
	return models.RestResponse{}
}

func (this *TaskResource) StopTask(req models.RestRequest) models.RestResponse {
	return models.RestResponse{}
}

func (this *TaskResource) UpdateTask(req models.RestRequest) models.RestResponse {
	return models.RestResponse{}
}

func (this *TaskResource) CreateTask(req models.RestRequest) models.RestResponse {
	return models.RestResponse{}
}
