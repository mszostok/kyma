package storage_test

import (
	"testing"

	"github.com/coreos/etcd/clientv3"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"github.com/kyma-project/kyma/components/helm-broker/internal/storage"
	"github.com/kyma-project/kyma/components/helm-broker/internal/storage/driver/etcd"
	"github.com/kyma-project/kyma/components/helm-broker/internal/storage/driver/memory"
	"github.com/kyma-project/kyma/components/helm-broker/internal/storage/testdata"
)

func TestNewFactory(t *testing.T) {
	for s, tc := range map[string]struct {
		cfgGen               func() storage.ConfigList
		expBundle            interface{}
		expChart             interface{}
		expInstance          interface{}
		expInstanceOperation interface{}
	}{
		"MemorySingleAll":        {testdata.GoldenConfigMemorySingleAll, &memory.Bundle{}, &memory.Chart{}, &memory.Instance{}, &memory.InstanceOperation{}},
		"MemorySingleSeparate":   {testdata.GoldenConfigMemorySingleSeparate, &memory.Bundle{}, &memory.Chart{}, &memory.Instance{}, &memory.InstanceOperation{}},
		"MemoryMultipleSeparate": {testdata.GoldenConfigMemoryMultipleSeparate, &memory.Bundle{}, &memory.Chart{}, &memory.Instance{}, &memory.InstanceOperation{}},
		"EtcdSingleAll":          {testdata.GoldenConfigEtcdSingleAll, &etcd.Bundle{}, &etcd.Chart{}, &etcd.Instance{}, &etcd.InstanceOperation{}},
		"EtcdSingleSeparate":     {testdata.GoldenConfigEtcdSingleSeparate, &etcd.Bundle{}, &etcd.Chart{}, &etcd.Instance{}, &etcd.InstanceOperation{}},
		"EtcdMultipleSeparate":   {testdata.GoldenConfigEtcdMultipleSeparate, &etcd.Bundle{}, &etcd.Chart{}, &etcd.Instance{}, &etcd.InstanceOperation{}},
		"MixEMMESeparate":        {testdata.GoldenConfigMixEMMESeparate, &etcd.Bundle{}, &memory.Chart{}, &memory.Instance{}, &etcd.InstanceOperation{}},
		"MixMMEEGrouped":         {testdata.GoldenConfigMixMMEEGrouped, &memory.Bundle{}, &memory.Chart{}, &etcd.Instance{}, &etcd.InstanceOperation{}},
	} {
		t.Run(s, func(t *testing.T) {
			// GIVEN:
			cfg := tc.cfgGen()

			// This is unit test and we do not want to create a real client
			// to ETCD in `storage.NewFactory` in case of ETCD Driver
			// because we do not have any etcd server available
			fakeEtcdCli := clientv3.Client{}
			for i := range cfg {
				if cfg[i].Driver == storage.DriverEtcd {
					cfg[i].Etcd.ForceClient = &fakeEtcdCli
				}
			}

			// WHEN:
			got, err := storage.NewFactory(&cfg)

			// THEN:
			require.NoError(t, err)
			assert.NotNil(t, got)
			assert.IsType(t, tc.expBundle, got.Bundle())
			assert.IsType(t, tc.expChart, got.Chart())
			assert.IsType(t, tc.expInstance, got.Instance())
			assert.IsType(t, tc.expInstanceOperation, got.InstanceOperation())
		})
	}
}
