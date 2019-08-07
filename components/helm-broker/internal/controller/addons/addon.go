package addons

import (
	"strings"

	"github.com/kyma-project/kyma/components/helm-broker/internal"
	"k8s.io/helm/pkg/proto/hapi/chart"
)

// AddonController is a wraper for Addon element with extra fields like URL, bunlde or charts
type AddonController struct {
	ID            string
	URL           string
	CompleteAddon *internal.Addon
	Charts        []*chart.Chart
}

// NewAddon returns pointer to new AddonController based on name, version and url
func NewAddon(n, v, u string) *AddonController {
	return &AddonController{
		URL: u,
	}
}

// IsReady informs addon is in ready status
func (a *AddonController) IsReady() bool {
	return false
}

// IsComplete informs AddonController has no fetching/loading error, what means own ID (from addon)
func (a *AddonController) IsComplete() bool {
	return a.ID != ""
}

// FetchingError sets addons as failed, sets addon reason as FetchingError
func (a *AddonController) FetchingError(err error) {
	a.failed()
	//a.setAddonFailedInfo(v1alpha3.AddonFetchingError, a.limitMessage(err.Error()))
}

// LoadingError sets addons as failed, sets addon reason as LoadingError
func (a *AddonController) LoadingError(err error) {
	a.failed()
	//a.setAddonFailedInfo(v1alpha3.AddonLoadingError, err.Error())
}

// ConflictInSpecifiedRepositories sets addons as failed, sets addon reason as ConflictInSpecifiedRepositories
func (a *AddonController) ConflictInSpecifiedRepositories(err error) {
	a.failed()
	//a.setAddonFailedInfo(v1alpha3.AddonConflictInSpecifiedRepositories, err.Error())
}

// ConflictWithAlreadyRegisteredAddons sets addons as failed, sets addon reason as ConflictWithAlreadyRegisteredAddons
func (a *AddonController) ConflictWithAlreadyRegisteredAddons(err error) {
	a.failed()
	//a.setAddonFailedInfo(v1alpha3.AddonConflictWithAlreadyRegisteredAddons, err.Error())
}

// RegisteringError sets addons as failed, sets addon reason as RegisteringError
func (a *AddonController) RegisteringError(err error) {
	a.failed()
	//a.setAddonFailedInfo(v1alpha3.AddonRegisteringError, err.Error())
}

func (a *AddonController) failed() {
	//a.Addon.Status = v1alpha3.AddonStatusFailed
}

//func (a *AddonController) setAddonFailedInfo(reason v1alpha3.AddonStatusReason, message string) {
//	a.Addon.Reason = reason
//	a.Addon.Message = fmt.Sprintf(reason.Message(), message)
//}

// limitMessage limits content of message field for AddonConfiguration which e.g. for fetching error
// could be very long. Full message occurs in controller log
func (a *AddonController) limitMessage(content string) string {
	parts := strings.Split(content, ":")
	if len(parts) <= 4 {
		return content
	}

	return strings.Join(parts[:4], ":")
}
