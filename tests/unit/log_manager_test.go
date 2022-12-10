package unit

import (
	"testing"

	"github.com/J-Obog/tasket/src/managers"
	"github.com/J-Obog/tasket/src/types"
	"github.com/J-Obog/tasket/tests"
	"github.com/J-Obog/tasket/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestLogManager(t *testing.T) {
	mockStore := mocks.NewILogStore(t)
	mockUUIDProvider := mocks.NewIUUIDProvider(t)
	mockTimeProvider := mocks.NewITimeProvider(t)

	manager := managers.NewLogManager(mockStore, mockTimeProvider, mockUUIDProvider)

	t.Run("gets logs for task", func(t *testing.T) {
		log1 := types.Log{Id: "log1"}
		log2 := types.Log{Id: "log2"}
		log3 := types.Log{Id: "log3"}

		expectedLogs := []types.Log{log1, log2, log3}

		mockStore.On("GetByTask", tests.TEST_TASK_ID).Return(expectedLogs, nil)

		actualLogs, actualErr := manager.GetLogsByTask(tests.TEST_TASK_ID)

		assert.ElementsMatch(t, actualLogs, expectedLogs)
		assert.NoError(t, actualErr)
	})

	t.Run("gets logs by filter", func(t *testing.T) {
		first := new(int64)
		*first = 2

		opts := types.LogOptions{
			First: first,
		}

		log1 := types.Log{Id: "log1"}
		log2 := types.Log{Id: "log2"}

		expectedLogs := []types.Log{log1, log2}

		mockStore.On("GetByFilter", tests.TEST_TASK_ID, opts).Return(expectedLogs, nil)

		actualLogs, actualErr := manager.GetLogsByFilter(tests.TEST_TASK_ID, opts)

		assert.ElementsMatch(t, actualLogs, expectedLogs)
		assert.NoError(t, actualErr)
	})

	t.Run("creates a log", func(t *testing.T) {
		newLog := types.NewLog{
			TaskId:  tests.TEST_TASK_ID,
			Source:  types.LogSource_TASK,
			Content: []byte(`hello world`),
		}

		log := types.Log{
			Id:        tests.TEST_LOG_ID,
			TaskId:    newLog.TaskId,
			Source:    newLog.Source,
			Content:   newLog.Content,
			CreatedAt: tests.TEST_TIMESTAMP,
		}

		mockUUIDProvider.On("UUID").Return(tests.TEST_LOG_ID)
		mockTimeProvider.On("Now").Return(tests.TEST_TIMESTAMP)
		mockStore.On("Insert", log).Return(nil)

		actualErr := manager.CreateLog(tests.TEST_TASK_ID, newLog)

		assert.NoError(t, actualErr)
	})

}
