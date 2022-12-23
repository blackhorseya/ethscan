// Code generated by mockery v2.15.0. DO NOT EDIT.

package repo

import (
	contextx "github.com/blackhorseya/ethscan/pkg/contextx"
	mock "github.com/stretchr/testify/mock"

	model "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/model"
)

// MockIRepo is an autogenerated mock type for the IRepo type
type MockIRepo struct {
	mock.Mock
}

// FetchTxByHash provides a mock function with given fields: ctx, hash
func (_m *MockIRepo) FetchTxByHash(ctx contextx.Contextx, hash string) (*model.Transaction, error) {
	ret := _m.Called(ctx, hash)

	var r0 *model.Transaction
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string) *model.Transaction); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Transaction)
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
