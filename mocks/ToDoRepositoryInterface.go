// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	entities "go-todo-api/entities"

	mock "github.com/stretchr/testify/mock"
)

// ToDoRepositoryInterface is an autogenerated mock type for the ToDoRepositoryInterface type
type ToDoRepositoryInterface struct {
	mock.Mock
}

// Delete provides a mock function with given fields: i
func (_m *ToDoRepositoryInterface) Delete(i *entities.ToDo) error {
	ret := _m.Called(i)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.ToDo) error); ok {
		r0 = rf(i)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: result
func (_m *ToDoRepositoryInterface) FindAll(result []*entities.ToDo) ([]*entities.ToDo, error) {
	ret := _m.Called(result)

	var r0 []*entities.ToDo
	if rf, ok := ret.Get(0).(func([]*entities.ToDo) []*entities.ToDo); ok {
		r0 = rf(result)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.ToDo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*entities.ToDo) error); ok {
		r1 = rf(result)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: result
func (_m *ToDoRepositoryInterface) FindByID(result *entities.ToDo) (*entities.ToDo, error) {
	ret := _m.Called(result)

	var r0 *entities.ToDo
	if rf, ok := ret.Get(0).(func(*entities.ToDo) *entities.ToDo); ok {
		r0 = rf(result)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.ToDo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entities.ToDo) error); ok {
		r1 = rf(result)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: i
func (_m *ToDoRepositoryInterface) Insert(i *entities.ToDo) error {
	ret := _m.Called(i)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.ToDo) error); ok {
		r0 = rf(i)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: i
func (_m *ToDoRepositoryInterface) Update(i *entities.ToDo) (*entities.ToDo, error) {
	ret := _m.Called(i)

	var r0 *entities.ToDo
	if rf, ok := ret.Get(0).(func(*entities.ToDo) *entities.ToDo); ok {
		r0 = rf(i)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.ToDo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entities.ToDo) error); ok {
		r1 = rf(i)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
