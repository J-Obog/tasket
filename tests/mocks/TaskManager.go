// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	model "github.com/J-Obog/tasket/src/model"
	mock "github.com/stretchr/testify/mock"
)

// TaskManager is an autogenerated mock type for the TaskManager type
type TaskManager struct {
	mock.Mock
}

// Create provides a mock function with given fields: userId, name, config
func (_m *TaskManager) Create(userId string, name string, config model.TaskConfig) error {
	ret := _m.Called(userId, name, config)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, model.TaskConfig) error); ok {
		r0 = rf(userId, name, config)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *TaskManager) Get(id string) (*model.Task, error) {
	ret := _m.Called(id)

	var r0 *model.Task
	if rf, ok := ret.Get(0).(func(string) *model.Task); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Task)
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

// GetBy provides a mock function with given fields: userId, options
func (_m *TaskManager) GetBy(userId string, options model.TaskOptions) ([]model.Task, error) {
	ret := _m.Called(userId, options)

	var r0 []model.Task
	if rf, ok := ret.Get(0).(func(string, model.TaskOptions) []model.Task); ok {
		r0 = rf(userId, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, model.TaskOptions) error); ok {
		r1 = rf(userId, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Stop provides a mock function with given fields: id
func (_m *TaskManager) Stop(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: id, taskUpdate
func (_m *TaskManager) Update(id string, taskUpdate model.TaskUpdate) error {
	ret := _m.Called(id, taskUpdate)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, model.TaskUpdate) error); ok {
		r0 = rf(id, taskUpdate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTaskManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewTaskManager creates a new instance of TaskManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTaskManager(t mockConstructorTestingTNewTaskManager) *TaskManager {
	mock := &TaskManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}