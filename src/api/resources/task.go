package resources

import (
	"encoding/json"
	"net/http"

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

func (this *TaskResource) GetTask(req models.RestRequest) models.RestResponse {
	id := utils.TaskIdParam()

	task, err := this.taskManager.GetTaskById(id)

	if err != nil {
		return utils.MakeServerError()
	}

	return models.RestResponse{
		Object: map[string]interface{}{
			"data": &task,
		},
		Status: http.StatusOK,
	}
}

func (this *TaskResource) GetTasks(req models.RestRequest) models.RestResponse {
	id := utils.CurrentUserId(req)

	tasks, err := this.taskManager.GetTasksByUser(id)

	if err != nil {
		return utils.MakeServerError()
	}

	return models.RestResponse{
		Object: map[string]interface{}{
			"data": tasks,
		},
		Status: http.StatusOK,
	}
}

func (this *TaskResource) GetTaskLogs(req models.RestRequest) models.RestResponse {
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

	return models.RestResponse{
		Object: map[string]interface{}{
			"data": logs,
		},
		Status: http.StatusOK,
	}
}

func (this *TaskResource) StopTask(req models.RestRequest) models.RestResponse {
	id := utils.TaskIdParam()

	err := this.taskManager.StopTask(id)

	if err != nil {
		return utils.MakeServerError()
	}

	return models.RestResponse{
		Object: map[string]interface{}{
			"message": "Task stopped successfully",
		},
		Status: http.StatusOK,
	}
}

// todo: implement
func (this *TaskResource) UpdateTask(req models.RestRequest) models.RestResponse {
	return models.RestResponse{}
}

func (this *TaskResource) CreateTask(req models.RestRequest) models.RestResponse {
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

	return models.RestResponse{
		Object: map[string]interface{}{
			"message": "Task creaetd successfully",
		},
		Status: http.StatusOK,
	}
}
