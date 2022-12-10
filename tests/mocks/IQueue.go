// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	types "github.com/J-Obog/tasket/src/types"
	mock "github.com/stretchr/testify/mock"
)

// IQueue is an autogenerated mock type for the IQueue type
type IQueue struct {
	mock.Mock
}

// Pull provides a mock function with given fields:
func (_m *IQueue) Pull() (types.EventMessage, error) {
	ret := _m.Called()

	var r0 types.EventMessage
	if rf, ok := ret.Get(0).(func() types.EventMessage); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.EventMessage)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Push provides a mock function with given fields: message
func (_m *IQueue) Push(message types.EventMessage) error {
	ret := _m.Called(message)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.EventMessage) error); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIQueue interface {
	mock.TestingT
	Cleanup(func())
}

// NewIQueue creates a new instance of IQueue. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIQueue(t mockConstructorTestingTNewIQueue) *IQueue {
	mock := &IQueue{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
