package managers

import (
	"github.com/J-Obog/tasket/src/models"
	"github.com/J-Obog/tasket/src/store"
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
	return nil
}

func (this *TaskManager) GetTasksByUser(userId string) []models.Task {
	return nil
}

func (this *TaskManager) GetTasksByFilter(userId string, filter models.TaskFilter) []models.Task {
	return nil
}

func (this *TaskManager) UpdateTaskName(id string, name string) {

}

func (this *TaskManager) UpdateTaskStatus(id string, status string) {

}

func (this *TaskManager) StopTask(id string) {

}

func (this *TaskManager) CreateTask(userId string, taskReq models.TaskRequest) {

}
