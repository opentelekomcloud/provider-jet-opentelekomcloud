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

type ProtectiongroupV1Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ProtectiongroupV1Parameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	DomainID *string `json:"domainId" tf:"domain_id,omitempty"`

	// +kubebuilder:validation:Optional
	DrType *string `json:"drType,omitempty" tf:"dr_type,omitempty"`

	// +kubebuilder:validation:Required
	SourceAvailabilityZone *string `json:"sourceAvailabilityZone" tf:"source_availability_zone,omitempty"`

	// +kubebuilder:validation:Required
	SourceVPCID *string `json:"sourceVpcId" tf:"source_vpc_id,omitempty"`

	// +kubebuilder:validation:Required
	TargetAvailabilityZone *string `json:"targetAvailabilityZone" tf:"target_availability_zone,omitempty"`
}

// ProtectiongroupV1Spec defines the desired state of ProtectiongroupV1
type ProtectiongroupV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProtectiongroupV1Parameters `json:"forProvider"`
}

// ProtectiongroupV1Status defines the observed state of ProtectiongroupV1.
type ProtectiongroupV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProtectiongroupV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProtectiongroupV1 is the Schema for the ProtectiongroupV1s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,otcjet}
type ProtectiongroupV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProtectiongroupV1Spec   `json:"spec"`
	Status            ProtectiongroupV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProtectiongroupV1List contains a list of ProtectiongroupV1s
type ProtectiongroupV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProtectiongroupV1 `json:"items"`
}

// Repository type metadata.
var (
	ProtectiongroupV1_Kind             = "ProtectiongroupV1"
	ProtectiongroupV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProtectiongroupV1_Kind}.String()
	ProtectiongroupV1_KindAPIVersion   = ProtectiongroupV1_Kind + "." + CRDGroupVersion.String()
	ProtectiongroupV1_GroupVersionKind = CRDGroupVersion.WithKind(ProtectiongroupV1_Kind)
)

func init() {
	SchemeBuilder.Register(&ProtectiongroupV1{}, &ProtectiongroupV1List{})
}