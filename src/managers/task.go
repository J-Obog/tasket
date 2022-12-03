package managers

import (
	"fmt"

	"github.com/J-Obog/tasket/src/logger"
	"github.com/J-Obog/tasket/src/models"
	"github.com/J-Obog/tasket/src/queue"
	"github.com/J-Obog/tasket/src/store"
	"github.com/J-Obog/tasket/src/utils"
)

type TaskManager struct {
	taskStore store.TaskStore
	taskQueue queue.Queue
	logger    logger.Logger
}

func NewTaskManager(taskStore store.TaskStore, taskQueue queue.Queue, logger logger.Logger) *TaskManager {
	return &TaskManager{
		taskStore: taskStore,
		taskQueue: taskQueue,
		logger:    logger,
	}
}

func (this *TaskManager) GetTaskById(id string) (*models.Task, error) {
	return this.taskStore.GetById(id)
}

func (this *TaskManager) GetTasksByFilter(userId string, options models.TaskOptions) ([]models.Task, error) {
	return this.taskStore.GetByFilter(userId, options)
}

func (this *TaskManager) UpdateTask(id string, updatedTask models.UpdatedTask) error {
	return this.taskStore.Update(id, updatedTask)
}

func (this *TaskManager) UpdateTaskStatus(id string, status models.TaskStatus) error {
	return this.taskStore.UpdateStatus(id, status)
}

func (this *TaskManager) StopTask(id string) error {
	event := queue.TaskStoppedEvent{
		Type:   queue.EventType_TASK_STOPPED,
		TaskId: id,
	}

	err := this.taskQueue.Push(&event)
	if err != nil {
		return err
	}

	return nil
}

func (this *TaskManager) CreateTask(userId string, newTask models.NewTask) error {
	id := utils.GenerateUUID()
	now := utils.TimeNow()
	config := newTask.Config

	task := models.Task{
		Id:          id,
		UserId:      userId,
		Name:        newTask.Name,
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

	event := queue.TaskScheduledEvent{
		Type:       queue.EventType_TASK_SCHEDULED,
		TaskId:     id,
		TaskConfig: config,
	}

	err = this.taskQueue.Push(&event)

	if err != nil {
		msg := fmt.Sprintf("Failed to push %s event for task %s", queue.EventType_TASK_SCHEDULED, id)
		this.logger.Error(msg)
	}

	return nil
}
