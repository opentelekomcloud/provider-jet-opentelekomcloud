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

type MonitorV2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MonitorV2Parameters struct {

	// +kubebuilder:validation:Optional
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// +kubebuilder:validation:Required
	Delay *int64 `json:"delay" tf:"delay,omitempty"`

	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// +kubebuilder:validation:Optional
	ExpectedCodes *string `json:"expectedCodes,omitempty" tf:"expected_codes,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	// +kubebuilder:validation:Required
	MaxRetries *int64 `json:"maxRetries" tf:"max_retries,omitempty"`

	// +kubebuilder:validation:Optional
	MonitorPort *int64 `json:"monitorPort,omitempty" tf:"monitor_port,omitempty"`

	// +kubebuilder:validation:Required
	PoolID *string `json:"poolId" tf:"pool_id,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// +kubebuilder:validation:Required
	Timeout *int64 `json:"timeout" tf:"timeout,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	URLPath *string `json:"urlPath,omitempty" tf:"url_path,omitempty"`
}

// MonitorV2Spec defines the desired state of MonitorV2
type MonitorV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorV2Parameters `json:"forProvider"`
}

// MonitorV2Status defines the observed state of MonitorV2.
type MonitorV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorV2 is the Schema for the MonitorV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,otcjet}
type MonitorV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorV2Spec   `json:"spec"`
	Status            MonitorV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorV2List contains a list of MonitorV2s
type MonitorV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitorV2 `json:"items"`
}

// Repository type metadata.
var (
	MonitorV2_Kind             = "MonitorV2"
	MonitorV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitorV2_Kind}.String()
	MonitorV2_KindAPIVersion   = MonitorV2_Kind + "." + CRDGroupVersion.String()
	MonitorV2_GroupVersionKind = CRDGroupVersion.WithKind(MonitorV2_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitorV2{}, &MonitorV2List{})
}
