// Code generated by mockery v2.15.0. DO NOT EDIT.

package repo

import (
	contextx "github.com/blackhorseya/ethscan/pkg/contextx"
	kafka "github.com/confluentinc/confluent-kafka-go/kafka"

	mock "github.com/stretchr/testify/mock"

	model "github.com/blackhorseya/ethscan/pkg/entity/domain/block/model"
)

// MockIRepo is an autogenerated mock type for the IRepo type
type MockIRepo struct {
	mock.Mock
}

// CountRecord provides a mock function with given fields: ctx, condition
func (_m *MockIRepo) CountRecord(ctx contextx.Contextx, condition ListRecordCondition) (int, error) {
	ret := _m.Called(ctx, condition)

	var r0 int
	if rf, ok := ret.Get(0).(func(contextx.Contextx, ListRecordCondition) int); ok {
		r0 = rf(ctx, condition)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, ListRecordCondition) error); ok {
		r1 = rf(ctx, condition)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRecord provides a mock function with given fields: ctx, record
func (_m *MockIRepo) CreateRecord(ctx contextx.Contextx, record *model.BlockRecord) error {
	ret := _m.Called(ctx, record)

	var r0 error
	if rf, ok := ret.Get(0).(func(contextx.Contextx, *model.BlockRecord) error); ok {
		r0 = rf(ctx, record)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchCurrentHeight provides a mock function with given fields: ctx
func (_m *MockIRepo) FetchCurrentHeight(ctx contextx.Contextx) (uint64, error) {
	ret := _m.Called(ctx)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(contextx.Contextx) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchRecordByHeight provides a mock function with given fields: ctx, height
func (_m *MockIRepo) FetchRecordByHeight(ctx contextx.Contextx, height uint64) (*model.BlockRecord, error) {
	ret := _m.Called(ctx, height)

	var r0 *model.BlockRecord
	if rf, ok := ret.Get(0).(func(contextx.Contextx, uint64) *model.BlockRecord); ok {
		r0 = rf(ctx, height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.BlockRecord)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, uint64) error); ok {
		r1 = rf(ctx, height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestRecord provides a mock function with given fields: ctx
func (_m *MockIRepo) GetLatestRecord(ctx contextx.Contextx) (*model.BlockRecord, error) {
	ret := _m.Called(ctx)

	var r0 *model.BlockRecord
	if rf, ok := ret.Get(0).(func(contextx.Contextx) *model.BlockRecord); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.BlockRecord)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRecordByHash provides a mock function with given fields: ctx, hash
func (_m *MockIRepo) GetRecordByHash(ctx contextx.Contextx, hash string) (*model.BlockRecord, error) {
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

// ListRecord provides a mock function with given fields: ctx, condition
func (_m *MockIRepo) ListRecord(ctx contextx.Contextx, condition ListRecordCondition) ([]*model.BlockRecord, error) {
	ret := _m.Called(ctx, condition)

	var r0 []*model.BlockRecord
	if rf, ok := ret.Get(0).(func(contextx.Contextx, ListRecordCondition) []*model.BlockRecord); ok {
		r0 = rf(ctx, condition)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.BlockRecord)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, ListRecordCondition) error); ok {
		r1 = rf(ctx, condition)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProduceRecord provides a mock function with given fields: ctx, record, delivery
func (_m *MockIRepo) ProduceRecord(ctx contextx.Contextx, record *model.BlockRecord, delivery chan kafka.Event) error {
	ret := _m.Called(ctx, record, delivery)

	var r0 error
	if rf, ok := ret.Get(0).(func(contextx.Contextx, *model.BlockRecord, chan kafka.Event) error); ok {
		r0 = rf(ctx, record, delivery)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockIRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockIRepo creates a new instance of MockIRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockIRepo(t mockConstructorTestingTNewMockIRepo) *MockIRepo {
	mock := &MockIRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}