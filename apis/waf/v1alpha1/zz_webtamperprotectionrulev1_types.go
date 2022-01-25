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

type WebtamperprotectionRuleV1Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type WebtamperprotectionRuleV1Parameters struct {

	// +kubebuilder:validation:Required
	Hostname *string `json:"hostname" tf:"hostname,omitempty"`

	// +kubebuilder:validation:Required
	PolicyID *string `json:"policyId" tf:"policy_id,omitempty"`

	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`
}

// WebtamperprotectionRuleV1Spec defines the desired state of WebtamperprotectionRuleV1
type WebtamperprotectionRuleV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WebtamperprotectionRuleV1Parameters `json:"forProvider"`
}

// WebtamperprotectionRuleV1Status defines the observed state of WebtamperprotectionRuleV1.
type WebtamperprotectionRuleV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WebtamperprotectionRuleV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WebtamperprotectionRuleV1 is the Schema for the WebtamperprotectionRuleV1s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,otcjet}
type WebtamperprotectionRuleV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WebtamperprotectionRuleV1Spec   `json:"spec"`
	Status            WebtamperprotectionRuleV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebtamperprotectionRuleV1List contains a list of WebtamperprotectionRuleV1s
type WebtamperprotectionRuleV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebtamperprotectionRuleV1 `json:"items"`
}

// Repository type metadata.
var (
	WebtamperprotectionRuleV1_Kind             = "WebtamperprotectionRuleV1"
	WebtamperprotectionRuleV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WebtamperprotectionRuleV1_Kind}.String()
	WebtamperprotectionRuleV1_KindAPIVersion   = WebtamperprotectionRuleV1_Kind + "." + CRDGroupVersion.String()
	WebtamperprotectionRuleV1_GroupVersionKind = CRDGroupVersion.WithKind(WebtamperprotectionRuleV1_Kind)
)

func init() {
	SchemeBuilder.Register(&WebtamperprotectionRuleV1{}, &WebtamperprotectionRuleV1List{})
}
