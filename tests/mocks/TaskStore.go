// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	models "github.com/J-Obog/tasket/src/models"
	mock "github.com/stretchr/testify/mock"
)

// TaskStore is an autogenerated mock type for the TaskStore type
type TaskStore struct {
	mock.Mock
}

// GetByFilter provides a mock function with given fields: userId, options
func (_m *TaskStore) GetByFilter(userId string, options models.TaskOptions) ([]models.Task, error) {
	ret := _m.Called(userId, options)

	var r0 []models.Task
	if rf, ok := ret.Get(0).(func(string, models.TaskOptions) []models.Task); ok {
		r0 = rf(userId, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, models.TaskOptions) error); ok {
		r1 = rf(userId, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *TaskStore) GetById(id string) (*models.Task, error) {
	ret := _m.Called(id)

	var r0 *models.Task
	if rf, ok := ret.Get(0).(func(string) *models.Task); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: task
func (_m *TaskStore) Insert(task models.Task) error {
	ret := _m.Called(task)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Task) error); ok {
		r0 = rf(task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: id, updatedTask
func (_m *TaskStore) Update(id string, updatedTask models.UpdatedTask) error {
	ret := _m.Called(id, updatedTask)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, models.UpdatedTask) error); ok {
		r0 = rf(id, updatedTask)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateStatus provides a mock function with given fields: id, status
func (_m *TaskStore) UpdateStatus(id string, status models.TaskStatus) error {
	ret := _m.Called(id, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, models.TaskStatus) error); ok {
		r0 = rf(id, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTaskStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewTaskStore creates a new instance of TaskStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTaskStore(t mockConstructorTestingTNewTaskStore) *TaskStore {
	mock := &TaskStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
