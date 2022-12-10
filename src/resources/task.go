package resources

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/J-Obog/tasket/src/types"
)

type TaskResource struct {
	taskManager types.ITaskManager
	logManager  types.ILogManager
}

func NewTaskResource(taskManager types.ITaskManager, logManager types.ILogManager) *TaskResource {
	return &TaskResource{
		taskManager: taskManager,
		logManager:  logManager,
	}
}

func (this *TaskResource) GetTask(req types.RestRequest) types.RestResponse {
	id := this.taskId(req.UrlParams)

	task, err := this.taskManager.GetTaskById(id)

	if err != nil {
		return ServerErrorResponse()
	}

	if task == nil {
		resErr := MakeError("resource not found", fmt.Sprintf("Task %s could not be found", id))
		return MakeResponse(&resErr, http.StatusNotFound)
	}

	return MakeResponse(&task, http.StatusOK)
}

func (this *TaskResource) GetTasks(req types.RestRequest) types.RestResponse {
	var taskOptions types.TaskOptions

	id := UserId(req.Metadata)

	tasks, err := this.taskManager.GetTasksByFilter(id, taskOptions)

	if err != nil {
		return ServerErrorResponse()
	}

	return MakeResponse(tasks, http.StatusOK)
}

func (this *TaskResource) GetTaskLogs(req types.RestRequest) types.RestResponse {
	id := this.taskId(req.UrlParams)

	var options types.LogOptions

	err := json.Unmarshal(req.Body, &options)

	if err != nil {
		return ServerErrorResponse()
	}

	//validate log filter

	logs, err := this.logManager.GetLogsByFilter(id, options)

	if err != nil {
		return ServerErrorResponse()
	}

	return MakeResponse(logs, http.StatusOK)
}

func (this *TaskResource) StopTask(req types.RestRequest) types.RestResponse {
	id := this.taskId(req.UrlParams)

	task, err := this.taskManager.GetTaskById(id)

	if err != nil {
		return ServerErrorResponse()
	}

	if task != nil {
		resErr := MakeError("not found", fmt.Sprintf("Task %s was not found", id))
		return MakeResponse(&resErr, http.StatusNotFound)
	}

	err = this.taskManager.StopTask(id)

	if err != nil {
		return ServerErrorResponse()
	}

	msg := MakeMessage(fmt.Sprintf("Task %s is set to be stopped", id))
	return MakeResponse(&msg, http.StatusOK)
}

func (this *TaskResource) UpdateTask(req types.RestRequest) types.RestResponse {
	id := this.taskId(req.UrlParams)

	var updatedTask types.UpdatedTask

	err := json.Unmarshal(req.Body, &updatedTask)

	if err != nil {
		return ServerErrorResponse()
	}

	task, err := this.taskManager.GetTaskById(id)
	if err != nil {
		return ServerErrorResponse()
	}

	if task == nil {
		resErr := MakeError("not found", fmt.Sprintf("Task %s was not found", id))
		return MakeResponse(&resErr, http.StatusNotFound)
	}

	err = this.taskManager.UpdateTask(id, updatedTask)

	if err != nil {
		return ServerErrorResponse()
	}

	msg := MakeMessage(fmt.Sprintf("Task %s updated successfully", id))
	return MakeResponse(&msg, http.StatusOK)
}

func (this *TaskResource) CreateTask(req types.RestRequest) types.RestResponse {
	id := UserId(req.Metadata)

	var newTask types.NewTask

	err := json.Unmarshal(req.Body, &newTask)

	if err != nil {
		return ServerErrorResponse()
	}

	//validate task request

	err = this.taskManager.CreateTask(id, newTask)

	if err != nil {
		return ServerErrorResponse()
	}

	msg := MakeMessage("Task created successfully")
	return MakeResponse(&msg, http.StatusOK)
}

func (this *TaskResource) taskId(urlParams map[string]interface{}) string {
	return urlParams["id"].(string)
}
