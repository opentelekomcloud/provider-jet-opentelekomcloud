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

type AgencyV3Observation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	ExpireTime *string `json:"expireTime,omitempty" tf:"expire_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AgencyV3Parameters struct {

	// +kubebuilder:validation:Required
	DelegatedDomainName *string `json:"delegatedDomainName" tf:"delegated_domain_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DomainRoles []*string `json:"domainRoles,omitempty" tf:"domain_roles,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectRole []ProjectRoleParameters `json:"projectRole,omitempty" tf:"project_role,omitempty"`
}

type ProjectRoleObservation struct {
}

type ProjectRoleParameters struct {

	// +kubebuilder:validation:Required
	Project *string `json:"project" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Roles []*string `json:"roles" tf:"roles,omitempty"`
}

// AgencyV3Spec defines the desired state of AgencyV3
type AgencyV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AgencyV3Parameters `json:"forProvider"`
}

// AgencyV3Status defines the observed state of AgencyV3.
type AgencyV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AgencyV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AgencyV3 is the Schema for the AgencyV3s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,otcjet}
type AgencyV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AgencyV3Spec   `json:"spec"`
	Status            AgencyV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AgencyV3List contains a list of AgencyV3s
type AgencyV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AgencyV3 `json:"items"`
}

// Repository type metadata.
var (
	AgencyV3_Kind             = "AgencyV3"
	AgencyV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AgencyV3_Kind}.String()
	AgencyV3_KindAPIVersion   = AgencyV3_Kind + "." + CRDGroupVersion.String()
	AgencyV3_GroupVersionKind = CRDGroupVersion.WithKind(AgencyV3_Kind)
)

func init() {
	SchemeBuilder.Register(&AgencyV3{}, &AgencyV3List{})
}
