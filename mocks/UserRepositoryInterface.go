// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	entities "go-todo-api/entities"

	mock "github.com/stretchr/testify/mock"
)

// UserRepositoryInterface is an autogenerated mock type for the UserRepositoryInterface type
type UserRepositoryInterface struct {
	mock.Mock
}

// Delete provides a mock function with given fields: user
func (_m *UserRepositoryInterface) Delete(user *entities.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: user
func (_m *UserRepositoryInterface) FindAll(user []*entities.User) ([]*entities.User, error) {
	ret := _m.Called(user)

	var r0 []*entities.User
	if rf, ok := ret.Get(0).(func([]*entities.User) []*entities.User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*entities.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: user
func (_m *UserRepositoryInterface) FindByID(user *entities.User) (*entities.User, error) {
	ret := _m.Called(user)

	var r0 *entities.User
	if rf, ok := ret.Get(0).(func(*entities.User) *entities.User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entities.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: user
func (_m *UserRepositoryInterface) Insert(user *entities.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: user
func (_m *UserRepositoryInterface) Update(user *entities.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
