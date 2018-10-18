// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	scheme "github.com/mszostok/kyma/components/event-bus/generated/ea/clientset/versioned/scheme"
	v1alpha1 "github.com/mszostok/kyma/components/event-bus/internal/ea/apis/applicationconnector.kyma.cx/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EventActivationsGetter has a method to return a EventActivationInterface.
// A group's client should implement this interface.
type EventActivationsGetter interface {
	EventActivations(namespace string) EventActivationInterface
}

// EventActivationInterface has methods to work with EventActivation resources.
type EventActivationInterface interface {
	Create(*v1alpha1.EventActivation) (*v1alpha1.EventActivation, error)
	Update(*v1alpha1.EventActivation) (*v1alpha1.EventActivation, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.EventActivation, error)
	List(opts v1.ListOptions) (*v1alpha1.EventActivationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EventActivation, err error)
	EventActivationExpansion
}

// eventActivations implements EventActivationInterface
type eventActivations struct {
	client rest.Interface
	ns     string
}

// newEventActivations returns a EventActivations
func newEventActivations(c *ApplicationconnectorV1alpha1Client, namespace string) *eventActivations {
	return &eventActivations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the eventActivation, and returns the corresponding eventActivation object, and an error if there is any.
func (c *eventActivations) Get(name string, options v1.GetOptions) (result *v1alpha1.EventActivation, err error) {
	result = &v1alpha1.EventActivation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eventactivations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EventActivations that match those selectors.
func (c *eventActivations) List(opts v1.ListOptions) (result *v1alpha1.EventActivationList, err error) {
	result = &v1alpha1.EventActivationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eventactivations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested eventActivations.
func (c *eventActivations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("eventactivations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a eventActivation and creates it.  Returns the server's representation of the eventActivation, and an error, if there is any.
func (c *eventActivations) Create(eventActivation *v1alpha1.EventActivation) (result *v1alpha1.EventActivation, err error) {
	result = &v1alpha1.EventActivation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("eventactivations").
		Body(eventActivation).
		Do().
		Into(result)
	return
}

// Update takes the representation of a eventActivation and updates it. Returns the server's representation of the eventActivation, and an error, if there is any.
func (c *eventActivations) Update(eventActivation *v1alpha1.EventActivation) (result *v1alpha1.EventActivation, err error) {
	result = &v1alpha1.EventActivation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("eventactivations").
		Name(eventActivation.Name).
		Body(eventActivation).
		Do().
		Into(result)
	return
}

// Delete takes name of the eventActivation and deletes it. Returns an error if one occurs.
func (c *eventActivations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eventactivations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *eventActivations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eventactivations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched eventActivation.
func (c *eventActivations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EventActivation, err error) {
	result = &v1alpha1.EventActivation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("eventactivations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
