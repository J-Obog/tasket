package unit

import (
	"testing"

	"github.com/J-Obog/tasket/src/managers"
	"github.com/J-Obog/tasket/src/models"
	"github.com/J-Obog/tasket/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTaskManager(t *testing.T) {
	mockStore := mocks.NewITaskStore(t)
	mockQueue := mocks.NewIQueue(t)
	mockLogger := mocks.NewILogger(t)
	manager := managers.NewTaskManager(mockStore, mockQueue, mockLogger)

	t.Run("gets task by id", func(t *testing.T) {
		taskId := "task_12345"

		expectedTask := &models.Task{Id: taskId}

		mockStore.On("GetById", taskId).Return(expectedTask, nil)

		actualTask, actualErr := manager.GetTaskById(taskId)

		assert.NoError(t, actualErr)
		assert.Equal(t, actualTask, expectedTask)
	})

	t.Run("gets tasks by filter", func(t *testing.T) {
		now := int64(1234567)
		userId := "user_12345"

		opts := models.TaskOptions{
			CreatedBefore: &now,
		}

		task1 := models.Task{Name: "t1"}
		task2 := models.Task{Name: "t2"}

		expectedTasks := []models.Task{task1, task2}

		mockStore.On("GetByFilter", userId, opts).Return(expectedTasks, nil)

		actualTasks, actualErr := manager.GetTasksByFilter(userId, opts)

		assert.NoError(t, actualErr)
		assert.ElementsMatch(t, actualTasks, expectedTasks)
	})

	t.Run("updates a task", func(t *testing.T) {
		taskId := "task_42434"

		updatedTask := models.UpdatedTask{
			Name: "new name",
		}

		mockStore.On("Update", taskId, updatedTask).Return(nil)

		actualErr := manager.UpdateTask(taskId, updatedTask)

		assert.NoError(t, actualErr)
	})

	t.Run("updates a task's status", func(t *testing.T) {
		taskId := "task_1234"

		newStatus := models.TaskStatus_COMPLETED

		mockStore.On("UpdateStatus", taskId, newStatus).Return(nil)

		actualErr := manager.UpdateTaskStatus(taskId, newStatus)

		assert.NoError(t, actualErr)
	})

	t.Run("stops a task", func(t *testing.T) {
		taskId := "task_12345"

		mockQueue.On("Push", mock.Anything).Return(nil)

		manager.StopTask(taskId)

		mockQueue.AssertNumberOfCalls(t, "Push", 1)
	})

}
