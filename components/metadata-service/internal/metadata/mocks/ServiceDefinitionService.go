// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import apperrors "github.com/mszostok/kyma/components/metadata-service/internal/apperrors"
import metadata "github.com/mszostok/kyma/components/metadata-service/internal/metadata"
import mock "github.com/stretchr/testify/mock"
import serviceapi "github.com/mszostok/kyma/components/metadata-service/internal/metadata/serviceapi"

// ServiceDefinitionService is an autogenerated mock type for the ServiceDefinitionService type
type ServiceDefinitionService struct {
	mock.Mock
}

// Create provides a mock function with given fields: remoteEnvironment, serviceDefinition
func (_m *ServiceDefinitionService) Create(remoteEnvironment string, serviceDefinition *metadata.ServiceDefinition) (string, apperrors.AppError) {
	ret := _m.Called(remoteEnvironment, serviceDefinition)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, *metadata.ServiceDefinition) string); ok {
		r0 = rf(remoteEnvironment, serviceDefinition)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(string, *metadata.ServiceDefinition) apperrors.AppError); ok {
		r1 = rf(remoteEnvironment, serviceDefinition)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// Delete provides a mock function with given fields: remoteEnvironment, id
func (_m *ServiceDefinitionService) Delete(remoteEnvironment string, id string) apperrors.AppError {
	ret := _m.Called(remoteEnvironment, id)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, string) apperrors.AppError); ok {
		r0 = rf(remoteEnvironment, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// GetAPI provides a mock function with given fields: remoteEnvironment, serviceId
func (_m *ServiceDefinitionService) GetAPI(remoteEnvironment string, serviceId string) (*serviceapi.API, apperrors.AppError) {
	ret := _m.Called(remoteEnvironment, serviceId)

	var r0 *serviceapi.API
	if rf, ok := ret.Get(0).(func(string, string) *serviceapi.API); ok {
		r0 = rf(remoteEnvironment, serviceId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*serviceapi.API)
		}
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(string, string) apperrors.AppError); ok {
		r1 = rf(remoteEnvironment, serviceId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: remoteEnvironment
func (_m *ServiceDefinitionService) GetAll(remoteEnvironment string) ([]metadata.ServiceDefinition, apperrors.AppError) {
	ret := _m.Called(remoteEnvironment)

	var r0 []metadata.ServiceDefinition
	if rf, ok := ret.Get(0).(func(string) []metadata.ServiceDefinition); ok {
		r0 = rf(remoteEnvironment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]metadata.ServiceDefinition)
		}
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(string) apperrors.AppError); ok {
		r1 = rf(remoteEnvironment)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: remoteEnvironment, id
func (_m *ServiceDefinitionService) GetByID(remoteEnvironment string, id string) (metadata.ServiceDefinition, apperrors.AppError) {
	ret := _m.Called(remoteEnvironment, id)

	var r0 metadata.ServiceDefinition
	if rf, ok := ret.Get(0).(func(string, string) metadata.ServiceDefinition); ok {
		r0 = rf(remoteEnvironment, id)
	} else {
		r0 = ret.Get(0).(metadata.ServiceDefinition)
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(string, string) apperrors.AppError); ok {
		r1 = rf(remoteEnvironment, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// Update provides a mock function with given fields: remoteEnvironment, id, serviceDef
func (_m *ServiceDefinitionService) Update(remoteEnvironment string, id string, serviceDef *metadata.ServiceDefinition) (metadata.ServiceDefinition, apperrors.AppError) {
	ret := _m.Called(remoteEnvironment, id, serviceDef)

	var r0 metadata.ServiceDefinition
	if rf, ok := ret.Get(0).(func(string, string, *metadata.ServiceDefinition) metadata.ServiceDefinition); ok {
		r0 = rf(remoteEnvironment, id, serviceDef)
	} else {
		r0 = ret.Get(0).(metadata.ServiceDefinition)
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(string, string, *metadata.ServiceDefinition) apperrors.AppError); ok {
		r1 = rf(remoteEnvironment, id, serviceDef)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}
