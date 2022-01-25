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
func (in *TrackerV1) DeepCopyInto(out *TrackerV1) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrackerV1.
func (in *TrackerV1) DeepCopy() *TrackerV1 {
	if in == nil {
		return nil
	}
	out := new(TrackerV1)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrackerV1) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrackerV1List) DeepCopyInto(out *TrackerV1List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TrackerV1, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrackerV1List.
func (in *TrackerV1List) DeepCopy() *TrackerV1List {
	if in == nil {
		return nil
	}
	out := new(TrackerV1List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrackerV1List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrackerV1Observation) DeepCopyInto(out *TrackerV1Observation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.TrackerName != nil {
		in, out := &in.TrackerName, &out.TrackerName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrackerV1Observation.
func (in *TrackerV1Observation) DeepCopy() *TrackerV1Observation {
	if in == nil {
		return nil
	}
	out := new(TrackerV1Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrackerV1Parameters) DeepCopyInto(out *TrackerV1Parameters) {
	*out = *in
	if in.BucketName != nil {
		in, out := &in.BucketName, &out.BucketName
		*out = new(string)
		**out = **in
	}
	if in.FilePrefixName != nil {
		in, out := &in.FilePrefixName, &out.FilePrefixName
		*out = new(string)
		**out = **in
	}
	if in.IsSendAllKeyOperation != nil {
		in, out := &in.IsSendAllKeyOperation, &out.IsSendAllKeyOperation
		*out = new(bool)
		**out = **in
	}
	if in.IsSupportSmn != nil {
		in, out := &in.IsSupportSmn, &out.IsSupportSmn
		*out = new(bool)
		**out = **in
	}
	if in.NeedNotifyUserList != nil {
		in, out := &in.NeedNotifyUserList, &out.NeedNotifyUserList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Operations != nil {
		in, out := &in.Operations, &out.Operations
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ProjectName != nil {
		in, out := &in.ProjectName, &out.ProjectName
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.TopicID != nil {
		in, out := &in.TopicID, &out.TopicID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrackerV1Parameters.
func (in *TrackerV1Parameters) DeepCopy() *TrackerV1Parameters {
	if in == nil {
		return nil
	}
	out := new(TrackerV1Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrackerV1Spec) DeepCopyInto(out *TrackerV1Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrackerV1Spec.
func (in *TrackerV1Spec) DeepCopy() *TrackerV1Spec {
	if in == nil {
		return nil
	}
	out := new(TrackerV1Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrackerV1Status) DeepCopyInto(out *TrackerV1Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrackerV1Status.
func (in *TrackerV1Status) DeepCopy() *TrackerV1Status {
	if in == nil {
		return nil
	}
	out := new(TrackerV1Status)
	in.DeepCopyInto(out)
	return out
}
