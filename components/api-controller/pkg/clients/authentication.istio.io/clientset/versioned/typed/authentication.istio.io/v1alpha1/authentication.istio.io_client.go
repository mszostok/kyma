// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/mszostok/kyma/components/api-controller/pkg/apis/authentication.istio.io/v1alpha1"
	"github.com/mszostok/kyma/components/api-controller/pkg/clients/authentication.istio.io/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type AuthenticationV1alpha1Interface interface {
	RESTClient() rest.Interface
	PoliciesGetter
}

// AuthenticationV1alpha1Client is used to interact with features provided by the authentication.istio.io group.
type AuthenticationV1alpha1Client struct {
	restClient rest.Interface
}

func (c *AuthenticationV1alpha1Client) Policies(namespace string) PolicyInterface {
	return newPolicies(c, namespace)
}

// NewForConfig creates a new AuthenticationV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*AuthenticationV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &AuthenticationV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new AuthenticationV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *AuthenticationV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new AuthenticationV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *AuthenticationV1alpha1Client {
	return &AuthenticationV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *AuthenticationV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
