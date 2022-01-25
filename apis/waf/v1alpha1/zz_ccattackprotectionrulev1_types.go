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

type CcattackprotectionRuleV1Observation struct {
	Default *bool `json:"default,omitempty" tf:"default,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CcattackprotectionRuleV1Parameters struct {

	// +kubebuilder:validation:Required
	ActionCategory *string `json:"actionCategory" tf:"action_category,omitempty"`

	// +kubebuilder:validation:Optional
	BlockContent *string `json:"blockContent,omitempty" tf:"block_content,omitempty"`

	// +kubebuilder:validation:Optional
	BlockContentType *string `json:"blockContentType,omitempty" tf:"block_content_type,omitempty"`

	// +kubebuilder:validation:Required
	LimitNum *int64 `json:"limitNum" tf:"limit_num,omitempty"`

	// +kubebuilder:validation:Required
	LimitPeriod *int64 `json:"limitPeriod" tf:"limit_period,omitempty"`

	// +kubebuilder:validation:Optional
	LockTime *int64 `json:"lockTime,omitempty" tf:"lock_time,omitempty"`

	// +kubebuilder:validation:Required
	PolicyID *string `json:"policyId" tf:"policy_id,omitempty"`

	// +kubebuilder:validation:Optional
	TagCategory *string `json:"tagCategory,omitempty" tf:"tag_category,omitempty"`

	// +kubebuilder:validation:Optional
	TagContents []*string `json:"tagContents,omitempty" tf:"tag_contents,omitempty"`

	// +kubebuilder:validation:Optional
	TagIndex *string `json:"tagIndex,omitempty" tf:"tag_index,omitempty"`

	// +kubebuilder:validation:Required
	TagType *string `json:"tagType" tf:"tag_type,omitempty"`

	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`
}

// CcattackprotectionRuleV1Spec defines the desired state of CcattackprotectionRuleV1
type CcattackprotectionRuleV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CcattackprotectionRuleV1Parameters `json:"forProvider"`
}

// CcattackprotectionRuleV1Status defines the observed state of CcattackprotectionRuleV1.
type CcattackprotectionRuleV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CcattackprotectionRuleV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CcattackprotectionRuleV1 is the Schema for the CcattackprotectionRuleV1s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,otcjet}
type CcattackprotectionRuleV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CcattackprotectionRuleV1Spec   `json:"spec"`
	Status            CcattackprotectionRuleV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CcattackprotectionRuleV1List contains a list of CcattackprotectionRuleV1s
type CcattackprotectionRuleV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CcattackprotectionRuleV1 `json:"items"`
}

// Repository type metadata.
var (
	CcattackprotectionRuleV1_Kind             = "CcattackprotectionRuleV1"
	CcattackprotectionRuleV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CcattackprotectionRuleV1_Kind}.String()
	CcattackprotectionRuleV1_KindAPIVersion   = CcattackprotectionRuleV1_Kind + "." + CRDGroupVersion.String()
	CcattackprotectionRuleV1_GroupVersionKind = CRDGroupVersion.WithKind(CcattackprotectionRuleV1_Kind)
)

func init() {
	SchemeBuilder.Register(&CcattackprotectionRuleV1{}, &CcattackprotectionRuleV1List{})
}
