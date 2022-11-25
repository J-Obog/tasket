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

	body, err := json.Marshal(&task)

	if err != nil {
		return utils.MakeServerError()
	}

	return models.RestResponse{
		Body:   body,
		Status: http.StatusOK,
	}
}

func (this *TaskResource) GetTasks(req models.RestRequest) models.RestResponse {
	id := utils.CurrentUserId(req)

	tasks, err := this.taskManager.GetTasksByUser(id)

	if err != nil {
		return utils.MakeServerError()
	}

	body, err := json.Marshal(&tasks)

	if err != nil {
		return utils.MakeServerError()
	}

	return models.RestResponse{
		Body:   body,
		Status: http.StatusOK,
	}
}

func (this *TaskResource) GetTaskLogs(req models.RestRequest) models.RestResponse {
	id := utils.TaskIdParam()

	var logFilter models.LogFilter

	err := json.Unmarshal(req.Body, &logFilter)

	if err != nil {
		return utils.MakeServerError()
	}

	//validate log filter

	logs, err := this.logManager.GetLogsByFilter(id, logFilter)

	if err != nil {
		return utils.MakeServerError()
	}

	body, err := json.Marshal(&logs)

	if err != nil {
		return utils.MakeServerError()
	}

	return models.RestResponse{
		Body:   body,
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
		Body:   make([]byte, 0),
		Status: http.StatusOK,
	}
}

// todo: implement
func (this *TaskResource) UpdateTask(req models.RestRequest) models.RestResponse {
	return models.RestResponse{}
}

func (this *TaskResource) CreateTask(req models.RestRequest) models.RestResponse {
	id := utils.CurrentUserId(req)

	var taskReq models.TaskRequest

	err := json.Unmarshal(req.Body, &taskReq)

	if err != nil {
		return utils.MakeServerError()
	}

	//validate task request

	err = this.taskManager.CreateTask(id, taskReq)

	if err != nil {
		return utils.MakeServerError()
	}

	return models.RestResponse{
		Body:   make([]byte, 0),
		Status: http.StatusOK,
	}
}
