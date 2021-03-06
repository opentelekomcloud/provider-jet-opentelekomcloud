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

type CertificateV2Observation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ExpireTime *string `json:"expireTime,omitempty" tf:"expire_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type CertificateV2Parameters struct {

	// +kubebuilder:validation:Required
	Certificate *string `json:"certificate" tf:"certificate,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateKey *string `json:"privateKey,omitempty" tf:"private_key,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// CertificateV2Spec defines the desired state of CertificateV2
type CertificateV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateV2Parameters `json:"forProvider"`
}

// CertificateV2Status defines the observed state of CertificateV2.
type CertificateV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateV2 is the Schema for the CertificateV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,otcjet}
type CertificateV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateV2Spec   `json:"spec"`
	Status            CertificateV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateV2List contains a list of CertificateV2s
type CertificateV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateV2 `json:"items"`
}

// Repository type metadata.
var (
	CertificateV2_Kind             = "CertificateV2"
	CertificateV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CertificateV2_Kind}.String()
	CertificateV2_KindAPIVersion   = CertificateV2_Kind + "." + CRDGroupVersion.String()
	CertificateV2_GroupVersionKind = CRDGroupVersion.WithKind(CertificateV2_Kind)
)

func init() {
	SchemeBuilder.Register(&CertificateV2{}, &CertificateV2List{})
}
