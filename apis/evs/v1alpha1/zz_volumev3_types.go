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

type AttachmentObservation struct {
	Device *string `json:"device,omitempty" tf:"device,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`
}

type AttachmentParameters struct {
}

type VolumeV3Observation struct {
	Attachment []AttachmentObservation `json:"attachment,omitempty" tf:"attachment,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Wwn *string `json:"wwn,omitempty" tf:"wwn,omitempty"`
}

type VolumeV3Parameters struct {

	// +kubebuilder:validation:Required
	AvailabilityZone *string `json:"availabilityZone" tf:"availability_zone,omitempty"`

	// +kubebuilder:validation:Optional
	BackupID *string `json:"backupId,omitempty" tf:"backup_id,omitempty"`

	// +kubebuilder:validation:Optional
	Cascade *bool `json:"cascade,omitempty" tf:"cascade,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DeviceType *string `json:"deviceType,omitempty" tf:"device_type,omitempty"`

	// +kubebuilder:validation:Optional
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// +kubebuilder:validation:Optional
	KMSID *string `json:"kmsId,omitempty" tf:"kms_id,omitempty"`

	// +kubebuilder:validation:Optional
	Multiattach *bool `json:"multiattach,omitempty" tf:"multiattach,omitempty"`

	// +kubebuilder:validation:Optional
	Size *int64 `json:"size,omitempty" tf:"size,omitempty"`

	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	VolumeType *string `json:"volumeType" tf:"volume_type,omitempty"`
}

// VolumeV3Spec defines the desired state of VolumeV3
type VolumeV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VolumeV3Parameters `json:"forProvider"`
}

// VolumeV3Status defines the observed state of VolumeV3.
type VolumeV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VolumeV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeV3 is the Schema for the VolumeV3s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,otcjet}
type VolumeV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VolumeV3Spec   `json:"spec"`
	Status            VolumeV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeV3List contains a list of VolumeV3s
type VolumeV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VolumeV3 `json:"items"`
}

// Repository type metadata.
var (
	VolumeV3_Kind             = "VolumeV3"
	VolumeV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VolumeV3_Kind}.String()
	VolumeV3_KindAPIVersion   = VolumeV3_Kind + "." + CRDGroupVersion.String()
	VolumeV3_GroupVersionKind = CRDGroupVersion.WithKind(VolumeV3_Kind)
)

func init() {
	SchemeBuilder.Register(&VolumeV3{}, &VolumeV3List{})
}