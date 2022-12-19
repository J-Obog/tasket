package managers

import (
	"github.com/J-Obog/tasket/src/model"
	"github.com/J-Obog/tasket/src/store"
	"github.com/J-Obog/tasket/src/types"
)

type LogManagerImpl struct {
	logStore     store.LogStore
	timeProvider types.ITimeProvider
	uuidProvider types.IUUIDProvider
}

func NewLogManager(
	logStore store.LogStore, timeProvider types.ITimeProvider, uuidProvider types.IUUIDProvider) *LogManagerImpl {
	return &LogManagerImpl{
		logStore:     logStore,
		timeProvider: timeProvider,
		uuidProvider: uuidProvider,
	}
}

func (this *LogManagerImpl) GetBy(taskId string, opts model.LogOptions) ([]model.Log, error) {
	return this.logStore.GetBy(taskId, opts)
}

func (this *LogManagerImpl) Create(taskId string, source model.LogSource, content []byte) error {
	log := model.Log{
		Id:        this.uuidProvider.UUID(),
		TaskId:    taskId,
		Source:    source,
		Content:   content,
		CreatedAt: this.timeProvider.Now(),
	}

	return this.logStore.Insert(log)
}
