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

type GatewayV2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GatewayV2Parameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	InternalNetworkID *string `json:"internalNetworkId" tf:"internal_network_id,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	RouterID *string `json:"routerId" tf:"router_id,omitempty"`

	// +kubebuilder:validation:Required
	Spec *string `json:"spec" tf:"spec,omitempty"`

	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

// GatewayV2Spec defines the desired state of GatewayV2
type GatewayV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayV2Parameters `json:"forProvider"`
}

// GatewayV2Status defines the observed state of GatewayV2.
type GatewayV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayV2 is the Schema for the GatewayV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,otcjet}
type GatewayV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewayV2Spec   `json:"spec"`
	Status            GatewayV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayV2List contains a list of GatewayV2s
type GatewayV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GatewayV2 `json:"items"`
}

// Repository type metadata.
var (
	GatewayV2_Kind             = "GatewayV2"
	GatewayV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GatewayV2_Kind}.String()
	GatewayV2_KindAPIVersion   = GatewayV2_Kind + "." + CRDGroupVersion.String()
	GatewayV2_GroupVersionKind = CRDGroupVersion.WithKind(GatewayV2_Kind)
)

func init() {
	SchemeBuilder.Register(&GatewayV2{}, &GatewayV2List{})
}
