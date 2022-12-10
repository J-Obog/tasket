package unit

import (
	"testing"

	"github.com/J-Obog/tasket/src/managers"
	"github.com/J-Obog/tasket/src/types"
	"github.com/J-Obog/tasket/tests"
	"github.com/J-Obog/tasket/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestTaskManager(t *testing.T) {
	mockStore := mocks.NewITaskStore(t)
	mockQueue := mocks.NewIQueue(t)
	mockLogger := mocks.NewILogger(t)
	mockUUIDProvider := mocks.NewIUUIDProvider(t)
	mockTimeProvider := mocks.NewITimeProvider(t)

	manager := managers.NewTaskManager(
		mockStore,
		mockQueue,
		mockLogger,
		mockUUIDProvider,
		mockTimeProvider,
	)

	t.Run("gets task by id", func(t *testing.T) {
		expectedTask := &types.Task{Id: tests.TEST_TASK_ID}

		mockStore.On("GetById", tests.TEST_TASK_ID).Return(expectedTask, nil)

		actualTask, actualErr := manager.GetTaskById(tests.TEST_TASK_ID)

		assert.NoError(t, actualErr)
		assert.Equal(t, actualTask, expectedTask)
	})

	t.Run("gets tasks by filter", func(t *testing.T) {
		beforeTime := new(int64)
		*beforeTime = tests.TEST_TIMESTAMP

		opts := types.TaskOptions{
			CreatedBefore: beforeTime,
		}

		task1 := types.Task{Name: "t1"}
		task2 := types.Task{Name: "t2"}

		expectedTasks := []types.Task{task1, task2}

		mockStore.On("GetByFilter", tests.TEST_USER_ID, opts).Return(expectedTasks, nil)

		actualTasks, actualErr := manager.GetTasksByFilter(tests.TEST_USER_ID, opts)

		assert.NoError(t, actualErr)
		assert.ElementsMatch(t, actualTasks, expectedTasks)
	})

	t.Run("updates a task", func(t *testing.T) {
		updatedTask := types.UpdatedTask{
			Name: "new name",
		}

		mockTimeProvider.On("Now").Return(tests.TEST_TIMESTAMP)

		mockStore.On("Update", tests.TEST_TASK_ID, updatedTask, tests.TEST_TIMESTAMP).Return(nil)

		actualErr := manager.UpdateTask(tests.TEST_TASK_ID, updatedTask)

		assert.NoError(t, actualErr)
	})

	t.Run("updates a task's status", func(t *testing.T) {
		newStatus := types.TaskStatus_COMPLETED

		mockTimeProvider.On("Now").Return(tests.TEST_TIMESTAMP)

		mockStore.On("UpdateStatus", tests.TEST_TASK_ID, newStatus, tests.TEST_TIMESTAMP).Return(nil)

		actualErr := manager.UpdateTaskStatus(tests.TEST_TASK_ID, newStatus)

		assert.NoError(t, actualErr)
	})

	t.Run("stops a task", func(t *testing.T) {
		eventMsg := types.EventMessage{
			Type: types.EventType_TASK_STOPPED,
			Data: &types.TaskStoppedMessage{
				TaskId: tests.TEST_TASK_ID,
			},
			Timestamp: tests.TEST_TIMESTAMP,
		}

		mockTimeProvider.On("Now").Return(tests.TEST_TIMESTAMP)
		mockQueue.On("Push", eventMsg).Return(nil)

		actualErr := manager.StopTask(tests.TEST_TASK_ID)

		assert.NoError(t, actualErr)
	})

	t.Run("creates a task", func(t *testing.T) {
		newTask := types.NewTask{
			Name: "foo bar",
			Config: types.TaskConfig{
				MaxTimeout: 60 * 5,
				MaxMemory:  1024 * 10,
			},
		}

		task := types.Task{
			Id:          tests.TEST_TASK_ID,
			UserId:      tests.TEST_USER_ID,
			Name:        newTask.Name,
			Status:      types.TaskStatus_PENDING,
			Config:      newTask.Config,
			StartedAt:   nil,
			CompletedAt: nil,
			CreatedAt:   tests.TEST_TIMESTAMP,
			UpdatedAt:   tests.TEST_TIMESTAMP,
		}

		eventMsg := types.EventMessage{
			Type: types.EventType_TASK_SCHEDULED,
			Data: &types.TaskScheduledMessage{
				TaskId:     tests.TEST_TASK_ID,
				TaskConfig: newTask.Config,
			},
			Timestamp: tests.TEST_TIMESTAMP,
		}

		mockTimeProvider.On("Now").Return(tests.TEST_TIMESTAMP)
		mockUUIDProvider.On("UUID").Return(tests.TEST_TASK_ID)
		mockStore.On("Insert", task).Return(nil)
		mockQueue.On("Push", eventMsg).Return(nil)

		actualErr := manager.CreateTask(tests.TEST_USER_ID, newTask)

		assert.NoError(t, actualErr)
	})

}
