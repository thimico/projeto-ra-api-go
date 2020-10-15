// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"
	model "projeto-ra-api-go/pkg/api/model"

	mock "github.com/stretchr/testify/mock"
)

// ComplainRepository is an autogenerated mock type for the ComplainRepository type
type ComplainRepository struct {
	mock.Mock
}

// DeleteById provides a mock function with given fields: ctx, id
func (_m *ComplainRepository) DeleteById(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindById provides a mock function with given fields: ctx, id
func (_m *ComplainRepository) FindById(ctx context.Context, id string) (*model.Complain, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.Complain
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Complain); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Complain)
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

// FindByParam provides a mock function with given fields: ctx, param
func (_m *ComplainRepository) FindByParam(ctx context.Context, param *model.ComplainIn) ([]model.Complain, error) {
	ret := _m.Called(ctx, param)

	var r0 []model.Complain
	if rf, ok := ret.Get(0).(func(context.Context, *model.ComplainIn) []model.Complain); ok {
		r0 = rf(ctx, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Complain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.ComplainIn) error); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: parentContext, complain
func (_m *ComplainRepository) Save(parentContext context.Context, complain *model.Complain) (string, error) {
	ret := _m.Called(parentContext, complain)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *model.Complain) string); ok {
		r0 = rf(parentContext, complain)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.Complain) error); ok {
		r1 = rf(parentContext, complain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, p, id
func (_m *ComplainRepository) Update(ctx context.Context, p *model.Complain, id string) error {
	ret := _m.Called(ctx, p, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Complain, string) error); ok {
		r0 = rf(ctx, p, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
