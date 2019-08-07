package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// +kubebuilder:object:root=true

// VirtualService is the Schema for the addonsconfigurations API
// Important: Run "make generates" to regenerate files after modifying this struct
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type VirtualService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualServiceSpec   `json:"spec,omitempty"`
	//Status AddonsConfigurationStatus `json:"status,omitempty"`
}

type VirtualServiceSpec struct {

	Hosts []string `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`

	Gateways []string `protobuf:"bytes,2,rep,name=gateways,proto3" json:"gateways,omitempty"`

	ExportTo             []string `protobuf:"bytes,6,rep,name=export_to,json=exportTo,proto3" json:"export_to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
// VirtualServiceList contains a list of VirtualService
// Important: Run "make generates" to regenerate files after modifying this struct
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type VirtualServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualService `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&VirtualService{}, &VirtualServiceList{},
	)
}
