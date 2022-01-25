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
func (in *DpdObservation) DeepCopyInto(out *DpdObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DpdObservation.
func (in *DpdObservation) DeepCopy() *DpdObservation {
	if in == nil {
		return nil
	}
	out := new(DpdObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DpdParameters) DeepCopyInto(out *DpdParameters) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(int64)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DpdParameters.
func (in *DpdParameters) DeepCopy() *DpdParameters {
	if in == nil {
		return nil
	}
	out := new(DpdParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointGroupV2) DeepCopyInto(out *EndpointGroupV2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointGroupV2.
func (in *EndpointGroupV2) DeepCopy() *EndpointGroupV2 {
	if in == nil {
		return nil
	}
	out := new(EndpointGroupV2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EndpointGroupV2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointGroupV2List) DeepCopyInto(out *EndpointGroupV2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EndpointGroupV2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointGroupV2List.
func (in *EndpointGroupV2List) DeepCopy() *EndpointGroupV2List {
	if in == nil {
		return nil
	}
	out := new(EndpointGroupV2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EndpointGroupV2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointGroupV2Observation) DeepCopyInto(out *EndpointGroupV2Observation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointGroupV2Observation.
func (in *EndpointGroupV2Observation) DeepCopy() *EndpointGroupV2Observation {
	if in == nil {
		return nil
	}
	out := new(EndpointGroupV2Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointGroupV2Parameters) DeepCopyInto(out *EndpointGroupV2Parameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.ValueSpecs != nil {
		in, out := &in.ValueSpecs, &out.ValueSpecs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointGroupV2Parameters.
func (in *EndpointGroupV2Parameters) DeepCopy() *EndpointGroupV2Parameters {
	if in == nil {
		return nil
	}
	out := new(EndpointGroupV2Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointGroupV2Spec) DeepCopyInto(out *EndpointGroupV2Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointGroupV2Spec.
func (in *EndpointGroupV2Spec) DeepCopy() *EndpointGroupV2Spec {
	if in == nil {
		return nil
	}
	out := new(EndpointGroupV2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointGroupV2Status) DeepCopyInto(out *EndpointGroupV2Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointGroupV2Status.
func (in *EndpointGroupV2Status) DeepCopy() *EndpointGroupV2Status {
	if in == nil {
		return nil
	}
	out := new(EndpointGroupV2Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IkePolicyV2) DeepCopyInto(out *IkePolicyV2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IkePolicyV2.
func (in *IkePolicyV2) DeepCopy() *IkePolicyV2 {
	if in == nil {
		return nil
	}
	out := new(IkePolicyV2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IkePolicyV2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IkePolicyV2List) DeepCopyInto(out *IkePolicyV2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IkePolicyV2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IkePolicyV2List.
func (in *IkePolicyV2List) DeepCopy() *IkePolicyV2List {
	if in == nil {
		return nil
	}
	out := new(IkePolicyV2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IkePolicyV2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IkePolicyV2Observation) DeepCopyInto(out *IkePolicyV2Observation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IkePolicyV2Observation.
func (in *IkePolicyV2Observation) DeepCopy() *IkePolicyV2Observation {
	if in == nil {
		return nil
	}
	out := new(IkePolicyV2Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IkePolicyV2Parameters) DeepCopyInto(out *IkePolicyV2Parameters) {
	*out = *in
	if in.AuthAlgorithm != nil {
		in, out := &in.AuthAlgorithm, &out.AuthAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EncryptionAlgorithm != nil {
		in, out := &in.EncryptionAlgorithm, &out.EncryptionAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.IkeVersion != nil {
		in, out := &in.IkeVersion, &out.IkeVersion
		*out = new(string)
		**out = **in
	}
	if in.Lifetime != nil {
		in, out := &in.Lifetime, &out.Lifetime
		*out = make([]LifetimeParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Pfs != nil {
		in, out := &in.Pfs, &out.Pfs
		*out = new(string)
		**out = **in
	}
	if in.Phase1NegotiationMode != nil {
		in, out := &in.Phase1NegotiationMode, &out.Phase1NegotiationMode
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
	if in.ValueSpecs != nil {
		in, out := &in.ValueSpecs, &out.ValueSpecs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IkePolicyV2Parameters.
func (in *IkePolicyV2Parameters) DeepCopy() *IkePolicyV2Parameters {
	if in == nil {
		return nil
	}
	out := new(IkePolicyV2Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IkePolicyV2Spec) DeepCopyInto(out *IkePolicyV2Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IkePolicyV2Spec.
func (in *IkePolicyV2Spec) DeepCopy() *IkePolicyV2Spec {
	if in == nil {
		return nil
	}
	out := new(IkePolicyV2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IkePolicyV2Status) DeepCopyInto(out *IkePolicyV2Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IkePolicyV2Status.
func (in *IkePolicyV2Status) DeepCopy() *IkePolicyV2Status {
	if in == nil {
		return nil
	}
	out := new(IkePolicyV2Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpsecPolicyV2) DeepCopyInto(out *IpsecPolicyV2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpsecPolicyV2.
func (in *IpsecPolicyV2) DeepCopy() *IpsecPolicyV2 {
	if in == nil {
		return nil
	}
	out := new(IpsecPolicyV2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IpsecPolicyV2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpsecPolicyV2LifetimeObservation) DeepCopyInto(out *IpsecPolicyV2LifetimeObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpsecPolicyV2LifetimeObservation.
func (in *IpsecPolicyV2LifetimeObservation) DeepCopy() *IpsecPolicyV2LifetimeObservation {
	if in == nil {
		return nil
	}
	out := new(IpsecPolicyV2LifetimeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpsecPolicyV2LifetimeParameters) DeepCopyInto(out *IpsecPolicyV2LifetimeParameters) {
	*out = *in
	if in.Units != nil {
		in, out := &in.Units, &out.Units
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpsecPolicyV2LifetimeParameters.
func (in *IpsecPolicyV2LifetimeParameters) DeepCopy() *IpsecPolicyV2LifetimeParameters {
	if in == nil {
		return nil
	}
	out := new(IpsecPolicyV2LifetimeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpsecPolicyV2List) DeepCopyInto(out *IpsecPolicyV2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IpsecPolicyV2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpsecPolicyV2List.
func (in *IpsecPolicyV2List) DeepCopy() *IpsecPolicyV2List {
	if in == nil {
		return nil
	}
	out := new(IpsecPolicyV2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IpsecPolicyV2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpsecPolicyV2Observation) DeepCopyInto(out *IpsecPolicyV2Observation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpsecPolicyV2Observation.
func (in *IpsecPolicyV2Observation) DeepCopy() *IpsecPolicyV2Observation {
	if in == nil {
		return nil
	}
	out := new(IpsecPolicyV2Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpsecPolicyV2Parameters) DeepCopyInto(out *IpsecPolicyV2Parameters) {
	*out = *in
	if in.AuthAlgorithm != nil {
		in, out := &in.AuthAlgorithm, &out.AuthAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EncapsulationMode != nil {
		in, out := &in.EncapsulationMode, &out.EncapsulationMode
		*out = new(string)
		**out = **in
	}
	if in.EncryptionAlgorithm != nil {
		in, out := &in.EncryptionAlgorithm, &out.EncryptionAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.Lifetime != nil {
		in, out := &in.Lifetime, &out.Lifetime
		*out = make([]IpsecPolicyV2LifetimeParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Pfs != nil {
		in, out := &in.Pfs, &out.Pfs
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
	if in.TransformProtocol != nil {
		in, out := &in.TransformProtocol, &out.TransformProtocol
		*out = new(string)
		**out = **in
	}
	if in.ValueSpecs != nil {
		in, out := &in.ValueSpecs, &out.ValueSpecs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpsecPolicyV2Parameters.
func (in *IpsecPolicyV2Parameters) DeepCopy() *IpsecPolicyV2Parameters {
	if in == nil {
		return nil
	}
	out := new(IpsecPolicyV2Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpsecPolicyV2Spec) DeepCopyInto(out *IpsecPolicyV2Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpsecPolicyV2Spec.
func (in *IpsecPolicyV2Spec) DeepCopy() *IpsecPolicyV2Spec {
	if in == nil {
		return nil
	}
	out := new(IpsecPolicyV2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpsecPolicyV2Status) DeepCopyInto(out *IpsecPolicyV2Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpsecPolicyV2Status.
func (in *IpsecPolicyV2Status) DeepCopy() *IpsecPolicyV2Status {
	if in == nil {
		return nil
	}
	out := new(IpsecPolicyV2Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LifetimeObservation) DeepCopyInto(out *LifetimeObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LifetimeObservation.
func (in *LifetimeObservation) DeepCopy() *LifetimeObservation {
	if in == nil {
		return nil
	}
	out := new(LifetimeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LifetimeParameters) DeepCopyInto(out *LifetimeParameters) {
	*out = *in
	if in.Units != nil {
		in, out := &in.Units, &out.Units
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LifetimeParameters.
func (in *LifetimeParameters) DeepCopy() *LifetimeParameters {
	if in == nil {
		return nil
	}
	out := new(LifetimeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceV2) DeepCopyInto(out *ServiceV2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceV2.
func (in *ServiceV2) DeepCopy() *ServiceV2 {
	if in == nil {
		return nil
	}
	out := new(ServiceV2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceV2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceV2List) DeepCopyInto(out *ServiceV2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceV2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceV2List.
func (in *ServiceV2List) DeepCopy() *ServiceV2List {
	if in == nil {
		return nil
	}
	out := new(ServiceV2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceV2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceV2Observation) DeepCopyInto(out *ServiceV2Observation) {
	*out = *in
	if in.ExternalV4IP != nil {
		in, out := &in.ExternalV4IP, &out.ExternalV4IP
		*out = new(string)
		**out = **in
	}
	if in.ExternalV6IP != nil {
		in, out := &in.ExternalV6IP, &out.ExternalV6IP
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceV2Observation.
func (in *ServiceV2Observation) DeepCopy() *ServiceV2Observation {
	if in == nil {
		return nil
	}
	out := new(ServiceV2Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceV2Parameters) DeepCopyInto(out *ServiceV2Parameters) {
	*out = *in
	if in.AdminStateUp != nil {
		in, out := &in.AdminStateUp, &out.AdminStateUp
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RouterID != nil {
		in, out := &in.RouterID, &out.RouterID
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
	if in.ValueSpecs != nil {
		in, out := &in.ValueSpecs, &out.ValueSpecs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceV2Parameters.
func (in *ServiceV2Parameters) DeepCopy() *ServiceV2Parameters {
	if in == nil {
		return nil
	}
	out := new(ServiceV2Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceV2Spec) DeepCopyInto(out *ServiceV2Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceV2Spec.
func (in *ServiceV2Spec) DeepCopy() *ServiceV2Spec {
	if in == nil {
		return nil
	}
	out := new(ServiceV2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceV2Status) DeepCopyInto(out *ServiceV2Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceV2Status.
func (in *ServiceV2Status) DeepCopy() *ServiceV2Status {
	if in == nil {
		return nil
	}
	out := new(ServiceV2Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteConnectionV2) DeepCopyInto(out *SiteConnectionV2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteConnectionV2.
func (in *SiteConnectionV2) DeepCopy() *SiteConnectionV2 {
	if in == nil {
		return nil
	}
	out := new(SiteConnectionV2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SiteConnectionV2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteConnectionV2List) DeepCopyInto(out *SiteConnectionV2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SiteConnectionV2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteConnectionV2List.
func (in *SiteConnectionV2List) DeepCopy() *SiteConnectionV2List {
	if in == nil {
		return nil
	}
	out := new(SiteConnectionV2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SiteConnectionV2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteConnectionV2Observation) DeepCopyInto(out *SiteConnectionV2Observation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteConnectionV2Observation.
func (in *SiteConnectionV2Observation) DeepCopy() *SiteConnectionV2Observation {
	if in == nil {
		return nil
	}
	out := new(SiteConnectionV2Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteConnectionV2Parameters) DeepCopyInto(out *SiteConnectionV2Parameters) {
	*out = *in
	if in.AdminStateUp != nil {
		in, out := &in.AdminStateUp, &out.AdminStateUp
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Dpd != nil {
		in, out := &in.Dpd, &out.Dpd
		*out = make([]DpdParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IkepolicyID != nil {
		in, out := &in.IkepolicyID, &out.IkepolicyID
		*out = new(string)
		**out = **in
	}
	if in.Initiator != nil {
		in, out := &in.Initiator, &out.Initiator
		*out = new(string)
		**out = **in
	}
	if in.IpsecpolicyID != nil {
		in, out := &in.IpsecpolicyID, &out.IpsecpolicyID
		*out = new(string)
		**out = **in
	}
	if in.LocalEpGroupID != nil {
		in, out := &in.LocalEpGroupID, &out.LocalEpGroupID
		*out = new(string)
		**out = **in
	}
	if in.LocalID != nil {
		in, out := &in.LocalID, &out.LocalID
		*out = new(string)
		**out = **in
	}
	if in.Mtu != nil {
		in, out := &in.Mtu, &out.Mtu
		*out = new(int64)
		**out = **in
	}
	if in.PeerAddress != nil {
		in, out := &in.PeerAddress, &out.PeerAddress
		*out = new(string)
		**out = **in
	}
	if in.PeerCidrs != nil {
		in, out := &in.PeerCidrs, &out.PeerCidrs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PeerEpGroupID != nil {
		in, out := &in.PeerEpGroupID, &out.PeerEpGroupID
		*out = new(string)
		**out = **in
	}
	if in.PeerID != nil {
		in, out := &in.PeerID, &out.PeerID
		*out = new(string)
		**out = **in
	}
	if in.Psk != nil {
		in, out := &in.Psk, &out.Psk
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
	if in.ValueSpecs != nil {
		in, out := &in.ValueSpecs, &out.ValueSpecs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.VpnserviceID != nil {
		in, out := &in.VpnserviceID, &out.VpnserviceID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteConnectionV2Parameters.
func (in *SiteConnectionV2Parameters) DeepCopy() *SiteConnectionV2Parameters {
	if in == nil {
		return nil
	}
	out := new(SiteConnectionV2Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteConnectionV2Spec) DeepCopyInto(out *SiteConnectionV2Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteConnectionV2Spec.
func (in *SiteConnectionV2Spec) DeepCopy() *SiteConnectionV2Spec {
	if in == nil {
		return nil
	}
	out := new(SiteConnectionV2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteConnectionV2Status) DeepCopyInto(out *SiteConnectionV2Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteConnectionV2Status.
func (in *SiteConnectionV2Status) DeepCopy() *SiteConnectionV2Status {
	if in == nil {
		return nil
	}
	out := new(SiteConnectionV2Status)
	in.DeepCopyInto(out)
	return out
}
