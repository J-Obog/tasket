// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	types "github.com/J-Obog/tasket/src/types"
	mock "github.com/stretchr/testify/mock"
)

// IUserResource is an autogenerated mock type for the IUserResource type
type IUserResource struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: req
func (_m *IUserResource) CreateUser(req types.RestRequest) types.RestResponse {
	ret := _m.Called(req)

	var r0 types.RestResponse
	if rf, ok := ret.Get(0).(func(types.RestRequest) types.RestResponse); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(types.RestResponse)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: req
func (_m *IUserResource) DeleteUser(req types.RestRequest) types.RestResponse {
	ret := _m.Called(req)

	var r0 types.RestResponse
	if rf, ok := ret.Get(0).(func(types.RestRequest) types.RestResponse); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(types.RestResponse)
	}

	return r0
}

// GetUser provides a mock function with given fields: req
func (_m *IUserResource) GetUser(req types.RestRequest) types.RestResponse {
	ret := _m.Called(req)

	var r0 types.RestResponse
	if rf, ok := ret.Get(0).(func(types.RestRequest) types.RestResponse); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(types.RestResponse)
	}

	return r0
}

type mockConstructorTestingTNewIUserResource interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUserResource creates a new instance of IUserResource. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUserResource(t mockConstructorTestingTNewIUserResource) *IUserResource {
	mock := &IUserResource{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}