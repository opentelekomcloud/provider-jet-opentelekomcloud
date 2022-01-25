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

type PortObservation struct {
}

type PortParameters struct {

	// +kubebuilder:validation:Required
	ClientPort *int64 `json:"clientPort" tf:"client_port,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Required
	ServerPort *int64 `json:"serverPort" tf:"server_port,omitempty"`
}

type ServiceV1Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ServiceV1Parameters struct {

	// +kubebuilder:validation:Optional
	ApprovalEnabled *bool `json:"approvalEnabled,omitempty" tf:"approval_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	PoolID *string `json:"poolId,omitempty" tf:"pool_id,omitempty"`

	// +kubebuilder:validation:Required
	Port []PortParameters `json:"port" tf:"port,omitempty"`

	// +kubebuilder:validation:Required
	PortID *string `json:"portId" tf:"port_id,omitempty"`

	// +kubebuilder:validation:Required
	ServerType *string `json:"serverType" tf:"server_type,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceType *string `json:"serviceType,omitempty" tf:"service_type,omitempty"`

	// +kubebuilder:validation:Optional
	TCPProxy *string `json:"tcpProxy,omitempty" tf:"tcp_proxy,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	VPCID *string `json:"vpcId" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	VipPortID *string `json:"vipPortId,omitempty" tf:"vip_port_id,omitempty"`
}

// ServiceV1Spec defines the desired state of ServiceV1
type ServiceV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceV1Parameters `json:"forProvider"`
}

// ServiceV1Status defines the observed state of ServiceV1.
type ServiceV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceV1 is the Schema for the ServiceV1s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,otcjet}
type ServiceV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceV1Spec   `json:"spec"`
	Status            ServiceV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceV1List contains a list of ServiceV1s
type ServiceV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceV1 `json:"items"`
}

// Repository type metadata.
var (
	ServiceV1_Kind             = "ServiceV1"
	ServiceV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceV1_Kind}.String()
	ServiceV1_KindAPIVersion   = ServiceV1_Kind + "." + CRDGroupVersion.String()
	ServiceV1_GroupVersionKind = CRDGroupVersion.WithKind(ServiceV1_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceV1{}, &ServiceV1List{})
}
