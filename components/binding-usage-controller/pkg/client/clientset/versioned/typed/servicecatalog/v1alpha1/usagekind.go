// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/mszostok/kyma/components/binding-usage-controller/pkg/apis/servicecatalog/v1alpha1"
	scheme "github.com/mszostok/kyma/components/binding-usage-controller/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// UsageKindsGetter has a method to return a UsageKindInterface.
// A group's client should implement this interface.
type UsageKindsGetter interface {
	UsageKinds() UsageKindInterface
}

// UsageKindInterface has methods to work with UsageKind resources.
type UsageKindInterface interface {
	Create(*v1alpha1.UsageKind) (*v1alpha1.UsageKind, error)
	Update(*v1alpha1.UsageKind) (*v1alpha1.UsageKind, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.UsageKind, error)
	List(opts v1.ListOptions) (*v1alpha1.UsageKindList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.UsageKind, err error)
	UsageKindExpansion
}

// usageKinds implements UsageKindInterface
type usageKinds struct {
	client rest.Interface
}

// newUsageKinds returns a UsageKinds
func newUsageKinds(c *ServicecatalogV1alpha1Client) *usageKinds {
	return &usageKinds{
		client: c.RESTClient(),
	}
}

// Get takes name of the usageKind, and returns the corresponding usageKind object, and an error if there is any.
func (c *usageKinds) Get(name string, options v1.GetOptions) (result *v1alpha1.UsageKind, err error) {
	result = &v1alpha1.UsageKind{}
	err = c.client.Get().
		Resource("usagekinds").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of UsageKinds that match those selectors.
func (c *usageKinds) List(opts v1.ListOptions) (result *v1alpha1.UsageKindList, err error) {
	result = &v1alpha1.UsageKindList{}
	err = c.client.Get().
		Resource("usagekinds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested usageKinds.
func (c *usageKinds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("usagekinds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a usageKind and creates it.  Returns the server's representation of the usageKind, and an error, if there is any.
func (c *usageKinds) Create(usageKind *v1alpha1.UsageKind) (result *v1alpha1.UsageKind, err error) {
	result = &v1alpha1.UsageKind{}
	err = c.client.Post().
		Resource("usagekinds").
		Body(usageKind).
		Do().
		Into(result)
	return
}

// Update takes the representation of a usageKind and updates it. Returns the server's representation of the usageKind, and an error, if there is any.
func (c *usageKinds) Update(usageKind *v1alpha1.UsageKind) (result *v1alpha1.UsageKind, err error) {
	result = &v1alpha1.UsageKind{}
	err = c.client.Put().
		Resource("usagekinds").
		Name(usageKind.Name).
		Body(usageKind).
		Do().
		Into(result)
	return
}

// Delete takes name of the usageKind and deletes it. Returns an error if one occurs.
func (c *usageKinds) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("usagekinds").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *usageKinds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("usagekinds").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched usageKind.
func (c *usageKinds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.UsageKind, err error) {
	result = &v1alpha1.UsageKind{}
	err = c.client.Patch(pt).
		Resource("usagekinds").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
