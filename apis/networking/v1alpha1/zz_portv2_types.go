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

type AllowedAddressPairsObservation struct {
}

type AllowedAddressPairsParameters struct {

	// +kubebuilder:validation:Required
	IPAddress *string `json:"ipAddress" tf:"ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`
}

type FixedIPObservation struct {
}

type FixedIPParameters struct {

	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// +kubebuilder:validation:Required
	SubnetID *string `json:"subnetId" tf:"subnet_id,omitempty"`
}

type PortV2Observation struct {
	AllFixedIps []*string `json:"allFixedIps,omitempty" tf:"all_fixed_ips,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PortV2Parameters struct {

	// +kubebuilder:validation:Optional
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// +kubebuilder:validation:Optional
	AllowedAddressPairs []AllowedAddressPairsParameters `json:"allowedAddressPairs,omitempty" tf:"allowed_address_pairs,omitempty"`

	// +kubebuilder:validation:Optional
	DeviceID *string `json:"deviceId,omitempty" tf:"device_id,omitempty"`

	// +kubebuilder:validation:Optional
	DeviceOwner *string `json:"deviceOwner,omitempty" tf:"device_owner,omitempty"`

	// +kubebuilder:validation:Optional
	FixedIP []FixedIPParameters `json:"fixedIp,omitempty" tf:"fixed_ip,omitempty"`

	// +kubebuilder:validation:Optional
	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`

	// +kubebuilder:validation:Required
	NetworkID *string `json:"networkId" tf:"network_id,omitempty"`

	// +kubebuilder:validation:Optional
	NoSecurityGroups *bool `json:"noSecurityGroups,omitempty" tf:"no_security_groups,omitempty"`

	// +kubebuilder:validation:Optional
	PortSecurityEnabled *bool `json:"portSecurityEnabled,omitempty" tf:"port_security_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// +kubebuilder:validation:Optional
	ValueSpecs map[string]string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

// PortV2Spec defines the desired state of PortV2
type PortV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PortV2Parameters `json:"forProvider"`
}

// PortV2Status defines the observed state of PortV2.
type PortV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PortV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PortV2 is the Schema for the PortV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,otcjet}
type PortV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PortV2Spec   `json:"spec"`
	Status            PortV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PortV2List contains a list of PortV2s
type PortV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PortV2 `json:"items"`
}

// Repository type metadata.
var (
	PortV2_Kind             = "PortV2"
	PortV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PortV2_Kind}.String()
	PortV2_KindAPIVersion   = PortV2_Kind + "." + CRDGroupVersion.String()
	PortV2_GroupVersionKind = CRDGroupVersion.WithKind(PortV2_Kind)
)

func init() {
	SchemeBuilder.Register(&PortV2{}, &PortV2List{})
}
