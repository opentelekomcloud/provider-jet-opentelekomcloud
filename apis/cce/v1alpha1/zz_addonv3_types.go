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

type AddonV3Observation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AddonV3Parameters struct {

	// +kubebuilder:validation:Required
	ClusterID *string `json:"clusterId" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Required
	TemplateName *string `json:"templateName" tf:"template_name,omitempty"`

	// +kubebuilder:validation:Required
	TemplateVersion *string `json:"templateVersion" tf:"template_version,omitempty"`

	// +kubebuilder:validation:Required
	Values []ValuesParameters `json:"values" tf:"values,omitempty"`
}

type ValuesObservation struct {
}

type ValuesParameters struct {

	// +kubebuilder:validation:Required
	Basic map[string]string `json:"basic" tf:"basic,omitempty"`

	// +kubebuilder:validation:Required
	Custom map[string]string `json:"custom" tf:"custom,omitempty"`
}

// AddonV3Spec defines the desired state of AddonV3
type AddonV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AddonV3Parameters `json:"forProvider"`
}

// AddonV3Status defines the observed state of AddonV3.
type AddonV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AddonV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AddonV3 is the Schema for the AddonV3s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,otcjet}
type AddonV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AddonV3Spec   `json:"spec"`
	Status            AddonV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AddonV3List contains a list of AddonV3s
type AddonV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AddonV3 `json:"items"`
}

// Repository type metadata.
var (
	AddonV3_Kind             = "AddonV3"
	AddonV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AddonV3_Kind}.String()
	AddonV3_KindAPIVersion   = AddonV3_Kind + "." + CRDGroupVersion.String()
	AddonV3_GroupVersionKind = CRDGroupVersion.WithKind(AddonV3_Kind)
)

func init() {
	SchemeBuilder.Register(&AddonV3{}, &AddonV3List{})
}