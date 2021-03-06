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

type OperationDefinitionObservation struct {
}

type OperationDefinitionParameters struct {

	// +kubebuilder:validation:Optional
	DayBackups *int64 `json:"dayBackups,omitempty" tf:"day_backups,omitempty"`

	// +kubebuilder:validation:Optional
	MaxBackups *int64 `json:"maxBackups,omitempty" tf:"max_backups,omitempty"`

	// +kubebuilder:validation:Optional
	MonthBackups *int64 `json:"monthBackups,omitempty" tf:"month_backups,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionDurationDays *int64 `json:"retentionDurationDays,omitempty" tf:"retention_duration_days,omitempty"`

	// +kubebuilder:validation:Required
	Timezone *string `json:"timezone" tf:"timezone,omitempty"`

	// +kubebuilder:validation:Optional
	WeekBackups *int64 `json:"weekBackups,omitempty" tf:"week_backups,omitempty"`

	// +kubebuilder:validation:Optional
	YearBackups *int64 `json:"yearBackups,omitempty" tf:"year_backups,omitempty"`
}

type PolicyV3Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type PolicyV3Parameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	OperationDefinition []OperationDefinitionParameters `json:"operationDefinition,omitempty" tf:"operation_definition,omitempty"`

	// +kubebuilder:validation:Required
	OperationType *string `json:"operationType" tf:"operation_type,omitempty"`

	// +kubebuilder:validation:Required
	TriggerPattern []*string `json:"triggerPattern" tf:"trigger_pattern,omitempty"`
}

// PolicyV3Spec defines the desired state of PolicyV3
type PolicyV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyV3Parameters `json:"forProvider"`
}

// PolicyV3Status defines the observed state of PolicyV3.
type PolicyV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyV3 is the Schema for the PolicyV3s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,otcjet}
type PolicyV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicyV3Spec   `json:"spec"`
	Status            PolicyV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyV3List contains a list of PolicyV3s
type PolicyV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyV3 `json:"items"`
}

// Repository type metadata.
var (
	PolicyV3_Kind             = "PolicyV3"
	PolicyV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyV3_Kind}.String()
	PolicyV3_KindAPIVersion   = PolicyV3_Kind + "." + CRDGroupVersion.String()
	PolicyV3_GroupVersionKind = CRDGroupVersion.WithKind(PolicyV3_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyV3{}, &PolicyV3List{})
}
