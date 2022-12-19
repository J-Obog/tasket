package resource

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/J-Obog/tasket/src/managers"
	"github.com/J-Obog/tasket/src/model"
	"github.com/J-Obog/tasket/src/rest"
)

type TaskResourceImpl struct {
	taskManager managers.TaskManager
	logManager  managers.LogManager
}

func NewTaskResource(taskManager managers.TaskManager, logManager managers.LogManager) *TaskResourceImpl {
	return &TaskResourceImpl{
		taskManager: taskManager,
		logManager:  logManager,
	}
}

func (this *TaskResourceImpl) Get(req rest.RestRequest) rest.RestResponse {
	id := this.taskId(req.UrlParams)

	task, err := this.taskManager.Get(id)

	if err != nil {
		return ServerErrorResponse()
	}

	if task == nil {
		resErr := MakeError("resource not found", fmt.Sprintf("Task %s could not be found", id))
		return MakeResponse(&resErr, http.StatusNotFound)
	}

	return MakeResponse(&task, http.StatusOK)
}

func (this *TaskResourceImpl) GetBy(req rest.RestRequest) rest.RestResponse {
	var taskOptions model.TaskOptions

	id := UserId(req.Metadata)

	tasks, err := this.taskManager.GetBy(id, taskOptions)

	if err != nil {
		return ServerErrorResponse()
	}

	return MakeResponse(tasks, http.StatusOK)
}

func (this *TaskResourceImpl) GetLogs(req rest.RestRequest) rest.RestResponse {
	id := this.taskId(req.UrlParams)

	var options model.LogOptions

	err := json.Unmarshal(req.Body, &options)

	if err != nil {
		return ServerErrorResponse()
	}

	//validate log filter

	logs, err := this.logManager.GetBy(id, options)

	if err != nil {
		return ServerErrorResponse()
	}

	return MakeResponse(logs, http.StatusOK)
}

func (this *TaskResourceImpl) Stop(req rest.RestRequest) rest.RestResponse {
	id := this.taskId(req.UrlParams)

	task, err := this.taskManager.Get(id)

	if err != nil {
		return ServerErrorResponse()
	}

	if task != nil {
		resErr := MakeError("not found", fmt.Sprintf("Task %s was not found", id))
		return MakeResponse(&resErr, http.StatusNotFound)
	}

	err = this.taskManager.Stop(id)

	if err != nil {
		return ServerErrorResponse()
	}

	msg := MakeMessage(fmt.Sprintf("Task %s is set to be stopped", id))
	return MakeResponse(&msg, http.StatusOK)
}

func (this *TaskResourceImpl) Update(req rest.RestRequest) rest.RestResponse {
	id := this.taskId(req.UrlParams)

	var updatedTask model.TaskUpdate

	err := json.Unmarshal(req.Body, &updatedTask)

	if err != nil {
		return ServerErrorResponse()
	}

	task, err := this.taskManager.Get(id)
	if err != nil {
		return ServerErrorResponse()
	}

	if task == nil {
		resErr := MakeError("not found", fmt.Sprintf("Task %s was not found", id))
		return MakeResponse(&resErr, http.StatusNotFound)
	}

	err = this.taskManager.Update(id, updatedTask)

	if err != nil {
		return ServerErrorResponse()
	}

	msg := MakeMessage(fmt.Sprintf("Task %s updated successfully", id))
	return MakeResponse(&msg, http.StatusOK)
}

func (this *TaskResourceImpl) Create(req rest.RestRequest) rest.RestResponse {
	//id := UserId(req.Metadata)
	return MakeResponse(nil, http.StatusOK)
}

func (this *TaskResourceImpl) taskId(urlParams map[string]interface{}) string {
	return urlParams["id"].(string)
}
