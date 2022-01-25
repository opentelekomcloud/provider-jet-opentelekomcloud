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

type WhiteblackipRuleV1Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type WhiteblackipRuleV1Parameters struct {

	// +kubebuilder:validation:Required
	Addr *string `json:"addr" tf:"addr,omitempty"`

	// +kubebuilder:validation:Required
	PolicyID *string `json:"policyId" tf:"policy_id,omitempty"`

	// +kubebuilder:validation:Optional
	White *int64 `json:"white,omitempty" tf:"white,omitempty"`
}

// WhiteblackipRuleV1Spec defines the desired state of WhiteblackipRuleV1
type WhiteblackipRuleV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WhiteblackipRuleV1Parameters `json:"forProvider"`
}

// WhiteblackipRuleV1Status defines the observed state of WhiteblackipRuleV1.
type WhiteblackipRuleV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WhiteblackipRuleV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WhiteblackipRuleV1 is the Schema for the WhiteblackipRuleV1s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,otcjet}
type WhiteblackipRuleV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WhiteblackipRuleV1Spec   `json:"spec"`
	Status            WhiteblackipRuleV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WhiteblackipRuleV1List contains a list of WhiteblackipRuleV1s
type WhiteblackipRuleV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WhiteblackipRuleV1 `json:"items"`
}

// Repository type metadata.
var (
	WhiteblackipRuleV1_Kind             = "WhiteblackipRuleV1"
	WhiteblackipRuleV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WhiteblackipRuleV1_Kind}.String()
	WhiteblackipRuleV1_KindAPIVersion   = WhiteblackipRuleV1_Kind + "." + CRDGroupVersion.String()
	WhiteblackipRuleV1_GroupVersionKind = CRDGroupVersion.WithKind(WhiteblackipRuleV1_Kind)
)

func init() {
	SchemeBuilder.Register(&WhiteblackipRuleV1{}, &WhiteblackipRuleV1List{})
}
