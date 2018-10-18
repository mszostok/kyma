// Code generated by mockery v1.0.0
package automock

import internal "github.com/mszostok/kyma/components/remote-environment-broker/internal"
import mock "github.com/stretchr/testify/mock"

// RemoteEnvironmentRemover is an autogenerated mock type for the RemoteEnvironmentRemover type
type RemoteEnvironmentRemover struct {
	mock.Mock
}

// Remove provides a mock function with given fields: name
func (_m *RemoteEnvironmentRemover) Remove(name internal.RemoteEnvironmentName) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(internal.RemoteEnvironmentName) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
