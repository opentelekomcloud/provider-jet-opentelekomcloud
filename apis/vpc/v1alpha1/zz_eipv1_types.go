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

type BandwidthObservation struct {
}

type BandwidthParameters struct {

	// +kubebuilder:validation:Optional
	ChargeMode *string `json:"chargeMode,omitempty" tf:"charge_mode,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ShareType *string `json:"shareType" tf:"share_type,omitempty"`

	// +kubebuilder:validation:Required
	Size *int64 `json:"size" tf:"size,omitempty"`
}

type EIPV1Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EIPV1Parameters struct {

	// +kubebuilder:validation:Required
	Bandwidth []BandwidthParameters `json:"bandwidth" tf:"bandwidth,omitempty"`

	// +kubebuilder:validation:Required
	Publicip []PublicipParameters `json:"publicip" tf:"publicip,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	ValueSpecs map[string]string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

type PublicipObservation struct {
}

type PublicipParameters struct {

	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// EIPV1Spec defines the desired state of EIPV1
type EIPV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EIPV1Parameters `json:"forProvider"`
}

// EIPV1Status defines the observed state of EIPV1.
type EIPV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EIPV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EIPV1 is the Schema for the EIPV1s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,otcjet}
type EIPV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EIPV1Spec   `json:"spec"`
	Status            EIPV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EIPV1List contains a list of EIPV1s
type EIPV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EIPV1 `json:"items"`
}

// Repository type metadata.
var (
	EIPV1_Kind             = "EIPV1"
	EIPV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EIPV1_Kind}.String()
	EIPV1_KindAPIVersion   = EIPV1_Kind + "." + CRDGroupVersion.String()
	EIPV1_GroupVersionKind = CRDGroupVersion.WithKind(EIPV1_Kind)
)

func init() {
	SchemeBuilder.Register(&EIPV1{}, &EIPV1List{})
}