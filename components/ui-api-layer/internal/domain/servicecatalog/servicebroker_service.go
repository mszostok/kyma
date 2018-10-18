package servicecatalog

import (
	"fmt"

	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"github.com/mszostok/kyma/components/ui-api-layer/internal/pager"
	"github.com/mszostok/kyma/components/ui-api-layer/pkg/resource"
	"k8s.io/client-go/tools/cache"
)

type serviceBrokerService struct {
	informer cache.SharedIndexInformer
	notifier notifier
}

func newServiceBrokerService(informer cache.SharedIndexInformer) *serviceBrokerService {
	return &serviceBrokerService{
		informer: informer,
		notifier: resource.NewNotifier(),
	}
}

func (svc *serviceBrokerService) Find(name, environment string) (*v1beta1.ServiceBroker, error) {
	key := fmt.Sprintf("%s/%s", environment, name)
	item, exists, err := svc.informer.GetStore().GetByKey(key)
	if err != nil || !exists {
		return nil, err
	}

	serviceBroker, ok := item.(*v1beta1.ServiceBroker)
	if !ok {
		return nil, fmt.Errorf("Incorrect item type: %T, should be: *ServiceBroker", item)
	}

	return serviceBroker, nil
}

func (svc *serviceBrokerService) List(environment string, pagingParams pager.PagingParams) ([]*v1beta1.ServiceBroker, error) {
	items, err := pager.FromIndexer(svc.informer.GetIndexer(), "namespace", environment).Limit(pagingParams)
	if err != nil {
		return nil, err
	}

	var serviceBrokers []*v1beta1.ServiceBroker
	for _, item := range items {
		serviceBroker, ok := item.(*v1beta1.ServiceBroker)
		if !ok {
			return nil, fmt.Errorf("Incorrect item type: %T, should be: *ServiceBroker", item)
		}
		serviceBrokers = append(serviceBrokers, serviceBroker)
	}

	return serviceBrokers, nil
}

func (svc *serviceBrokerService) Subscribe(listener resource.Listener) {
	svc.notifier.AddListener(listener)
}

func (svc *serviceBrokerService) Unsubscribe(listener resource.Listener) {
	svc.notifier.DeleteListener(listener)
}
