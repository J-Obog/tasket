package managers

import (
	"github.com/J-Obog/tasket/src/models"
	"github.com/J-Obog/tasket/src/store"
)

type LogManager struct {
	logStore store.LogStore
}

func NewLogManager(logStore store.LogStore) *LogManager {
	return &LogManager{
		logStore: logStore,
	}
}

func (this *LogManager) GetLogsByTask(taskId string) []models.Log {
	return nil
}

func (this *LogManager) GetLogsByFilter(taskId string, filter models.LogFilter) []models.Log {
	return nil
}

func (this *LogManager) CreateLog(taskId string, logReq models.LogRequest) {

}
