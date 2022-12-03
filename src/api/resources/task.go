package resources

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/J-Obog/tasket/src/api"
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

func (this *TaskResource) GetTask(req api.RestRequest) (api.RestResponse, int) {
	id := this.taskId(req.UrlParams)

	task, err := this.taskManager.GetTaskById(id)

	if err != nil {
		return api.ServerErrorResponse(), http.StatusInternalServerError
	}

	if task == nil {
		res := api.RestResponse{
			Error: api.RestError{
				Type:   "resource not found",
				Detail: fmt.Sprintf("Task %s could not be found", id),
			},
		}

		return res, http.StatusNotFound
	}

	res := api.RestResponse{
		Data: &task,
	}

	return res, http.StatusOK
}

func (this *TaskResource) GetTasks(req api.RestRequest) (api.RestResponse, int) {
	var taskOptions models.TaskOptions

	id := api.UserId(req.Metadata)

	tasks, err := this.taskManager.GetTasksByFilter(id, taskOptions)

	if err != nil {
		return api.ServerErrorResponse(), http.StatusInternalServerError
	}

	res := api.RestResponse{
		Data: tasks,
	}

	return res, http.StatusOK
}

func (this *TaskResource) GetTaskLogs(req api.RestRequest) (api.RestResponse, int) {
	id := this.taskId(req.UrlParams)

	var options models.LogOptions

	err := json.Unmarshal(req.Body, &options)

	if err != nil {
		return api.ServerErrorResponse(), http.StatusInternalServerError
	}

	//validate log filter

	logs, err := this.logManager.GetLogsByFilter(id, options)

	if err != nil {
		return api.ServerErrorResponse(), http.StatusInternalServerError
	}

	res := api.RestResponse{
		Data: logs,
	}

	return res, http.StatusOK
}

func (this *TaskResource) StopTask(req api.RestRequest) (api.RestResponse, int) {
	id := this.taskId(req.UrlParams)

	task, err := this.taskManager.GetTaskById(id)

	if err != nil {
		return api.ServerErrorResponse(), http.StatusInternalServerError
	}

	if task != nil {
		res := api.RestResponse{
			Error: api.RestError{
				Type:   "not found",
				Detail: fmt.Sprintf("Task %s was not found", id),
			},
		}

		return res, http.StatusNotFound
	}

	err = this.taskManager.StopTask(id)

	if err != nil {
		return api.ServerErrorResponse(), http.StatusInternalServerError
	}

	res := api.RestResponse{
		Message: fmt.Sprintf("Task %s is set to be stopped", id),
	}

	return res, http.StatusOK
}

func (this *TaskResource) UpdateTask(req api.RestRequest) (api.RestResponse, int) {
	id := this.taskId(req.UrlParams)

	var updatedTask models.UpdatedTask

	err := json.Unmarshal(req.Body, &updatedTask)

	if err != nil {
		return api.ServerErrorResponse(), http.StatusInternalServerError
	}

	task, err := this.taskManager.GetTaskById(id)
	if err != nil {
		return api.ServerErrorResponse(), http.StatusInternalServerError
	}

	if task == nil {
		res := api.RestResponse{
			Error: api.RestError{
				Type:   "not found",
				Detail: fmt.Sprintf("Task %s was not found", id),
			},
		}

		return res, http.StatusNotFound
	}

	err = this.taskManager.UpdateTask(id, updatedTask)

	if err != nil {
		return api.ServerErrorResponse(), http.StatusInternalServerError
	}

	res := api.RestResponse{
		Message: fmt.Sprintf("Task %s updated successfully", id),
	}

	return res, http.StatusOK
}

func (this *TaskResource) CreateTask(req api.RestRequest) (api.RestResponse, int) {
	id := api.UserId(req.Metadata)

	var newTask models.NewTask

	err := json.Unmarshal(req.Body, &newTask)

	if err != nil {
		return api.ServerErrorResponse(), http.StatusInternalServerError
	}

	//validate task request

	err = this.taskManager.CreateTask(id, newTask)

	if err != nil {
		return api.ServerErrorResponse(), http.StatusInternalServerError
	}

	res := api.RestResponse{
		Message: "Task created successfully",
	}

	return res, http.StatusOK
}

func (this *TaskResource) taskId(urlParams map[string]interface{}) string {
	return urlParams["id"].(string)
}
