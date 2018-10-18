// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "github.com/mszostok/kyma/components/metadata-service/pkg/apis/istio/v1alpha2"
	scheme "github.com/mszostok/kyma/components/metadata-service/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DeniersGetter has a method to return a DenierInterface.
// A group's client should implement this interface.
type DeniersGetter interface {
	Deniers(namespace string) DenierInterface
}

// DenierInterface has methods to work with Denier resources.
type DenierInterface interface {
	Create(*v1alpha2.Denier) (*v1alpha2.Denier, error)
	Update(*v1alpha2.Denier) (*v1alpha2.Denier, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.Denier, error)
	List(opts v1.ListOptions) (*v1alpha2.DenierList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.Denier, err error)
	DenierExpansion
}

// deniers implements DenierInterface
type deniers struct {
	client rest.Interface
	ns     string
}

// newDeniers returns a Deniers
func newDeniers(c *IstioV1alpha2Client, namespace string) *deniers {
	return &deniers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the denier, and returns the corresponding denier object, and an error if there is any.
func (c *deniers) Get(name string, options v1.GetOptions) (result *v1alpha2.Denier, err error) {
	result = &v1alpha2.Denier{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deniers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Deniers that match those selectors.
func (c *deniers) List(opts v1.ListOptions) (result *v1alpha2.DenierList, err error) {
	result = &v1alpha2.DenierList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deniers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested deniers.
func (c *deniers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("deniers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a denier and creates it.  Returns the server's representation of the denier, and an error, if there is any.
func (c *deniers) Create(denier *v1alpha2.Denier) (result *v1alpha2.Denier, err error) {
	result = &v1alpha2.Denier{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("deniers").
		Body(denier).
		Do().
		Into(result)
	return
}

// Update takes the representation of a denier and updates it. Returns the server's representation of the denier, and an error, if there is any.
func (c *deniers) Update(denier *v1alpha2.Denier) (result *v1alpha2.Denier, err error) {
	result = &v1alpha2.Denier{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deniers").
		Name(denier.Name).
		Body(denier).
		Do().
		Into(result)
	return
}

// Delete takes name of the denier and deletes it. Returns an error if one occurs.
func (c *deniers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deniers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *deniers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deniers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched denier.
func (c *deniers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.Denier, err error) {
	result = &v1alpha2.Denier{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("deniers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
