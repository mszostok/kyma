package resource

import (
	"log"

	"github.com/avast/retry-go"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"

	"time"
)

type Manager struct {
	RetryOptions []retry.Option
}

//CreateResource creates a given k8s resource
func (m *Manager) CreateResource(client dynamic.Interface, resourceSchema schema.GroupVersionResource, namespace string, manifest unstructured.Unstructured) {
	panicOnErr(retry.Do(func() error {
		if _, err := client.Resource(resourceSchema).Namespace(namespace).Create(&manifest, metav1.CreateOptions{}); err != nil {
			log.Printf("Error: %+v", err)
			return err
		}
		return nil
	}, m.RetryOptions...))
}

//UpdateResource updates a given k8s resource
func (m *Manager) UpdateResource(client dynamic.Interface, resourceSchema schema.GroupVersionResource, namespace string, name string, updateTo unstructured.Unstructured) {
	panicOnErr(retry.Do(func() error {
		time.Sleep(5 * time.Second) //TODO: delete after waiting for resource creation is implemented
		toUpdate, err := client.Resource(resourceSchema).Namespace(namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		updateTo.SetResourceVersion(toUpdate.GetResourceVersion())
		_, err = client.Resource(resourceSchema).Namespace(namespace).Update(&updateTo, metav1.UpdateOptions{})
		if err != nil {
			return err
		}

		return nil
	}, m.RetryOptions...))
}

//DeleteResource deletes a given k8s resource
func (m *Manager) DeleteResource(client dynamic.Interface, resourceSchema schema.GroupVersionResource, namespace string, resourceName string) {
	panicOnErr(retry.Do(func() error {
		deletePolicy := metav1.DeletePropagationForeground
		deleteOptions := &metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}
		if err := client.Resource(resourceSchema).Namespace(namespace).Delete(resourceName, deleteOptions); err != nil {
			if !apierrors.IsNotFound(err) {
				return err
			}
		}
		return nil
	}, m.RetryOptions...))
}

//GetResource returns chosed k8s object
func (m *Manager) GetResource(client dynamic.Interface, resourceSchema schema.GroupVersionResource, namespace string, resourceName string) {
	panicOnErr(retry.Do(func() error {
		resource, err := client.Resource(resourceSchema).Namespace(namespace).Get(resourceName, metav1.GetOptions{})
		if err != nil {
			log.Printf("Error: %+v", err)
			return err
		}
		log.Printf("Resource found: %+v", resource.GetName())
		return nil
	}, m.RetryOptions...))
}

func panicOnErr(err error) {
	if err != nil {
		log.Panicf("Error: %+v", err)
	}
}
