// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	model "github.com/J-Obog/tasket/src/model"
	mock "github.com/stretchr/testify/mock"
)

// TaskStore is an autogenerated mock type for the TaskStore type
type TaskStore struct {
	mock.Mock
}

// Get provides a mock function with given fields: id
func (_m *TaskStore) Get(id string) (*model.Task, error) {
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
func (_m *TaskStore) GetBy(userId string, options model.TaskOptions) ([]model.Task, error) {
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

// Insert provides a mock function with given fields: task
func (_m *TaskStore) Insert(task model.Task) error {
	ret := _m.Called(task)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Task) error); ok {
		r0 = rf(task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: id, taskUpdate
func (_m *TaskStore) Update(id string, taskUpdate model.TaskUpdate) error {
	ret := _m.Called(id, taskUpdate)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, model.TaskUpdate) error); ok {
		r0 = rf(id, taskUpdate)
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