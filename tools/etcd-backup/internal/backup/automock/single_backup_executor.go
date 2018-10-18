// Code generated by mockery v1.0.0
package automock

import backup "github.com/mszostok/kyma/tools/etcd-backup/internal/backup"
import mock "github.com/stretchr/testify/mock"

// singleBackupExecutor is an autogenerated mock type for the singleBackupExecutor type
type singleBackupExecutor struct {
	mock.Mock
}

// SingleBackup provides a mock function with given fields: stopCh, blobPrefix
func (_m *singleBackupExecutor) SingleBackup(stopCh <-chan struct{}, blobPrefix string) (*backup.SingleBackupOutput, error) {
	ret := _m.Called(stopCh, blobPrefix)

	var r0 *backup.SingleBackupOutput
	if rf, ok := ret.Get(0).(func(<-chan struct{}, string) *backup.SingleBackupOutput); ok {
		r0 = rf(stopCh, blobPrefix)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backup.SingleBackupOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(<-chan struct{}, string) error); ok {
		r1 = rf(stopCh, blobPrefix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
