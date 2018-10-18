// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import httputil "net/http/httputil"
import mock "github.com/stretchr/testify/mock"
import proxycache "github.com/mszostok/kyma/components/gateway/internal/proxy/proxycache"

// HTTPProxyCache is an autogenerated mock type for the HTTPProxyCache type
type HTTPProxyCache struct {
	mock.Mock
}

// Add provides a mock function with given fields: id, oauthUrl, clientId, clientSecret, proxy
func (_m *HTTPProxyCache) Add(id string, oauthUrl string, clientId string, clientSecret string, proxy *httputil.ReverseProxy) *proxycache.Proxy {
	ret := _m.Called(id, oauthUrl, clientId, clientSecret, proxy)

	var r0 *proxycache.Proxy
	if rf, ok := ret.Get(0).(func(string, string, string, string, *httputil.ReverseProxy) *proxycache.Proxy); ok {
		r0 = rf(id, oauthUrl, clientId, clientSecret, proxy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proxycache.Proxy)
		}
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *HTTPProxyCache) Get(id string) (*proxycache.Proxy, bool) {
	ret := _m.Called(id)

	var r0 *proxycache.Proxy
	if rf, ok := ret.Get(0).(func(string) *proxycache.Proxy); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proxycache.Proxy)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}
