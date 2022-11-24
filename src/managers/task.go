package managers

import (
	"encoding/json"

	"github.com/J-Obog/tasket/src/models"
	"github.com/J-Obog/tasket/src/queue"
	"github.com/J-Obog/tasket/src/store"
	"github.com/J-Obog/tasket/src/utils"
)

type TaskManager struct {
	taskStore store.TaskStore
	taskQueue queue.Queue
}

func NewTaskManager(taskStore store.TaskStore, taskQueue queue.Queue) *TaskManager {
	return &TaskManager{
		taskStore: taskStore,
		taskQueue: taskQueue,
	}
}

func (this *TaskManager) GetTaskById(id string) (*models.Task, error) {
	return this.taskStore.GetById(id)
}

func (this *TaskManager) GetTasksByUser(userId string) ([]models.Task, error) {
	return this.taskStore.GetByUser(userId)
}

func (this *TaskManager) GetTasksByFilter(userId string, filter models.TaskFilter) ([]models.Task, error) {
	return this.taskStore.GetByFilter(userId, filter)
}

func (this *TaskManager) UpdateTaskName(id string, name string) error {
	return this.taskStore.UpdateName(id, name)
}

func (this *TaskManager) UpdateTaskStatus(id string, status models.TaskStatus) error {
	return this.taskStore.UpdateStatus(id, status)
}

func (this *TaskManager) StopTask(id string) error {
	return nil
}

func (this *TaskManager) CreateTask(userId string, taskReq models.TaskRequest) error {
	id := utils.GenerateUUID()
	now := utils.TimeNow()
	config := taskReq.Config

	task := models.Task{
		Id:          id,
		UserId:      userId,
		Name:        taskReq.Name,
		Status:      models.TaskStatus_PENDING,
		Config:      config,
		StartedAt:   nil,
		CompletedAt: nil,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	err := this.taskStore.Insert(task)
	if err != nil {
		return err
	}

	msg := models.ScheduledTask{
		TaskId:     id,
		TaskConfig: config,
	}

	data, err := json.Marshal(&msg)
	// log error

	err = this.taskQueue.Push(data)
	//log error

	return nil
}
