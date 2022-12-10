package managers

import (
	"github.com/J-Obog/tasket/src/types"
)

type LogManager struct {
	logStore     types.ILogStore
	timeProvider types.ITimeProvider
	uuidProvider types.IUUIDProvider
}

func NewLogManager(logStore types.ILogStore, timeProvider types.ITimeProvider, uuidProvider types.IUUIDProvider) *LogManager {
	return &LogManager{
		logStore:     logStore,
		timeProvider: timeProvider,
		uuidProvider: uuidProvider,
	}
}

func (this *LogManager) GetLogsByTask(taskId string) ([]types.Log, error) {
	return this.logStore.GetByTask(taskId)
}

func (this *LogManager) GetLogsByFilter(taskId string, filter types.LogOptions) ([]types.Log, error) {
	return this.logStore.GetByFilter(taskId, filter)
}

func (this *LogManager) CreateLog(taskId string, newLog types.NewLog) error {
	id := this.uuidProvider.UUID()
	now := this.timeProvider.Now()

	log := types.Log{
		Id:        id,
		TaskId:    newLog.TaskId,
		Source:    newLog.Source,
		Content:   newLog.Content,
		CreatedAt: now,
	}

	return this.logStore.Insert(log)
}
