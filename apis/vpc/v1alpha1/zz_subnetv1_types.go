/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SubnetV1Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type SubnetV1Parameters struct {

	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// +kubebuilder:validation:Required
	Cidr *string `json:"cidr" tf:"cidr,omitempty"`

	// +kubebuilder:validation:Optional
	DHCPEnable *bool `json:"dhcpEnable,omitempty" tf:"dhcp_enable,omitempty"`

	// +kubebuilder:validation:Optional
	DNSList []*string `json:"dnsList,omitempty" tf:"dns_list,omitempty"`

	// +kubebuilder:validation:Required
	GatewayIP *string `json:"gatewayIp" tf:"gateway_ip,omitempty"`

	// +kubebuilder:validation:Optional
	NtpAddresses *string `json:"ntpAddresses,omitempty" tf:"ntp_addresses,omitempty"`

	// +kubebuilder:validation:Optional
	PrimaryDNS *string `json:"primaryDns,omitempty" tf:"primary_dns,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	SecondaryDNS *string `json:"secondaryDns,omitempty" tf:"secondary_dns,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	VPCID *string `json:"vpcId" tf:"vpc_id,omitempty"`
}

// SubnetV1Spec defines the desired state of SubnetV1
type SubnetV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetV1Parameters `json:"forProvider"`
}

// SubnetV1Status defines the observed state of SubnetV1.
type SubnetV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetV1 is the Schema for the SubnetV1s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,otcjet}
type SubnetV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetV1Spec   `json:"spec"`
	Status            SubnetV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetV1List contains a list of SubnetV1s
type SubnetV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubnetV1 `json:"items"`
}

// Repository type metadata.
var (
	SubnetV1_Kind             = "SubnetV1"
	SubnetV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubnetV1_Kind}.String()
	SubnetV1_KindAPIVersion   = SubnetV1_Kind + "." + CRDGroupVersion.String()
	SubnetV1_GroupVersionKind = CRDGroupVersion.WithKind(SubnetV1_Kind)
)

func init() {
	SchemeBuilder.Register(&SubnetV1{}, &SubnetV1List{})
}
