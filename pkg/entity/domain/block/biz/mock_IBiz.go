// Code generated by mockery v2.15.0. DO NOT EDIT.

package biz

import (
	contextx "github.com/blackhorseya/portto/pkg/contextx"
	mock "github.com/stretchr/testify/mock"

	model "github.com/blackhorseya/portto/pkg/entity/domain/block/model"
)

// MockIBiz is an autogenerated mock type for the IBiz type
type MockIBiz struct {
	mock.Mock
}

// GetByHash provides a mock function with given fields: ctx, hash
func (_m *MockIBiz) GetByHash(ctx contextx.Contextx, hash string) (*model.BlockRecord, error) {
	ret := _m.Called(ctx, hash)

	var r0 *model.BlockRecord
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string) *model.BlockRecord); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.BlockRecord)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, string) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, cond
func (_m *MockIBiz) List(ctx contextx.Contextx, cond ListCondition) ([]*model.BlockRecord, int, error) {
	ret := _m.Called(ctx, cond)

	var r0 []*model.BlockRecord
	if rf, ok := ret.Get(0).(func(contextx.Contextx, ListCondition) []*model.BlockRecord); ok {
		r0 = rf(ctx, cond)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.BlockRecord)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(contextx.Contextx, ListCondition) int); ok {
		r1 = rf(ctx, cond)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(contextx.Contextx, ListCondition) error); ok {
		r2 = rf(ctx, cond)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ScanBlock provides a mock function with given fields: ctx, start
func (_m *MockIBiz) ScanBlock(ctx contextx.Contextx, start uint64) (uint64, chan *model.BlockRecord, chan struct{}, chan error) {
	ret := _m.Called(ctx, start)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(contextx.Contextx, uint64) uint64); ok {
		r0 = rf(ctx, start)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 chan *model.BlockRecord
	if rf, ok := ret.Get(1).(func(contextx.Contextx, uint64) chan *model.BlockRecord); ok {
		r1 = rf(ctx, start)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(chan *model.BlockRecord)
		}
	}

	var r2 chan struct{}
	if rf, ok := ret.Get(2).(func(contextx.Contextx, uint64) chan struct{}); ok {
		r2 = rf(ctx, start)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(chan struct{})
		}
	}

	var r3 chan error
	if rf, ok := ret.Get(3).(func(contextx.Contextx, uint64) chan error); ok {
		r3 = rf(ctx, start)
	} else {
		if ret.Get(3) != nil {
			r3 = ret.Get(3).(chan error)
		}
	}

	return r0, r1, r2, r3
}

type mockConstructorTestingTNewMockIBiz interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockIBiz creates a new instance of MockIBiz. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockIBiz(t mockConstructorTestingTNewMockIBiz) *MockIBiz {
	mock := &MockIBiz{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
