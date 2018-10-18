// Code generated by mockery v1.0.0
package automock

import gqlschema "github.com/mszostok/kyma/components/ui-api-layer/internal/gqlschema"
import mock "github.com/stretchr/testify/mock"

import v1alpha1 "github.com/mszostok/kyma/components/binding-usage-controller/pkg/apis/servicecatalog/v1alpha1"

// statusBindingUsageExtractor is an autogenerated mock type for the statusBindingUsageExtractor type
type statusBindingUsageExtractor struct {
	mock.Mock
}

// Status provides a mock function with given fields: conditions
func (_m *statusBindingUsageExtractor) Status(conditions []v1alpha1.ServiceBindingUsageCondition) gqlschema.ServiceBindingUsageStatus {
	ret := _m.Called(conditions)

	var r0 gqlschema.ServiceBindingUsageStatus
	if rf, ok := ret.Get(0).(func([]v1alpha1.ServiceBindingUsageCondition) gqlschema.ServiceBindingUsageStatus); ok {
		r0 = rf(conditions)
	} else {
		r0 = ret.Get(0).(gqlschema.ServiceBindingUsageStatus)
	}

	return r0
}
