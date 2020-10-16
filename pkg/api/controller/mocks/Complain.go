// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"
	model "projeto-ra-api-go/pkg/api/model"

	mock "github.com/stretchr/testify/mock"
)

// Complain is an autogenerated mock type for the Complain type
type Complain struct {
	mock.Mock
}

type ComplainFn struct {
	SaveFn		func(context.Context, *model.ComplainIn) (*model.ComplainOut, error)
	DeleteByIdFn func(context.Context, string) error
	UpdateFn func(context.Context, model.ComplainIn, string) error
	FindByIdFn func(context.Context, string) (*model.ComplainOut, error)
	FindByIdWithExternalFn func(context.Context, string) (*model.ComplainOut, error)
	FindByParamFn func(context.Context, *model.ComplainIn) ([]model.ComplainOut, error)
	FindByParamWithExternalFn func(context.Context, *model.ComplainIn) ([]model.ComplainOut, error)
}

// DeleteById provides a mock function with given fields: ctx, id
func (_m *Complain) DeleteById(ctx context.Context, id string) error {
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
func (_m *Complain) FindById(ctx context.Context, id string) (*model.ComplainOut, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.ComplainOut
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.ComplainOut); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ComplainOut)
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

// FindByIdWithExternal provides a mock function with given fields: ctx, id
func (_m *Complain) FindByIdWithExternal(ctx context.Context, id string) (*model.ComplainOut, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.ComplainOut
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.ComplainOut); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ComplainOut)
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
func (_m *Complain) FindByParam(ctx context.Context, param *model.ComplainIn) ([]model.ComplainOut, error) {
	ret := _m.Called(ctx, param)

	var r0 []model.ComplainOut
	if rf, ok := ret.Get(0).(func(context.Context, *model.ComplainIn) []model.ComplainOut); ok {
		r0 = rf(ctx, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.ComplainOut)
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

// FindByParamWithExternal provides a mock function with given fields: ctx, param
func (_m *Complain) FindByParamWithExternal(ctx context.Context, param *model.ComplainIn) ([]model.ComplainOut, error) {
	ret := _m.Called(ctx, param)

	var r0 []model.ComplainOut
	if rf, ok := ret.Get(0).(func(context.Context, *model.ComplainIn) []model.ComplainOut); ok {
		r0 = rf(ctx, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.ComplainOut)
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
func (_m *Complain) Save(parentContext context.Context, complain *model.ComplainIn) (*model.ComplainOut, error) {
	ret := _m.Called(parentContext, complain)

	var r0 *model.ComplainOut
	if rf, ok := ret.Get(0).(func(context.Context, *model.ComplainIn) *model.ComplainOut); ok {
		r0 = rf(parentContext, complain)
	} else {
		r0 = ret.Get(0).(*model.ComplainOut)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.ComplainIn) error); ok {
		r1 = rf(parentContext, complain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, p, id
func (_m *Complain) Update(ctx context.Context, p model.ComplainIn, id string) error {
	ret := _m.Called(ctx, p, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.ComplainIn, string) error); ok {
		r0 = rf(ctx, p, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
