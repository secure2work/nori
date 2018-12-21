// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import config "github.com/secure2work/nori/core/config"
import context "context"
import meta "github.com/secure2work/nori/core/plugins/meta"
import mock "github.com/stretchr/testify/mock"
import plugin "github.com/secure2work/nori/core/plugins/plugin"

// Plugin is an autogenerated mock type for the Plugin type
type Plugin struct {
	mock.Mock
}

// Init provides a mock function with given fields: ctx, _a1
func (_m *Plugin) Init(ctx context.Context, _a1 config.Manager) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, config.Manager) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Instance provides a mock function with given fields:
func (_m *Plugin) Instance() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// Meta provides a mock function with given fields:
func (_m *Plugin) Meta() meta.Meta {
	ret := _m.Called()

	var r0 meta.Meta
	if rf, ok := ret.Get(0).(func() meta.Meta); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(meta.Meta)
		}
	}

	return r0
}

// Start provides a mock function with given fields: ctx, registry
func (_m *Plugin) Start(ctx context.Context, registry plugin.Registry) error {
	ret := _m.Called(ctx, registry)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, plugin.Registry) error); ok {
		r0 = rf(ctx, registry)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields: ctx, registry
func (_m *Plugin) Stop(ctx context.Context, registry plugin.Registry) error {
	ret := _m.Called(ctx, registry)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, plugin.Registry) error); ok {
		r0 = rf(ctx, registry)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
