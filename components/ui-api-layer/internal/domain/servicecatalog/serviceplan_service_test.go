package servicecatalog_test

import (
	"testing"
	"time"

	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset/fake"
	"github.com/kubernetes-incubator/service-catalog/pkg/client/informers_generated/externalversions"
	"github.com/mszostok/kyma/components/ui-api-layer/internal/domain/servicecatalog"
	testingUtils "github.com/mszostok/kyma/components/ui-api-layer/internal/testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestServicePlanService_Find(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		environmentName := "env"
		planName := "testExample"
		servicePlan := fixServicePlan(planName, "test", planName, environmentName)
		client := fake.NewSimpleClientset(servicePlan)

		informerFactory := externalversions.NewSharedInformerFactory(client, 0)
		servicePlanInformer := informerFactory.Servicecatalog().V1beta1().ServicePlans().Informer()

		svc := servicecatalog.NewServicePlanService(servicePlanInformer)

		testingUtils.WaitForInformerStartAtMost(t, time.Second, servicePlanInformer)

		plan, err := svc.Find(planName, environmentName)
		require.NoError(t, err)
		assert.Equal(t, servicePlan, plan)
	})

	t.Run("NotFound", func(t *testing.T) {
		client := fake.NewSimpleClientset()

		informerFactory := externalversions.NewSharedInformerFactory(client, 0)
		servicePlanInformer := informerFactory.Servicecatalog().V1beta1().ServicePlans().Informer()

		svc := servicecatalog.NewServicePlanService(servicePlanInformer)

		testingUtils.WaitForInformerStartAtMost(t, time.Second, servicePlanInformer)

		plan, err := svc.Find("doesntExist", "env")
		require.NoError(t, err)
		assert.Nil(t, plan)
	})
}

func TestServicePlanService_FindByExternalNameForClass(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		environmentName := "env"
		className := "test"
		planName := "testExample"
		externalName := "testExternal"
		servicePlan := fixServicePlan(planName, className, externalName, environmentName)
		client := fake.NewSimpleClientset(servicePlan)

		informerFactory := externalversions.NewSharedInformerFactory(client, 0)
		servicePlanInformer := informerFactory.Servicecatalog().V1beta1().ServicePlans().Informer()

		svc := servicecatalog.NewServicePlanService(servicePlanInformer)

		testingUtils.WaitForInformerStartAtMost(t, time.Second, servicePlanInformer)

		plan, err := svc.FindByExternalName(externalName, className, environmentName)
		require.NoError(t, err)
		assert.Equal(t, servicePlan, plan)
	})

	t.Run("NotFound", func(t *testing.T) {
		client := fake.NewSimpleClientset()

		informerFactory := externalversions.NewSharedInformerFactory(client, 0)
		servicePlanInformer := informerFactory.Servicecatalog().V1beta1().ServicePlans().Informer()

		svc := servicecatalog.NewServicePlanService(servicePlanInformer)

		testingUtils.WaitForInformerStartAtMost(t, time.Second, servicePlanInformer)

		plan, err := svc.FindByExternalName("doesntExist", "none", "env")

		require.NoError(t, err)
		assert.Nil(t, plan)
	})

	t.Run("Error", func(t *testing.T) {
		environmentName := "env"
		className := "duplicateName"
		externalName := "duplicateName"
		client := fake.NewSimpleClientset(
			fixServicePlan("1", className, externalName, environmentName),
			fixServicePlan("2", className, externalName, environmentName),
			fixServicePlan("3", className, externalName, environmentName),
		)

		informerFactory := externalversions.NewSharedInformerFactory(client, 0)
		servicePlanInformer := informerFactory.Servicecatalog().V1beta1().ServicePlans().Informer()

		svc := servicecatalog.NewServicePlanService(servicePlanInformer)

		testingUtils.WaitForInformerStartAtMost(t, time.Second, servicePlanInformer)

		_, err := svc.FindByExternalName(externalName, className, environmentName)

		assert.Error(t, err)
	})
}

func TestServicePlanService_ListForClass(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		environmentName := "env"
		className := "testClassName"

		servicePlan1 := fixServicePlan("1", className, "1", environmentName)
		servicePlan2 := fixServicePlan("2", className, "2", environmentName)
		servicePlan3 := fixServicePlan("3", className, "3", environmentName)
		client := fake.NewSimpleClientset(servicePlan1, servicePlan2, servicePlan3)

		informerFactory := externalversions.NewSharedInformerFactory(client, 0)
		servicePlanInformer := informerFactory.Servicecatalog().V1beta1().ServicePlans().Informer()

		svc := servicecatalog.NewServicePlanService(servicePlanInformer)

		testingUtils.WaitForInformerStartAtMost(t, time.Second, servicePlanInformer)

		plans, err := svc.ListForServiceClass(className, environmentName)
		require.NoError(t, err)
		assert.Equal(t, []*v1beta1.ServicePlan{
			servicePlan1, servicePlan2, servicePlan3,
		}, plans)
	})

	t.Run("NotFound", func(t *testing.T) {
		client := fake.NewSimpleClientset()

		informerFactory := externalversions.NewSharedInformerFactory(client, 0)
		servicePlanInformer := informerFactory.Servicecatalog().V1beta1().ServicePlans().Informer()
		svc := servicecatalog.NewServicePlanService(servicePlanInformer)

		testingUtils.WaitForInformerStartAtMost(t, time.Second, servicePlanInformer)

		var emptyArray []*v1beta1.ServicePlan
		plans, err := svc.ListForServiceClass("doesntExist", "env")
		require.NoError(t, err)
		assert.Equal(t, emptyArray, plans)
	})

}

func fixServicePlan(name, relatedServiceClassName, externalName, environment string) *v1beta1.ServicePlan {
	plan := v1beta1.ServicePlan{
		Spec: v1beta1.ServicePlanSpec{
			CommonServicePlanSpec: v1beta1.CommonServicePlanSpec{
				ExternalName: externalName,
			},
			ServiceClassRef: v1beta1.LocalObjectReference{
				Name: relatedServiceClassName,
			},
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: environment,
		},
	}

	return &plan
}
