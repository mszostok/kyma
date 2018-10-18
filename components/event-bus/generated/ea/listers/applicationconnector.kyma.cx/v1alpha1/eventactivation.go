// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/mszostok/kyma/components/event-bus/internal/ea/apis/applicationconnector.kyma.cx/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EventActivationLister helps list EventActivations.
type EventActivationLister interface {
	// List lists all EventActivations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.EventActivation, err error)
	// EventActivations returns an object that can list and get EventActivations.
	EventActivations(namespace string) EventActivationNamespaceLister
	EventActivationListerExpansion
}

// eventActivationLister implements the EventActivationLister interface.
type eventActivationLister struct {
	indexer cache.Indexer
}

// NewEventActivationLister returns a new EventActivationLister.
func NewEventActivationLister(indexer cache.Indexer) EventActivationLister {
	return &eventActivationLister{indexer: indexer}
}

// List lists all EventActivations in the indexer.
func (s *eventActivationLister) List(selector labels.Selector) (ret []*v1alpha1.EventActivation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EventActivation))
	})
	return ret, err
}

// EventActivations returns an object that can list and get EventActivations.
func (s *eventActivationLister) EventActivations(namespace string) EventActivationNamespaceLister {
	return eventActivationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EventActivationNamespaceLister helps list and get EventActivations.
type EventActivationNamespaceLister interface {
	// List lists all EventActivations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.EventActivation, err error)
	// Get retrieves the EventActivation from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.EventActivation, error)
	EventActivationNamespaceListerExpansion
}

// eventActivationNamespaceLister implements the EventActivationNamespaceLister
// interface.
type eventActivationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EventActivations in the indexer for a given namespace.
func (s eventActivationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EventActivation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EventActivation))
	})
	return ret, err
}

// Get retrieves the EventActivation from the indexer for a given namespace and name.
func (s eventActivationNamespaceLister) Get(name string) (*v1alpha1.EventActivation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("eventactivation"), name)
	}
	return obj.(*v1alpha1.EventActivation), nil
}
