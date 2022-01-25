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

type WhitelistV2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type WhitelistV2Parameters struct {

	// +kubebuilder:validation:Optional
	EnableWhitelist *bool `json:"enableWhitelist,omitempty" tf:"enable_whitelist,omitempty"`

	// +kubebuilder:validation:Required
	ListenerID *string `json:"listenerId" tf:"listener_id,omitempty"`

	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// +kubebuilder:validation:Optional
	Whitelist *string `json:"whitelist,omitempty" tf:"whitelist,omitempty"`
}

// WhitelistV2Spec defines the desired state of WhitelistV2
type WhitelistV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WhitelistV2Parameters `json:"forProvider"`
}

// WhitelistV2Status defines the observed state of WhitelistV2.
type WhitelistV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WhitelistV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WhitelistV2 is the Schema for the WhitelistV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,otcjet}
type WhitelistV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WhitelistV2Spec   `json:"spec"`
	Status            WhitelistV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WhitelistV2List contains a list of WhitelistV2s
type WhitelistV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WhitelistV2 `json:"items"`
}

// Repository type metadata.
var (
	WhitelistV2_Kind             = "WhitelistV2"
	WhitelistV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WhitelistV2_Kind}.String()
	WhitelistV2_KindAPIVersion   = WhitelistV2_Kind + "." + CRDGroupVersion.String()
	WhitelistV2_GroupVersionKind = CRDGroupVersion.WithKind(WhitelistV2_Kind)
)

func init() {
	SchemeBuilder.Register(&WhitelistV2{}, &WhitelistV2List{})
}