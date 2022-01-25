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

type BillingObservation struct {
	Allocated *int64 `json:"allocated,omitempty" tf:"allocated,omitempty"`

	FrozenScene *string `json:"frozenScene,omitempty" tf:"frozen_scene,omitempty"`

	OrderID *string `json:"orderId,omitempty" tf:"order_id,omitempty"`

	ProductID *string `json:"productId,omitempty" tf:"product_id,omitempty"`

	SpecCode *string `json:"specCode,omitempty" tf:"spec_code,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	StorageUnit *string `json:"storageUnit,omitempty" tf:"storage_unit,omitempty"`

	Used *int64 `json:"used,omitempty" tf:"used,omitempty"`
}

type BillingParameters struct {

	// +kubebuilder:validation:Optional
	ChargingMode *string `json:"chargingMode,omitempty" tf:"charging_mode,omitempty"`

	// +kubebuilder:validation:Optional
	CloudType *string `json:"cloudType,omitempty" tf:"cloud_type,omitempty"`

	// +kubebuilder:validation:Optional
	ConsistentLevel *string `json:"consistentLevel,omitempty" tf:"consistent_level,omitempty"`

	// +kubebuilder:validation:Optional
	ConsoleURL *string `json:"consoleUrl,omitempty" tf:"console_url,omitempty"`

	// +kubebuilder:validation:Optional
	ExtraInfo map[string]string `json:"extraInfo,omitempty" tf:"extra_info,omitempty"`

	// +kubebuilder:validation:Optional
	IsAutoPay *bool `json:"isAutoPay,omitempty" tf:"is_auto_pay,omitempty"`

	// +kubebuilder:validation:Optional
	IsAutoRenew *bool `json:"isAutoRenew,omitempty" tf:"is_auto_renew,omitempty"`

	// +kubebuilder:validation:Required
	ObjectType *string `json:"objectType" tf:"object_type,omitempty"`

	// +kubebuilder:validation:Optional
	PeriodNum *int64 `json:"periodNum,omitempty" tf:"period_num,omitempty"`

	// +kubebuilder:validation:Optional
	PeriodType *string `json:"periodType,omitempty" tf:"period_type,omitempty"`

	// +kubebuilder:validation:Required
	ProtectType *string `json:"protectType" tf:"protect_type,omitempty"`

	// +kubebuilder:validation:Required
	Size *int64 `json:"size" tf:"size,omitempty"`
}

type BindRulesObservation struct {
}

type BindRulesParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type ResourceObservation struct {
	BackupCount *int64 `json:"backupCount,omitempty" tf:"backup_count,omitempty"`

	BackupSize *int64 `json:"backupSize,omitempty" tf:"backup_size,omitempty"`

	ProtectStatus *string `json:"protectStatus,omitempty" tf:"protect_status,omitempty"`

	Size *int64 `json:"size,omitempty" tf:"size,omitempty"`
}

type ResourceParameters struct {

	// +kubebuilder:validation:Optional
	ExtraInfo map[string]string `json:"extraInfo,omitempty" tf:"extra_info,omitempty"`

	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type VaultV3Observation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	ProviderID *string `json:"providerId,omitempty" tf:"provider_id,omitempty"`

	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type VaultV3Parameters struct {

	// +kubebuilder:validation:Optional
	AutoBind *bool `json:"autoBind,omitempty" tf:"auto_bind,omitempty"`

	// +kubebuilder:validation:Optional
	AutoExpand *bool `json:"autoExpand,omitempty" tf:"auto_expand,omitempty"`

	// +kubebuilder:validation:Optional
	BackupPolicyID *string `json:"backupPolicyId,omitempty" tf:"backup_policy_id,omitempty"`

	// +kubebuilder:validation:Required
	Billing []BillingParameters `json:"billing" tf:"billing,omitempty"`

	// +kubebuilder:validation:Optional
	BindRules []BindRulesParameters `json:"bindRules,omitempty" tf:"bind_rules,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// +kubebuilder:validation:Optional
	Resource []ResourceParameters `json:"resource,omitempty" tf:"resource,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// VaultV3Spec defines the desired state of VaultV3
type VaultV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VaultV3Parameters `json:"forProvider"`
}

// VaultV3Status defines the observed state of VaultV3.
type VaultV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VaultV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VaultV3 is the Schema for the VaultV3s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,otcjet}
type VaultV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VaultV3Spec   `json:"spec"`
	Status            VaultV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VaultV3List contains a list of VaultV3s
type VaultV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VaultV3 `json:"items"`
}

// Repository type metadata.
var (
	VaultV3_Kind             = "VaultV3"
	VaultV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VaultV3_Kind}.String()
	VaultV3_KindAPIVersion   = VaultV3_Kind + "." + CRDGroupVersion.String()
	VaultV3_GroupVersionKind = CRDGroupVersion.WithKind(VaultV3_Kind)
)

func init() {
	SchemeBuilder.Register(&VaultV3{}, &VaultV3List{})
}