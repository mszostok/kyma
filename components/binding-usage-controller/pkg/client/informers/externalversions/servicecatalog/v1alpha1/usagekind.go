// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	servicecatalog_v1alpha1 "github.com/mszostok/kyma/components/binding-usage-controller/pkg/apis/servicecatalog/v1alpha1"
	versioned "github.com/mszostok/kyma/components/binding-usage-controller/pkg/client/clientset/versioned"
	internalinterfaces "github.com/mszostok/kyma/components/binding-usage-controller/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/mszostok/kyma/components/binding-usage-controller/pkg/client/listers/servicecatalog/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// UsageKindInformer provides access to a shared informer and lister for
// UsageKinds.
type UsageKindInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.UsageKindLister
}

type usageKindInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewUsageKindInformer constructs a new informer for UsageKind type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewUsageKindInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredUsageKindInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredUsageKindInformer constructs a new informer for UsageKind type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredUsageKindInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServicecatalogV1alpha1().UsageKinds().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServicecatalogV1alpha1().UsageKinds().Watch(options)
			},
		},
		&servicecatalog_v1alpha1.UsageKind{},
		resyncPeriod,
		indexers,
	)
}

func (f *usageKindInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredUsageKindInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *usageKindInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&servicecatalog_v1alpha1.UsageKind{}, f.defaultInformer)
}

func (f *usageKindInformer) Lister() v1alpha1.UsageKindLister {
	return v1alpha1.NewUsageKindLister(f.Informer().GetIndexer())
}
