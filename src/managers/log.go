package managers

import (
	"github.com/J-Obog/tasket/src/models"
	"github.com/J-Obog/tasket/src/store"
	"github.com/J-Obog/tasket/src/utils"
)

type LogManager struct {
	logStore store.LogStore
}

func NewLogManager(logStore store.LogStore) *LogManager {
	return &LogManager{
		logStore: logStore,
	}
}

func (this *LogManager) GetLogsByTask(taskId string) ([]models.Log, error) {
	return this.logStore.GetByTask(taskId)
}

func (this *LogManager) GetLogsByFilter(taskId string, filter models.LogFilter) ([]models.Log, error) {
	return this.logStore.GetByFilter(taskId, filter)
}

func (this *LogManager) CreateLog(taskId string, logReq models.LogRequest) error {
	id := utils.GenerateUUID()
	now := utils.TimeNow()

	log := models.Log{
		Id:        id,
		TaskId:    logReq.TaskId,
		Source:    logReq.Source,
		Content:   logReq.Content,
		Timestamp: logReq.Timestamp,
		CreatedAt: now,
	}

	return this.logStore.Insert(log)
}
