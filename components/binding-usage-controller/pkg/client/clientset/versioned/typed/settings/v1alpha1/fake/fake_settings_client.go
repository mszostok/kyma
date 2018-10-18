// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/mszostok/kyma/components/binding-usage-controller/pkg/client/clientset/versioned/typed/settings/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSettingsV1alpha1 struct {
	*testing.Fake
}

func (c *FakeSettingsV1alpha1) PodPresets(namespace string) v1alpha1.PodPresetInterface {
	return &FakePodPresets{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSettingsV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
