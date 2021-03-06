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

type ServergroupV2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Members []*string `json:"members,omitempty" tf:"members,omitempty"`
}

type ServergroupV2Parameters struct {

	// +kubebuilder:validation:Optional
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	ValueSpecs map[string]string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

// ServergroupV2Spec defines the desired state of ServergroupV2
type ServergroupV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServergroupV2Parameters `json:"forProvider"`
}

// ServergroupV2Status defines the observed state of ServergroupV2.
type ServergroupV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServergroupV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServergroupV2 is the Schema for the ServergroupV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,otcjet}
type ServergroupV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServergroupV2Spec   `json:"spec"`
	Status            ServergroupV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServergroupV2List contains a list of ServergroupV2s
type ServergroupV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServergroupV2 `json:"items"`
}

// Repository type metadata.
var (
	ServergroupV2_Kind             = "ServergroupV2"
	ServergroupV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServergroupV2_Kind}.String()
	ServergroupV2_KindAPIVersion   = ServergroupV2_Kind + "." + CRDGroupVersion.String()
	ServergroupV2_GroupVersionKind = CRDGroupVersion.WithKind(ServergroupV2_Kind)
)

func init() {
	SchemeBuilder.Register(&ServergroupV2{}, &ServergroupV2List{})
}
