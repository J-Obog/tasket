package resources

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/J-Obog/tasket/src/api"
	"github.com/J-Obog/tasket/src/managers"
	"github.com/J-Obog/tasket/src/models"
	"github.com/J-Obog/tasket/src/utils"
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

func (this *TaskResource) GetTask(req api.RestRequest) api.RestResponse {
	id := utils.TaskIdParam()

	task, err := this.taskManager.GetTaskById(id)

	if err != nil {
		return utils.MakeServerError()
	}

	return api.RestResponse{
		Object: &task,
		Status: http.StatusOK,
	}
}

func (this *TaskResource) GetTasks(req api.RestRequest) api.RestResponse {
	id := utils.CurrentUserId(req)

	tasks, err := this.taskManager.GetTasksByUser(id)

	if err != nil {
		return utils.MakeServerError()
	}

	return api.RestResponse{
		Object: tasks,
		Status: http.StatusOK,
	}
}

func (this *TaskResource) GetTaskLogs(req api.RestRequest) api.RestResponse {
	id := utils.TaskIdParam()

	var options models.LogOptions

	err := json.Unmarshal(req.Body, &options)

	if err != nil {
		return utils.MakeServerError()
	}

	//validate log filter

	logs, err := this.logManager.GetLogsByFilter(id, options)

	if err != nil {
		return utils.MakeServerError()
	}

	return api.RestResponse{
		Object: logs,
		Status: http.StatusOK,
	}
}

func (this *TaskResource) StopTask(req api.RestRequest) api.RestResponse {
	id := utils.TaskIdParam()

	task, err := this.taskManager.GetTaskById(id)

	if err != nil {
		return utils.MakeServerError()
	}

	if task != nil {
		return api.RestResponse{
			Object: map[string]interface{}{
				"message": fmt.Sprintf("Task %s was not found", id),
			},
			Status: http.StatusNotFound,
		}
	}

	err = this.taskManager.StopTask(id)

	if err != nil {
		return utils.MakeServerError()
	}

	return api.RestResponse{
		Object: map[string]interface{}{
			"message": fmt.Sprintf("Task %s is set to be stopped", id),
		},
		Status: http.StatusOK,
	}
}

func (this *TaskResource) UpdateTask(req api.RestRequest) api.RestResponse {
	id := utils.TaskIdParam()

	var updatedTask models.UpdatedTask

	err := json.Unmarshal(req.Body, &updatedTask)

	if err != nil {
		return utils.MakeServerError()
	}

	task, err := this.taskManager.GetTaskById(id)
	if err != nil {
		return utils.MakeServerError()
	}

	if task == nil {
		return api.RestResponse{
			Object: map[string]interface{}{
				"message": fmt.Sprintf("Task with id %s was not found", id),
			},
			Status: http.StatusNotFound,
		}
	}

	err = this.taskManager.UpdateTask(id, updatedTask)

	if err != nil {
		return utils.MakeServerError()
	}

	return api.RestResponse{
		Object: map[string]interface{}{
			"message": "Task updated successfully",
		},
		Status: http.StatusOK,
	}
}

func (this *TaskResource) CreateTask(req api.RestRequest) api.RestResponse {
	id := utils.CurrentUserId(req)

	var newTask models.NewTask

	err := json.Unmarshal(req.Body, &newTask)

	if err != nil {
		return utils.MakeServerError()
	}

	//validate task request

	err = this.taskManager.CreateTask(id, newTask)

	if err != nil {
		return utils.MakeServerError()
	}

	return api.RestResponse{
		Object: map[string]interface{}{
			"message": "Task created successfully",
		},
		Status: http.StatusOK,
	}
}
