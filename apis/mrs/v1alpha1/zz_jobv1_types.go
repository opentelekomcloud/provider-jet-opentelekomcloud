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

type JobV1Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	JobState *string `json:"jobState,omitempty" tf:"job_state,omitempty"`
}

type JobV1Parameters struct {

	// +kubebuilder:validation:Optional
	Arguments *string `json:"arguments,omitempty" tf:"arguments,omitempty"`

	// +kubebuilder:validation:Required
	ClusterID *string `json:"clusterId" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	HiveScriptPath *string `json:"hiveScriptPath,omitempty" tf:"hive_script_path,omitempty"`

	// +kubebuilder:validation:Optional
	Input *string `json:"input,omitempty" tf:"input,omitempty"`

	// +kubebuilder:validation:Optional
	IsProtected *bool `json:"isProtected,omitempty" tf:"is_protected,omitempty"`

	// +kubebuilder:validation:Optional
	IsPublic *bool `json:"isPublic,omitempty" tf:"is_public,omitempty"`

	// +kubebuilder:validation:Required
	JarPath *string `json:"jarPath" tf:"jar_path,omitempty"`

	// +kubebuilder:validation:Optional
	JobLog *string `json:"jobLog,omitempty" tf:"job_log,omitempty"`

	// +kubebuilder:validation:Required
	JobName *string `json:"jobName" tf:"job_name,omitempty"`

	// +kubebuilder:validation:Required
	JobType *int64 `json:"jobType" tf:"job_type,omitempty"`

	// +kubebuilder:validation:Optional
	Output *string `json:"output,omitempty" tf:"output,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// JobV1Spec defines the desired state of JobV1
type JobV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     JobV1Parameters `json:"forProvider"`
}

// JobV1Status defines the observed state of JobV1.
type JobV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        JobV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// JobV1 is the Schema for the JobV1s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,otcjet}
type JobV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              JobV1Spec   `json:"spec"`
	Status            JobV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// JobV1List contains a list of JobV1s
type JobV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JobV1 `json:"items"`
}

// Repository type metadata.
var (
	JobV1_Kind             = "JobV1"
	JobV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: JobV1_Kind}.String()
	JobV1_KindAPIVersion   = JobV1_Kind + "." + CRDGroupVersion.String()
	JobV1_GroupVersionKind = CRDGroupVersion.WithKind(JobV1_Kind)
)

func init() {
	SchemeBuilder.Register(&JobV1{}, &JobV1List{})
}
