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

type EndpointV1Observation struct {
	DNSNames []*string `json:"dnsNames,omitempty" tf:"dns_names,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	MarkerID *int64 `json:"markerId,omitempty" tf:"marker_id,omitempty"`

	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	ServiceType *string `json:"serviceType,omitempty" tf:"service_type,omitempty"`
}

type EndpointV1Parameters struct {

	// +kubebuilder:validation:Optional
	EnableDNS *bool `json:"enableDns,omitempty" tf:"enable_dns,omitempty"`

	// +kubebuilder:validation:Optional
	EnableWhitelist *bool `json:"enableWhitelist,omitempty" tf:"enable_whitelist,omitempty"`

	// +kubebuilder:validation:Optional
	PortIP *string `json:"portIp,omitempty" tf:"port_ip,omitempty"`

	// +kubebuilder:validation:Optional
	RouteTables []*string `json:"routeTables,omitempty" tf:"route_tables,omitempty"`

	// +kubebuilder:validation:Required
	ServiceID *string `json:"serviceId" tf:"service_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	VPCID *string `json:"vpcId" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	Whitelist []*string `json:"whitelist,omitempty" tf:"whitelist,omitempty"`
}

// EndpointV1Spec defines the desired state of EndpointV1
type EndpointV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EndpointV1Parameters `json:"forProvider"`
}

// EndpointV1Status defines the observed state of EndpointV1.
type EndpointV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EndpointV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointV1 is the Schema for the EndpointV1s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,otcjet}
type EndpointV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointV1Spec   `json:"spec"`
	Status            EndpointV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointV1List contains a list of EndpointV1s
type EndpointV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EndpointV1 `json:"items"`
}

// Repository type metadata.
var (
	EndpointV1_Kind             = "EndpointV1"
	EndpointV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EndpointV1_Kind}.String()
	EndpointV1_KindAPIVersion   = EndpointV1_Kind + "." + CRDGroupVersion.String()
	EndpointV1_GroupVersionKind = CRDGroupVersion.WithKind(EndpointV1_Kind)
)

func init() {
	SchemeBuilder.Register(&EndpointV1{}, &EndpointV1List{})
}