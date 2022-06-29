// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	dtos "github.com/bancodobrasil/featws-api/dtos"
	mock "github.com/stretchr/testify/mock"
)

// Rulesheets is an autogenerated mock type for the Rulesheets type
type Rulesheets struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *Rulesheets) Create(_a0 context.Context, _a1 *dtos.Rulesheet) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *dtos.Rulesheet) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Rulesheets) Delete(ctx context.Context, id string) (bool, error) {
	ret := _m.Called(ctx, id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: ctx, filter
func (_m *Rulesheets) Find(ctx context.Context, filter interface{}) ([]*dtos.Rulesheet, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*dtos.Rulesheet
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) []*dtos.Rulesheet); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dtos.Rulesheet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, id
func (_m *Rulesheets) Get(ctx context.Context, id string) (*dtos.Rulesheet, error) {
	ret := _m.Called(ctx, id)

	var r0 *dtos.Rulesheet
	if rf, ok := ret.Get(0).(func(context.Context, string) *dtos.Rulesheet); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dtos.Rulesheet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, entity
func (_m *Rulesheets) Update(ctx context.Context, entity dtos.Rulesheet) (*dtos.Rulesheet, error) {
	ret := _m.Called(ctx, entity)

	var r0 *dtos.Rulesheet
	if rf, ok := ret.Get(0).(func(context.Context, dtos.Rulesheet) *dtos.Rulesheet); ok {
		r0 = rf(ctx, entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dtos.Rulesheet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, dtos.Rulesheet) error); ok {
		r1 = rf(ctx, entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewRulesheetsT interface {
	mock.TestingT
	Cleanup(func())
}

// NewRulesheets creates a new instance of Rulesheets. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRulesheets(t NewRulesheetsT) *Rulesheets {
	mock := &Rulesheets{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
