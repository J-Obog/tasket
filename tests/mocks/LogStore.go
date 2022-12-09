// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	models "github.com/J-Obog/tasket/src/models"
	mock "github.com/stretchr/testify/mock"
)

// LogStore is an autogenerated mock type for the LogStore type
type LogStore struct {
	mock.Mock
}

// GetByFilter provides a mock function with given fields: taskId, options
func (_m *LogStore) GetByFilter(taskId string, options models.LogOptions) ([]models.Log, error) {
	ret := _m.Called(taskId, options)

	var r0 []models.Log
	if rf, ok := ret.Get(0).(func(string, models.LogOptions) []models.Log); ok {
		r0 = rf(taskId, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Log)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, models.LogOptions) error); ok {
		r1 = rf(taskId, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByTask provides a mock function with given fields: taskId
func (_m *LogStore) GetByTask(taskId string) ([]models.Log, error) {
	ret := _m.Called(taskId)

	var r0 []models.Log
	if rf, ok := ret.Get(0).(func(string) []models.Log); ok {
		r0 = rf(taskId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Log)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(taskId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: log
func (_m *LogStore) Insert(log models.Log) error {
	ret := _m.Called(log)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Log) error); ok {
		r0 = rf(log)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewLogStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewLogStore creates a new instance of LogStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLogStore(t mockConstructorTestingTNewLogStore) *LogStore {
	mock := &LogStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
