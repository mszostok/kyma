// Code generated by mockery v1.0.0
package automock

import controller "github.com/mszostok/kyma/components/binding-usage-controller/internal/controller"
import mock "github.com/stretchr/testify/mock"

// KindsSupervisors is an autogenerated mock type for the KindsSupervisors type
type KindsSupervisors struct {
	mock.Mock
}

// Get provides a mock function with given fields: kind
func (_m *KindsSupervisors) Get(kind controller.Kind) (controller.KubernetesResourceSupervisor, error) {
	ret := _m.Called(kind)

	var r0 controller.KubernetesResourceSupervisor
	if rf, ok := ret.Get(0).(func(controller.Kind) controller.KubernetesResourceSupervisor); ok {
		r0 = rf(kind)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(controller.KubernetesResourceSupervisor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(controller.Kind) error); ok {
		r1 = rf(kind)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

