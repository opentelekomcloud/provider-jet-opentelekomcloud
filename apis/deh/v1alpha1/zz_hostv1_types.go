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

type AvailableInstanceCapacitiesObservation struct {
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`
}

type AvailableInstanceCapacitiesParameters struct {
}

type HostV1Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type HostV1Parameters struct {

	// +kubebuilder:validation:Optional
	AutoPlacement *string `json:"autoPlacement,omitempty" tf:"auto_placement,omitempty"`

	// +kubebuilder:validation:Required
	AvailabilityZone *string `json:"availabilityZone" tf:"availability_zone,omitempty"`

	// +kubebuilder:validation:Optional
	AvailableInstanceCapacities []AvailableInstanceCapacitiesParameters `json:"availableInstanceCapacities,omitempty" tf:"available_instance_capacities,omitempty"`

	// +kubebuilder:validation:Optional
	AvailableMemory *int64 `json:"availableMemory,omitempty" tf:"available_memory,omitempty"`

	// +kubebuilder:validation:Optional
	AvailableVcpus *int64 `json:"availableVcpus,omitempty" tf:"available_vcpus,omitempty"`

	// +kubebuilder:validation:Optional
	Cores *int64 `json:"cores,omitempty" tf:"cores,omitempty"`

	// +kubebuilder:validation:Required
	HostType *string `json:"hostType" tf:"host_type,omitempty"`

	// +kubebuilder:validation:Optional
	HostTypeName *string `json:"hostTypeName,omitempty" tf:"host_type_name,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceTotal *int64 `json:"instanceTotal,omitempty" tf:"instance_total,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceUuids []*string `json:"instanceUuids,omitempty" tf:"instance_uuids,omitempty"`

	// +kubebuilder:validation:Optional
	Memory *int64 `json:"memory,omitempty" tf:"memory,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	Sockets *int64 `json:"sockets,omitempty" tf:"sockets,omitempty"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// +kubebuilder:validation:Optional
	Vcpus *int64 `json:"vcpus,omitempty" tf:"vcpus,omitempty"`
}

// HostV1Spec defines the desired state of HostV1
type HostV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HostV1Parameters `json:"forProvider"`
}

// HostV1Status defines the observed state of HostV1.
type HostV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HostV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HostV1 is the Schema for the HostV1s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,otcjet}
type HostV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HostV1Spec   `json:"spec"`
	Status            HostV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostV1List contains a list of HostV1s
type HostV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostV1 `json:"items"`
}

// Repository type metadata.
var (
	HostV1_Kind             = "HostV1"
	HostV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HostV1_Kind}.String()
	HostV1_KindAPIVersion   = HostV1_Kind + "." + CRDGroupVersion.String()
	HostV1_GroupVersionKind = CRDGroupVersion.WithKind(HostV1_Kind)
)

func init() {
	SchemeBuilder.Register(&HostV1{}, &HostV1List{})
}