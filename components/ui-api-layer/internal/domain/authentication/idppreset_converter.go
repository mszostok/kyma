package authentication

import (
	"github.com/mszostok/kyma/components/idppreset/pkg/apis/authentication/v1alpha1"
	"github.com/mszostok/kyma/components/ui-api-layer/internal/gqlschema"
)

type idpPresetConverter struct{}

func (c *idpPresetConverter) ToGQL(in *v1alpha1.IDPPreset) gqlschema.IDPPreset {
	if in == nil {
		return gqlschema.IDPPreset{}
	}

	return gqlschema.IDPPreset{
		Name:    in.Name,
		Issuer:  in.Spec.Issuer,
		JwksURI: in.Spec.JwksUri,
	}
}
