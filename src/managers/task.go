package managers

import (
	"fmt"

	"github.com/J-Obog/tasket/src/queue"
	"github.com/J-Obog/tasket/src/types"
)

type TaskManager struct {
	taskStore    types.ITaskStore
	taskQueue    types.IQueue
	logger       types.ILogger
	uuidProvider types.IUUIDProvider
	timeProvider types.ITimeProvider
}

func NewTaskManager(
	taskStore types.ITaskStore,
	taskQueue types.IQueue,
	logger types.ILogger,
	uuidProvider types.IUUIDProvider,
	timeProvider types.ITimeProvider,
) *TaskManager {
	return &TaskManager{
		taskStore:    taskStore,
		taskQueue:    taskQueue,
		logger:       logger,
		uuidProvider: uuidProvider,
		timeProvider: timeProvider,
	}
}

func (this *TaskManager) GetTaskById(id string) (*types.Task, error) {
	return this.taskStore.GetById(id)
}

func (this *TaskManager) GetTasksByFilter(userId string, options types.TaskOptions) ([]types.Task, error) {
	return this.taskStore.GetByFilter(userId, options)
}

func (this *TaskManager) UpdateTask(id string, updatedTask types.UpdatedTask) error {
	now := this.timeProvider.Now()
	return this.taskStore.Update(id, updatedTask, now)
}

func (this *TaskManager) UpdateTaskStatus(id string, status types.TaskStatus) error {
	now := this.timeProvider.Now()
	return this.taskStore.UpdateStatus(id, status, now)
}

func (this *TaskManager) StopTask(id string) error {
	now := this.timeProvider.Now()

	msg := queue.MakeTaskStoppedMsg(id, now)
	err := this.taskQueue.Push(msg)

	if err != nil {
		this.logger.Error(fmt.Sprintf("Failed to push event %v", msg))
		return err
	}

	return nil
}

func (this *TaskManager) CreateTask(userId string, newTask types.NewTask) error {
	id := this.uuidProvider.UUID()
	now := this.timeProvider.Now()
	config := newTask.Config

	task := types.Task{
		Id:          id,
		UserId:      userId,
		Name:        newTask.Name,
		Status:      types.TaskStatus_PENDING,
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

	msg := queue.MakeTaskScheduledMsg(id, config, now)
	err = this.taskQueue.Push(msg)

	if err != nil {
		this.logger.Error(fmt.Sprintf("Failed to push event %v", msg))
	}

	return nil
}
