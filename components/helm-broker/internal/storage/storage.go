package storage

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"

	"github.com/kyma-project/kyma/components/helm-broker/internal/storage/driver/etcd"
	"github.com/kyma-project/kyma/components/helm-broker/internal/storage/driver/memory"
)

// Factory provides access to concrete storage.
// Multiple calls should to specific storage return the same storage instance.
type Factory interface {
	Bundle() Bundle
	Chart() Chart
	Instance() Instance
	InstanceOperation() InstanceOperation
	InstanceBindData() InstanceBindData
}

// DriverType defines type of data storage
type DriverType string

const (
	// DriverEtcd is a driver for key-value store - Etcd
	DriverEtcd DriverType = "etcd"
	// DriverMemory is a driver to local in-memory store
	DriverMemory DriverType = "memory"
)

// EntityName defines name of the entity in database
type EntityName string

const (
	// EntityAll represents name of all entities
	EntityAll EntityName = "all"
	// EntityChart represents name of chart entities
	EntityChart EntityName = "chart"
	// EntityBundle represents name of bundle entities
	EntityBundle EntityName = "bundle"
	// EntityInstance represents name of services instances entities
	EntityInstance EntityName = "instance"
	// EntityInstanceOperation represents name of instances operations entities
	EntityInstanceOperation EntityName = "instanceOperation"
	// EntityInstanceBindData represents name of bind data entities
	EntityInstanceBindData EntityName = "entityInstanceBindData"
)

// ProviderConfig provides configuration to the database provider
type ProviderConfig struct{}

// ProviderConfigMap contains map of provided configurations for given entities
type ProviderConfigMap map[EntityName]ProviderConfig

// Config contains database configuration.
type Config struct {
	Driver  DriverType        `json:"driver" valid:"required"`
	Provide ProviderConfigMap `json:"provide" valid:"required"`
	Etcd    etcd.Config       `json:"etcd"`
	Memory  memory.Config     `json:"memory"`
}

// ConfigList is a list of configurations
type ConfigList []Config

// ConfigParse is parsing yaml file to the ConfigList
func ConfigParse(inByte []byte) (*ConfigList, error) {
	var cl ConfigList

	if err := yaml.Unmarshal(inByte, &cl); err != nil {
		return nil, errors.Wrap(err, "while unmarshalling yaml")
	}

	return &cl, nil
}

// NewConfigListAllMemory returns configured configList with the memory driver for all entities.
func NewConfigListAllMemory() *ConfigList {
	return &ConfigList{{Driver: DriverMemory, Provide: ProviderConfigMap{EntityAll: ProviderConfig{}}}}
}

// NewFactory is a factory for entities based on given ConfigList
// TODO: add error handling
func NewFactory(cl *ConfigList) (Factory, error) {
	fact := concreteFactory{}

	for _, cfg := range *cl {

		var (
			bundleFact            func() Bundle
			chartFact             func() Chart
			instanceFact          func() Instance
			instanceOperationFact func() InstanceOperation
			instanceBindDataFact  func() InstanceBindData
		)

		switch cfg.Driver {
		case DriverMemory:
			bundleFact = func() Bundle {
				return memory.NewBundle()
			}
			chartFact = func() Chart {
				return memory.NewChart()
			}
			instanceFact = func() Instance {
				return memory.NewInstance()
			}
			instanceOperationFact = func() InstanceOperation {
				return memory.NewInstanceOperation()
			}
			instanceBindDataFact = func() InstanceBindData {
				return memory.NewInstanceBindData()
			}
		case DriverEtcd:
			var cli etcd.Client
			if cfg.Etcd.ForceClient != nil {
				cli = cfg.Etcd.ForceClient
			} else {
				var err error
				cli, err = etcd.NewClient(cfg.Etcd)
				if err != nil {
					return nil, errors.Wrap(err, "while creating new ETCD client")
				}
			}

			bundleFact = func() Bundle {
				return etcd.NewBundle(cli)
			}
			chartFact = func() Chart {
				return etcd.NewChart(cli)
			}
			instanceFact = func() Instance {
				return etcd.NewInstance(cli)
			}
			instanceOperationFact = func() InstanceOperation {
				return etcd.NewInstanceOperation(cli)
			}
			instanceBindDataFact = func() InstanceBindData {
				return etcd.NewInstanceBindData(cli)
			}
		default:
			return nil, errors.New("unknown driver type")
		}

		for em := range cfg.Provide {
			switch em {
			case EntityChart:
				fact.chart = chartFact()
			case EntityBundle:
				fact.bundle = bundleFact()
			case EntityInstance:
				fact.instance = instanceFact()
			case EntityInstanceOperation:
				fact.instanceOperation = instanceOperationFact()
			case EntityInstanceBindData:
				fact.instanceBindData = instanceBindDataFact()
			case EntityAll:
				fact.chart = chartFact()
				fact.bundle = bundleFact()
				fact.instance = instanceFact()
				fact.instanceOperation = instanceOperationFact()
				fact.instanceBindData = instanceBindDataFact()
			default:
			}
		}
	}

	return &fact, nil
}

type concreteFactory struct {
	bundle            Bundle
	chart             Chart
	instance          Instance
	instanceOperation InstanceOperation
	instanceBindData  InstanceBindData
}

func (f *concreteFactory) Bundle() Bundle {
	return f.bundle
}
func (f *concreteFactory) Chart() Chart {
	return f.chart
}
func (f *concreteFactory) Instance() Instance {
	return f.instance
}
func (f *concreteFactory) InstanceOperation() InstanceOperation {
	return f.instanceOperation
}
func (f *concreteFactory) InstanceBindData() InstanceBindData {
	return f.instanceBindData
}
