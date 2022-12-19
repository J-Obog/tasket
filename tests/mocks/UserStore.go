// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	model "github.com/J-Obog/tasket/src/model"
	mock "github.com/stretchr/testify/mock"
)

// UserStore is an autogenerated mock type for the UserStore type
type UserStore struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *UserStore) Delete(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *UserStore) Get(id string) (*model.User, error) {
	ret := _m.Called(id)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
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

// GetByEmail provides a mock function with given fields: email
func (_m *UserStore) GetByEmail(email string) (*model.User, error) {
	ret := _m.Called(email)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: user
func (_m *UserStore) Insert(user model.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUserStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserStore creates a new instance of UserStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserStore(t mockConstructorTestingTNewUserStore) *UserStore {
	mock := &UserStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}