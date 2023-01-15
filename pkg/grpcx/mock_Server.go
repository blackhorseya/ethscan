// Code generated by mockery v2.16.0. DO NOT EDIT.

package grpcx

import mock "github.com/stretchr/testify/mock"

// MockServer is an autogenerated mock type for the Server type
type MockServer struct {
	mock.Mock
}

// Start provides a mock function with given fields:
func (_m *MockServer) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *MockServer) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockServer creates a new instance of MockServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockServer(t mockConstructorTestingTNewMockServer) *MockServer {
	mock := &MockServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
