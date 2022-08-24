// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	dtos "github.com/bancodobrasil/featws-api/dtos"
	mock "github.com/stretchr/testify/mock"
)

// Gitlab is an autogenerated mock type for the Gitlab type
type Gitlab struct {
	mock.Mock
}

// Fill provides a mock function with given fields: rulesheet
func (_m *Gitlab) Fill(rulesheet *dtos.Rulesheet) error {
	ret := _m.Called(rulesheet)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dtos.Rulesheet) error); ok {
		r0 = rf(rulesheet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: rulesheet, commitMessage
func (_m *Gitlab) Save(rulesheet *dtos.Rulesheet, commitMessage string) error {
	ret := _m.Called(rulesheet, commitMessage)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dtos.Rulesheet, string) error); ok {
		r0 = rf(rulesheet, commitMessage)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewGitlab interface {
	mock.TestingT
	Cleanup(func())
}

// NewGitlab creates a new instance of Gitlab. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGitlab(t mockConstructorTestingTNewGitlab) *Gitlab {
	mock := &Gitlab{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
