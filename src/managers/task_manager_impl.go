package managers

import (
	"fmt"

	"github.com/J-Obog/tasket/src/logger"
	"github.com/J-Obog/tasket/src/model"
	"github.com/J-Obog/tasket/src/queue"
	"github.com/J-Obog/tasket/src/store"
	"github.com/J-Obog/tasket/src/types"
)

type TaskManagerImpl struct {
	taskStore    store.TaskStore
	taskQueue    queue.Queue
	logger       logger.Logger
	uuidProvider types.IUUIDProvider
	timeProvider types.ITimeProvider
}

func NewTaskManager(
	taskStore store.TaskStore,
	taskQueue queue.Queue,
	logger logger.Logger,
	uuidProvider types.IUUIDProvider,
	timeProvider types.ITimeProvider,
) *TaskManagerImpl {
	return &TaskManagerImpl{
		taskStore:    taskStore,
		taskQueue:    taskQueue,
		logger:       logger,
		uuidProvider: uuidProvider,
		timeProvider: timeProvider,
	}
}

func (this *TaskManagerImpl) Get(id string) (*model.Task, error) {
	return this.taskStore.Get(id)
}

func (this *TaskManagerImpl) GetBy(userId string, opts model.TaskOptions) ([]model.Task, error) {
	return this.taskStore.GetBy(userId, opts)
}

func (this *TaskManagerImpl) Update(id string, taskUpdate model.TaskUpdate) error {
	return this.taskStore.Update(id, taskUpdate)
}

func (this *TaskManagerImpl) Stop(id string) error {
	now := this.timeProvider.Now()
	msg := queue.MakeTaskStoppedMsg(id, now)
	err := this.taskQueue.Push(msg)

	if err != nil {
		this.logger.Error(fmt.Sprintf("Failed to push event %v", msg))
		return err
	}

	return nil
}

func (this *TaskManagerImpl) Create(userId string, name string, config model.TaskConfig) error {
	id := this.uuidProvider.UUID()
	now := this.timeProvider.Now()

	task := model.Task{
		Id:          id,
		UserId:      userId,
		Name:        name,
		Status:      model.TaskStatus_PENDING,
		Config:      config,
		StartedAt:   nil,
		CompletedAt: nil,
		CreatedAt:   now,
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
