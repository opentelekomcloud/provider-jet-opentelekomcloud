//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStrategyObservation) DeepCopyInto(out *BackupStrategyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStrategyObservation.
func (in *BackupStrategyObservation) DeepCopy() *BackupStrategyObservation {
	if in == nil {
		return nil
	}
	out := new(BackupStrategyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStrategyParameters) DeepCopyInto(out *BackupStrategyParameters) {
	*out = *in
	if in.KeepDays != nil {
		in, out := &in.KeepDays, &out.KeepDays
		*out = new(int64)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStrategyParameters.
func (in *BackupStrategyParameters) DeepCopy() *BackupStrategyParameters {
	if in == nil {
		return nil
	}
	out := new(BackupStrategyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreObservation) DeepCopyInto(out *DatastoreObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreObservation.
func (in *DatastoreObservation) DeepCopy() *DatastoreObservation {
	if in == nil {
		return nil
	}
	out := new(DatastoreObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreParameters) DeepCopyInto(out *DatastoreParameters) {
	*out = *in
	if in.StorageEngine != nil {
		in, out := &in.StorageEngine, &out.StorageEngine
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreParameters.
func (in *DatastoreParameters) DeepCopy() *DatastoreParameters {
	if in == nil {
		return nil
	}
	out := new(DatastoreParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlavorObservation) DeepCopyInto(out *FlavorObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlavorObservation.
func (in *FlavorObservation) DeepCopy() *FlavorObservation {
	if in == nil {
		return nil
	}
	out := new(FlavorObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlavorParameters) DeepCopyInto(out *FlavorParameters) {
	*out = *in
	if in.Num != nil {
		in, out := &in.Num, &out.Num
		*out = new(int64)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(int64)
		**out = **in
	}
	if in.SpecCode != nil {
		in, out := &in.SpecCode, &out.SpecCode
		*out = new(string)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlavorParameters.
func (in *FlavorParameters) DeepCopy() *FlavorParameters {
	if in == nil {
		return nil
	}
	out := new(FlavorParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceV3) DeepCopyInto(out *InstanceV3) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceV3.
func (in *InstanceV3) DeepCopy() *InstanceV3 {
	if in == nil {
		return nil
	}
	out := new(InstanceV3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceV3) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceV3List) DeepCopyInto(out *InstanceV3List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InstanceV3, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceV3List.
func (in *InstanceV3List) DeepCopy() *InstanceV3List {
	if in == nil {
		return nil
	}
	out := new(InstanceV3List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceV3List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceV3Observation) DeepCopyInto(out *InstanceV3Observation) {
	*out = *in
	if in.DBUsername != nil {
		in, out := &in.DBUsername, &out.DBUsername
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]NodesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PayMode != nil {
		in, out := &in.PayMode, &out.PayMode
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceV3Observation.
func (in *InstanceV3Observation) DeepCopy() *InstanceV3Observation {
	if in == nil {
		return nil
	}
	out := new(InstanceV3Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceV3Parameters) DeepCopyInto(out *InstanceV3Parameters) {
	*out = *in
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.BackupStrategy != nil {
		in, out := &in.BackupStrategy, &out.BackupStrategy
		*out = make([]BackupStrategyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Datastore != nil {
		in, out := &in.Datastore, &out.Datastore
		*out = make([]DatastoreParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DiskEncryptionID != nil {
		in, out := &in.DiskEncryptionID, &out.DiskEncryptionID
		*out = new(string)
		**out = **in
	}
	if in.Flavor != nil {
		in, out := &in.Flavor, &out.Flavor
		*out = make([]FlavorParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	out.PasswordSecretRef = in.PasswordSecretRef
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SSL != nil {
		in, out := &in.SSL, &out.SSL
		*out = new(bool)
		**out = **in
	}
	if in.SecurityGroupID != nil {
		in, out := &in.SecurityGroupID, &out.SecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceV3Parameters.
func (in *InstanceV3Parameters) DeepCopy() *InstanceV3Parameters {
	if in == nil {
		return nil
	}
	out := new(InstanceV3Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceV3Spec) DeepCopyInto(out *InstanceV3Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceV3Spec.
func (in *InstanceV3Spec) DeepCopy() *InstanceV3Spec {
	if in == nil {
		return nil
	}
	out := new(InstanceV3Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceV3Status) DeepCopyInto(out *InstanceV3Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceV3Status.
func (in *InstanceV3Status) DeepCopy() *InstanceV3Status {
	if in == nil {
		return nil
	}
	out := new(InstanceV3Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodesObservation) DeepCopyInto(out *NodesObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PrivateIP != nil {
		in, out := &in.PrivateIP, &out.PrivateIP
		*out = new(string)
		**out = **in
	}
	if in.PublicIP != nil {
		in, out := &in.PublicIP, &out.PublicIP
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodesObservation.
func (in *NodesObservation) DeepCopy() *NodesObservation {
	if in == nil {
		return nil
	}
	out := new(NodesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodesParameters) DeepCopyInto(out *NodesParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodesParameters.
func (in *NodesParameters) DeepCopy() *NodesParameters {
	if in == nil {
		return nil
	}
	out := new(NodesParameters)
	in.DeepCopyInto(out)
	return out
}
