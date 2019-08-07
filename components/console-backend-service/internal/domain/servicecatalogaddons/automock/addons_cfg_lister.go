// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import mock "github.com/stretchr/testify/mock"
import pager "github.com/kyma-project/kyma/components/console-backend-service/internal/pager"

import v1alpha1 "github.com/kyma-project/kyma/components/helm-broker/pkg/apis/networking/v1alpha3"

// addonsCfgLister is an autogenerated mock type for the addonsCfgLister type
type addonsCfgLister struct {
	mock.Mock
}

// List provides a mock function with given fields: pagingParams
func (_m *addonsCfgLister) List(pagingParams pager.PagingParams) ([]*v1alpha1.ClusterAddonsConfiguration, error) {
	ret := _m.Called(pagingParams)

	var r0 []*v1alpha1.ClusterAddonsConfiguration
	if rf, ok := ret.Get(0).(func(pager.PagingParams) []*v1alpha1.ClusterAddonsConfiguration); ok {
		r0 = rf(pagingParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1alpha1.ClusterAddonsConfiguration)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(pager.PagingParams) error); ok {
		r1 = rf(pagingParams)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
