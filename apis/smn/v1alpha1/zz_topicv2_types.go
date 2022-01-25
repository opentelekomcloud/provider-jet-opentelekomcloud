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

type TopicV2Observation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	PushPolicy *int64 `json:"pushPolicy,omitempty" tf:"push_policy,omitempty"`

	TopicUrn *string `json:"topicUrn,omitempty" tf:"topic_urn,omitempty"`

	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type TopicV2Parameters struct {

	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectName *string `json:"projectName,omitempty" tf:"project_name,omitempty"`
}

// TopicV2Spec defines the desired state of TopicV2
type TopicV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TopicV2Parameters `json:"forProvider"`
}

// TopicV2Status defines the observed state of TopicV2.
type TopicV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TopicV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TopicV2 is the Schema for the TopicV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,otcjet}
type TopicV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TopicV2Spec   `json:"spec"`
	Status            TopicV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TopicV2List contains a list of TopicV2s
type TopicV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TopicV2 `json:"items"`
}

// Repository type metadata.
var (
	TopicV2_Kind             = "TopicV2"
	TopicV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TopicV2_Kind}.String()
	TopicV2_KindAPIVersion   = TopicV2_Kind + "." + CRDGroupVersion.String()
	TopicV2_GroupVersionKind = CRDGroupVersion.WithKind(TopicV2_Kind)
)

func init() {
	SchemeBuilder.Register(&TopicV2{}, &TopicV2List{})
}
