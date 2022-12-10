// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ITimeProvider is an autogenerated mock type for the ITimeProvider type
type ITimeProvider struct {
	mock.Mock
}

// Now provides a mock function with given fields:
func (_m *ITimeProvider) Now() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

type mockConstructorTestingTNewITimeProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewITimeProvider creates a new instance of ITimeProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewITimeProvider(t mockConstructorTestingTNewITimeProvider) *ITimeProvider {
	mock := &ITimeProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}