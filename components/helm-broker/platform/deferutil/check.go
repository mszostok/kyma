package deferutil

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// CheckFn helps
// In the future it can be improved and hardcoded `logrus.Errorf` can be replace with dynamic ErrorHandlers, similar to
// this implementation: https://github.com/kubernetes/apimachinery/blob/release-1.13/pkg/util/runtime/runtime.go#L101-L114
func CheckFn(f func() error, msg string, args ...interface{}) {
	if err := f(); err != nil {
		logrus.Errorf("%s: %v", fmt.Sprintf(msg, args...), err)
	}
}
