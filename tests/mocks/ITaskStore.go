// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	types "github.com/J-Obog/tasket/src/types"
	mock "github.com/stretchr/testify/mock"
)

// ITaskStore is an autogenerated mock type for the ITaskStore type
type ITaskStore struct {
	mock.Mock
}

// GetByFilter provides a mock function with given fields: userId, options
func (_m *ITaskStore) GetByFilter(userId string, options types.TaskOptions) ([]types.Task, error) {
	ret := _m.Called(userId, options)

	var r0 []types.Task
	if rf, ok := ret.Get(0).(func(string, types.TaskOptions) []types.Task); ok {
		r0 = rf(userId, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, types.TaskOptions) error); ok {
		r1 = rf(userId, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *ITaskStore) GetById(id string) (*types.Task, error) {
	ret := _m.Called(id)

	var r0 *types.Task
	if rf, ok := ret.Get(0).(func(string) *types.Task); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Task)
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
func (_m *ITaskStore) Insert(task types.Task) error {
	ret := _m.Called(task)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Task) error); ok {
		r0 = rf(task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: id, updatedTask
func (_m *ITaskStore) Update(id string, updatedTask types.UpdatedTask) error {
	ret := _m.Called(id, updatedTask)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, types.UpdatedTask) error); ok {
		r0 = rf(id, updatedTask)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateStatus provides a mock function with given fields: id, status
func (_m *ITaskStore) UpdateStatus(id string, status types.TaskStatus) error {
	ret := _m.Called(id, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, types.TaskStatus) error); ok {
		r0 = rf(id, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewITaskStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewITaskStore creates a new instance of ITaskStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewITaskStore(t mockConstructorTestingTNewITaskStore) *ITaskStore {
	mock := &ITaskStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
