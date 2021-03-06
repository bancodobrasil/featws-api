// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/bancodobrasil/featws-api/models"
	mock "github.com/stretchr/testify/mock"

	repository "github.com/bancodobrasil/featws-api/repository"
)

// Rulesheets is an autogenerated mock type for the Rulesheets type
type Rulesheets struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, entity
func (_m *Rulesheets) Count(ctx context.Context, entity interface{}) (int64, error) {
	ret := _m.Called(ctx, entity)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) int64); ok {
		r0 = rf(ctx, entity)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, entity
func (_m *Rulesheets) Create(ctx context.Context, entity *models.Rulesheet) error {
	ret := _m.Called(ctx, entity)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Rulesheet) error); ok {
		r0 = rf(ctx, entity)
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

// Find provides a mock function with given fields: ctx, entity, options
func (_m *Rulesheets) Find(ctx context.Context, entity interface{}, options *repository.FindOptions) ([]*models.Rulesheet, error) {
	ret := _m.Called(ctx, entity, options)

	var r0 []*models.Rulesheet
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, *repository.FindOptions) []*models.Rulesheet); ok {
		r0 = rf(ctx, entity, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Rulesheet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, *repository.FindOptions) error); ok {
		r1 = rf(ctx, entity, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, id
func (_m *Rulesheets) Get(ctx context.Context, id string) (*models.Rulesheet, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Rulesheet
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Rulesheet); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Rulesheet)
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
func (_m *Rulesheets) Update(ctx context.Context, entity models.Rulesheet) (*models.Rulesheet, error) {
	ret := _m.Called(ctx, entity)

	var r0 *models.Rulesheet
	if rf, ok := ret.Get(0).(func(context.Context, models.Rulesheet) *models.Rulesheet); ok {
		r0 = rf(ctx, entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Rulesheet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.Rulesheet) error); ok {
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
