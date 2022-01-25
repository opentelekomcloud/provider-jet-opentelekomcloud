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

type RoleAssignmentV3Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RoleAssignmentV3Parameters struct {

	// +kubebuilder:validation:Optional
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Required
	RoleID *string `json:"roleId" tf:"role_id,omitempty"`

	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

// RoleAssignmentV3Spec defines the desired state of RoleAssignmentV3
type RoleAssignmentV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleAssignmentV3Parameters `json:"forProvider"`
}

// RoleAssignmentV3Status defines the observed state of RoleAssignmentV3.
type RoleAssignmentV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleAssignmentV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RoleAssignmentV3 is the Schema for the RoleAssignmentV3s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,otcjet}
type RoleAssignmentV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleAssignmentV3Spec   `json:"spec"`
	Status            RoleAssignmentV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleAssignmentV3List contains a list of RoleAssignmentV3s
type RoleAssignmentV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleAssignmentV3 `json:"items"`
}

// Repository type metadata.
var (
	RoleAssignmentV3_Kind             = "RoleAssignmentV3"
	RoleAssignmentV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RoleAssignmentV3_Kind}.String()
	RoleAssignmentV3_KindAPIVersion   = RoleAssignmentV3_Kind + "." + CRDGroupVersion.String()
	RoleAssignmentV3_GroupVersionKind = CRDGroupVersion.WithKind(RoleAssignmentV3_Kind)
)

func init() {
	SchemeBuilder.Register(&RoleAssignmentV3{}, &RoleAssignmentV3List{})
}
