package managers

import (
	"github.com/J-Obog/tasket/src/models"
	"github.com/J-Obog/tasket/src/store"
	"github.com/J-Obog/tasket/src/utils"
)

type TaskManager struct {
	taskStore store.TaskStore
}

func NewTaskManager(taskStore store.TaskStore) *TaskManager {
	return &TaskManager{
		taskStore: taskStore,
	}
}

func (this *TaskManager) GetTaskById(id string) *models.Task {
	return this.taskStore.GetById(id)
}

func (this *TaskManager) GetTasksByUser(userId string) []models.Task {
	return this.taskStore.GetByUser(userId)
}

func (this *TaskManager) GetTasksByFilter(userId string, filter models.TaskFilter) []models.Task {
	return this.taskStore.GetByFilter(userId, filter)
}

func (this *TaskManager) UpdateTaskName(id string, name string) {
	this.taskStore.UpdateName(id, name)
}

func (this *TaskManager) UpdateTaskStatus(id string, status models.TaskStatus) {
	this.taskStore.UpdateStatus(id, status)
}

func (this *TaskManager) StopTask(id string) {

}

func (this *TaskManager) CreateTask(userId string, taskReq models.TaskRequest) {
	id := utils.GenerateUUID()
	now := utils.TimeNow()

	task := models.Task{
		Id:          id,
		UserId:      userId,
		Name:        taskReq.Name,
		Status:      models.TaskStatus_PENDING,
		MaxTimeout:  taskReq.MaxTimeout,
		MaxMemory:   taskReq.MaxMemory,
		Commands:    taskReq.Commands,
		Environ:     taskReq.Environ,
		SourceMeta:  taskReq.SourceMeta,
		StartedAt:   nil,
		CompletedAt: nil,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	this.taskStore.Insert(task)
}
