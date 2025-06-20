// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package stub

import (
	"ergo.services/ergo/gen"
	mock "github.com/stretchr/testify/mock"
)

// NewResolver creates a new instance of Resolver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewResolver(t interface {
	mock.TestingT
	Cleanup(func())
}) *Resolver {
	mock := &Resolver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Resolver is an autogenerated mock type for the Resolver type
type Resolver struct {
	mock.Mock
}

type Resolver_Expecter struct {
	mock *mock.Mock
}

func (_m *Resolver) EXPECT() *Resolver_Expecter {
	return &Resolver_Expecter{mock: &_m.Mock}
}

// Resolve provides a mock function for the type Resolver
func (_mock *Resolver) Resolve(node gen.Atom) ([]gen.Route, error) {
	ret := _mock.Called(node)

	if len(ret) == 0 {
		panic("no return value specified for Resolve")
	}

	var r0 []gen.Route
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(gen.Atom) ([]gen.Route, error)); ok {
		return returnFunc(node)
	}
	if returnFunc, ok := ret.Get(0).(func(gen.Atom) []gen.Route); ok {
		r0 = returnFunc(node)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gen.Route)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(gen.Atom) error); ok {
		r1 = returnFunc(node)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// Resolver_Resolve_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Resolve'
type Resolver_Resolve_Call struct {
	*mock.Call
}

// Resolve is a helper method to define mock.On call
//   - node
func (_e *Resolver_Expecter) Resolve(node interface{}) *Resolver_Resolve_Call {
	return &Resolver_Resolve_Call{Call: _e.mock.On("Resolve", node)}
}

func (_c *Resolver_Resolve_Call) Run(run func(node gen.Atom)) *Resolver_Resolve_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(gen.Atom))
	})
	return _c
}

func (_c *Resolver_Resolve_Call) Return(routes []gen.Route, err error) *Resolver_Resolve_Call {
	_c.Call.Return(routes, err)
	return _c
}

func (_c *Resolver_Resolve_Call) RunAndReturn(run func(node gen.Atom) ([]gen.Route, error)) *Resolver_Resolve_Call {
	_c.Call.Return(run)
	return _c
}

// ResolveApplication provides a mock function for the type Resolver
func (_mock *Resolver) ResolveApplication(name gen.Atom) ([]gen.ApplicationRoute, error) {
	ret := _mock.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for ResolveApplication")
	}

	var r0 []gen.ApplicationRoute
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(gen.Atom) ([]gen.ApplicationRoute, error)); ok {
		return returnFunc(name)
	}
	if returnFunc, ok := ret.Get(0).(func(gen.Atom) []gen.ApplicationRoute); ok {
		r0 = returnFunc(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gen.ApplicationRoute)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(gen.Atom) error); ok {
		r1 = returnFunc(name)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// Resolver_ResolveApplication_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResolveApplication'
type Resolver_ResolveApplication_Call struct {
	*mock.Call
}

// ResolveApplication is a helper method to define mock.On call
//   - name
func (_e *Resolver_Expecter) ResolveApplication(name interface{}) *Resolver_ResolveApplication_Call {
	return &Resolver_ResolveApplication_Call{Call: _e.mock.On("ResolveApplication", name)}
}

func (_c *Resolver_ResolveApplication_Call) Run(run func(name gen.Atom)) *Resolver_ResolveApplication_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(gen.Atom))
	})
	return _c
}

func (_c *Resolver_ResolveApplication_Call) Return(applicationRoutes []gen.ApplicationRoute, err error) *Resolver_ResolveApplication_Call {
	_c.Call.Return(applicationRoutes, err)
	return _c
}

func (_c *Resolver_ResolveApplication_Call) RunAndReturn(run func(name gen.Atom) ([]gen.ApplicationRoute, error)) *Resolver_ResolveApplication_Call {
	_c.Call.Return(run)
	return _c
}

// ResolveProxy provides a mock function for the type Resolver
func (_mock *Resolver) ResolveProxy(node gen.Atom) ([]gen.ProxyRoute, error) {
	ret := _mock.Called(node)

	if len(ret) == 0 {
		panic("no return value specified for ResolveProxy")
	}

	var r0 []gen.ProxyRoute
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(gen.Atom) ([]gen.ProxyRoute, error)); ok {
		return returnFunc(node)
	}
	if returnFunc, ok := ret.Get(0).(func(gen.Atom) []gen.ProxyRoute); ok {
		r0 = returnFunc(node)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gen.ProxyRoute)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(gen.Atom) error); ok {
		r1 = returnFunc(node)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// Resolver_ResolveProxy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResolveProxy'
type Resolver_ResolveProxy_Call struct {
	*mock.Call
}

// ResolveProxy is a helper method to define mock.On call
//   - node
func (_e *Resolver_Expecter) ResolveProxy(node interface{}) *Resolver_ResolveProxy_Call {
	return &Resolver_ResolveProxy_Call{Call: _e.mock.On("ResolveProxy", node)}
}

func (_c *Resolver_ResolveProxy_Call) Run(run func(node gen.Atom)) *Resolver_ResolveProxy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(gen.Atom))
	})
	return _c
}

func (_c *Resolver_ResolveProxy_Call) Return(proxyRoutes []gen.ProxyRoute, err error) *Resolver_ResolveProxy_Call {
	_c.Call.Return(proxyRoutes, err)
	return _c
}

func (_c *Resolver_ResolveProxy_Call) RunAndReturn(run func(node gen.Atom) ([]gen.ProxyRoute, error)) *Resolver_ResolveProxy_Call {
	_c.Call.Return(run)
	return _c
}
