// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/mszostok/kyma/components/installer/pkg/apis/installer/v1alpha1"
	"github.com/mszostok/kyma/components/installer/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type InstallerV1alpha1Interface interface {
	RESTClient() rest.Interface
	InstallationsGetter
}

// InstallerV1alpha1Client is used to interact with features provided by the installer.kyma-project.io group.
type InstallerV1alpha1Client struct {
	restClient rest.Interface
}

func (c *InstallerV1alpha1Client) Installations(namespace string) InstallationInterface {
	return newInstallations(c, namespace)
}

// NewForConfig creates a new InstallerV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*InstallerV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &InstallerV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new InstallerV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *InstallerV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new InstallerV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *InstallerV1alpha1Client {
	return &InstallerV1alpha1Client{c}
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
func (c *InstallerV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
