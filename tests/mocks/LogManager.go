// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	model "github.com/J-Obog/tasket/src/model"
	mock "github.com/stretchr/testify/mock"
)

// LogManager is an autogenerated mock type for the LogManager type
type LogManager struct {
	mock.Mock
}

// Create provides a mock function with given fields: taskId, source, content
func (_m *LogManager) Create(taskId string, source model.LogSource, content []byte) error {
	ret := _m.Called(taskId, source, content)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, model.LogSource, []byte) error); ok {
		r0 = rf(taskId, source, content)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBy provides a mock function with given fields: taskId, options
func (_m *LogManager) GetBy(taskId string, options model.LogOptions) ([]model.Log, error) {
	ret := _m.Called(taskId, options)

	var r0 []model.Log
	if rf, ok := ret.Get(0).(func(string, model.LogOptions) []model.Log); ok {
		r0 = rf(taskId, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Log)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, model.LogOptions) error); ok {
		r1 = rf(taskId, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewLogManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewLogManager creates a new instance of LogManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLogManager(t mockConstructorTestingTNewLogManager) *LogManager {
	mock := &LogManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
