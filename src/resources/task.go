package resources

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/J-Obog/tasket/src/managers"
	"github.com/J-Obog/tasket/src/models"
	"github.com/J-Obog/tasket/src/rest"
)

type TaskResource struct {
	taskManager managers.ITaskManager
	logManager  managers.ILogManager
}

func NewTaskResource(taskManager managers.ITaskManager, logManager managers.ILogManager) *TaskResource {
	return &TaskResource{
		taskManager: taskManager,
		logManager:  logManager,
	}
}

func (this *TaskResource) GetTask(req rest.RestRequest) (rest.RestResponse, int) {
	id := this.taskId(req.UrlParams)

	task, err := this.taskManager.GetTaskById(id)

	if err != nil {
		return rest.ServerErrorResponse(), http.StatusInternalServerError
	}

	if task == nil {
		res := rest.RestResponse{
			Error: rest.RestError{
				Type:   "resource not found",
				Detail: fmt.Sprintf("Task %s could not be found", id),
			},
		}

		return res, http.StatusNotFound
	}

	res := rest.RestResponse{
		Data: &task,
	}

	return res, http.StatusOK
}

func (this *TaskResource) GetTasks(req rest.RestRequest) (rest.RestResponse, int) {
	var taskOptions models.TaskOptions

	id := rest.UserId(req.Metadata)

	tasks, err := this.taskManager.GetTasksByFilter(id, taskOptions)

	if err != nil {
		return rest.ServerErrorResponse(), http.StatusInternalServerError
	}

	res := rest.RestResponse{
		Data: tasks,
	}

	return res, http.StatusOK
}

func (this *TaskResource) GetTaskLogs(req rest.RestRequest) (rest.RestResponse, int) {
	id := this.taskId(req.UrlParams)

	var options models.LogOptions

	err := json.Unmarshal(req.Body, &options)

	if err != nil {
		return rest.ServerErrorResponse(), http.StatusInternalServerError
	}

	//validate log filter

	logs, err := this.logManager.GetLogsByFilter(id, options)

	if err != nil {
		return rest.ServerErrorResponse(), http.StatusInternalServerError
	}

	res := rest.RestResponse{
		Data: logs,
	}

	return res, http.StatusOK
}

func (this *TaskResource) StopTask(req rest.RestRequest) (rest.RestResponse, int) {
	id := this.taskId(req.UrlParams)

	task, err := this.taskManager.GetTaskById(id)

	if err != nil {
		return rest.ServerErrorResponse(), http.StatusInternalServerError
	}

	if task != nil {
		res := rest.RestResponse{
			Error: rest.RestError{
				Type:   "not found",
				Detail: fmt.Sprintf("Task %s was not found", id),
			},
		}

		return res, http.StatusNotFound
	}

	err = this.taskManager.StopTask(id)

	if err != nil {
		return rest.ServerErrorResponse(), http.StatusInternalServerError
	}

	res := rest.RestResponse{
		Message: fmt.Sprintf("Task %s is set to be stopped", id),
	}

	return res, http.StatusOK
}

func (this *TaskResource) UpdateTask(req rest.RestRequest) (rest.RestResponse, int) {
	id := this.taskId(req.UrlParams)

	var updatedTask models.UpdatedTask

	err := json.Unmarshal(req.Body, &updatedTask)

	if err != nil {
		return rest.ServerErrorResponse(), http.StatusInternalServerError
	}

	task, err := this.taskManager.GetTaskById(id)
	if err != nil {
		return rest.ServerErrorResponse(), http.StatusInternalServerError
	}

	if task == nil {
		res := rest.RestResponse{
			Error: rest.RestError{
				Type:   "not found",
				Detail: fmt.Sprintf("Task %s was not found", id),
			},
		}

		return res, http.StatusNotFound
	}

	err = this.taskManager.UpdateTask(id, updatedTask)

	if err != nil {
		return rest.ServerErrorResponse(), http.StatusInternalServerError
	}

	res := rest.RestResponse{
		Message: fmt.Sprintf("Task %s updated successfully", id),
	}

	return res, http.StatusOK
}

func (this *TaskResource) CreateTask(req rest.RestRequest) (rest.RestResponse, int) {
	id := rest.UserId(req.Metadata)

	var newTask models.NewTask

	err := json.Unmarshal(req.Body, &newTask)

	if err != nil {
		return rest.ServerErrorResponse(), http.StatusInternalServerError
	}

	//validate task request

	err = this.taskManager.CreateTask(id, newTask)

	if err != nil {
		return rest.ServerErrorResponse(), http.StatusInternalServerError
	}

	res := rest.RestResponse{
		Message: "Task created successfully",
	}

	return res, http.StatusOK
}

func (this *TaskResource) taskId(urlParams map[string]interface{}) string {
	return urlParams["id"].(string)
}
