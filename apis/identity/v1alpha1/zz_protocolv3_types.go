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

type MetadataObservation struct {
}

type MetadataParameters struct {

	// +kubebuilder:validation:Required
	DomainID *string `json:"domainId" tf:"domain_id,omitempty"`

	// +kubebuilder:validation:Required
	Metadata *string `json:"metadata" tf:"metadata,omitempty"`

	// +kubebuilder:validation:Optional
	XaccountType *string `json:"xaccountType,omitempty" tf:"xaccount_type,omitempty"`
}

type ProtocolV3Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Links map[string]string `json:"links,omitempty" tf:"links,omitempty"`
}

type ProtocolV3Parameters struct {

	// +kubebuilder:validation:Required
	MappingID *string `json:"mappingId" tf:"mapping_id,omitempty"`

	// +kubebuilder:validation:Optional
	Metadata []MetadataParameters `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Required
	ProviderID *string `json:"providerId" tf:"provider_id,omitempty"`
}

// ProtocolV3Spec defines the desired state of ProtocolV3
type ProtocolV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProtocolV3Parameters `json:"forProvider"`
}

// ProtocolV3Status defines the observed state of ProtocolV3.
type ProtocolV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProtocolV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProtocolV3 is the Schema for the ProtocolV3s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,otcjet}
type ProtocolV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProtocolV3Spec   `json:"spec"`
	Status            ProtocolV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProtocolV3List contains a list of ProtocolV3s
type ProtocolV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProtocolV3 `json:"items"`
}

// Repository type metadata.
var (
	ProtocolV3_Kind             = "ProtocolV3"
	ProtocolV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProtocolV3_Kind}.String()
	ProtocolV3_KindAPIVersion   = ProtocolV3_Kind + "." + CRDGroupVersion.String()
	ProtocolV3_GroupVersionKind = CRDGroupVersion.WithKind(ProtocolV3_Kind)
)

func init() {
	SchemeBuilder.Register(&ProtocolV3{}, &ProtocolV3List{})
}
