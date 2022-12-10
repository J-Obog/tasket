package managers

import (
	"fmt"

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
	return this.taskStore.Update(id, updatedTask)
}

func (this *TaskManager) UpdateTaskStatus(id string, status types.TaskStatus) error {
	return this.taskStore.UpdateStatus(id, status)
}

func (this *TaskManager) StopTask(id string) error {
	event := types.TaskStoppedEvent{
		Type:   types.EventType_TASK_STOPPED,
		TaskId: id,
	}

	err := this.taskQueue.Push(&event)
	if err != nil {
		this.logger.Error(fmt.Sprintf("Failed to push event %v", event))
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

	event := types.TaskScheduledEvent{
		Type:       types.EventType_TASK_SCHEDULED,
		TaskId:     id,
		TaskConfig: config,
	}

	err = this.taskQueue.Push(&event)

	if err != nil {
		this.logger.Error(fmt.Sprintf("Failed to push event %v", event))
	}

	return nil
}
