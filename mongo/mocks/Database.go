// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	mongo "application/project/mongo"

	mock "github.com/stretchr/testify/mock"
)

// Database is an autogenerated mock type for the Database type
type Database struct {
	mock.Mock
}

// Client provides a mock function with given fields:
func (_m *Database) Client() *mongo.MongoClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Client")
	}

	var r0 *mongo.MongoClient
	if rf, ok := ret.Get(0).(func() *mongo.MongoClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.MongoClient)
		}
	}

	return r0
}

// Collection provides a mock function with given fields: name
func (_m *Database) Collection(name string) *mongo.Collection {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for Collection")
	}

	var r0 *mongo.Collection
	if rf, ok := ret.Get(0).(func(string) *mongo.Collection); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.Collection)
		}
	}

	return r0
}

// NewDatabase creates a new instance of Database. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDatabase(t interface {
	mock.TestingT
	Cleanup(func())
}) *Database {
	mock := &Database{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
