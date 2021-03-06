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

type NetworkV2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type NetworkV2Parameters struct {

	// +kubebuilder:validation:Optional
	AdminStateUp *string `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	Segments []SegmentsParameters `json:"segments,omitempty" tf:"segments,omitempty"`

	// +kubebuilder:validation:Optional
	Shared *string `json:"shared,omitempty" tf:"shared,omitempty"`

	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// +kubebuilder:validation:Optional
	ValueSpecs map[string]string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

type SegmentsObservation struct {
}

type SegmentsParameters struct {

	// +kubebuilder:validation:Optional
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// +kubebuilder:validation:Optional
	PhysicalNetwork *string `json:"physicalNetwork,omitempty" tf:"physical_network,omitempty"`

	// +kubebuilder:validation:Optional
	SegmentationID *int64 `json:"segmentationId,omitempty" tf:"segmentation_id,omitempty"`
}

// NetworkV2Spec defines the desired state of NetworkV2
type NetworkV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkV2Parameters `json:"forProvider"`
}

// NetworkV2Status defines the observed state of NetworkV2.
type NetworkV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkV2 is the Schema for the NetworkV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,otcjet}
type NetworkV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkV2Spec   `json:"spec"`
	Status            NetworkV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkV2List contains a list of NetworkV2s
type NetworkV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkV2 `json:"items"`
}

// Repository type metadata.
var (
	NetworkV2_Kind             = "NetworkV2"
	NetworkV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkV2_Kind}.String()
	NetworkV2_KindAPIVersion   = NetworkV2_Kind + "." + CRDGroupVersion.String()
	NetworkV2_GroupVersionKind = CRDGroupVersion.WithKind(NetworkV2_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkV2{}, &NetworkV2List{})
}
