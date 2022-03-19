// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// RedisClientInterface is an autogenerated mock type for the RedisClientInterface type
type RedisClientInterface struct {
	mock.Mock
}

// GetData provides a mock function with given fields: key
func (_m *RedisClientInterface) GetData(key string) string {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// SetData provides a mock function with given fields: key, value, exp
func (_m *RedisClientInterface) SetData(key string, value string, exp time.Duration) error {
	ret := _m.Called(key, value, exp)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, time.Duration) error); ok {
		r0 = rf(key, value, exp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
