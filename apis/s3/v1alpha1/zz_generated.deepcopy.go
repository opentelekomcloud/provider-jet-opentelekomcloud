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
func (in *Bucket) DeepCopyInto(out *Bucket) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket.
func (in *Bucket) DeepCopy() *Bucket {
	if in == nil {
		return nil
	}
	out := new(Bucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Bucket) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketList) DeepCopyInto(out *BucketList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Bucket, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketList.
func (in *BucketList) DeepCopy() *BucketList {
	if in == nil {
		return nil
	}
	out := new(BucketList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketObject) DeepCopyInto(out *BucketObject) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketObject.
func (in *BucketObject) DeepCopy() *BucketObject {
	if in == nil {
		return nil
	}
	out := new(BucketObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketObject) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketObjectList) DeepCopyInto(out *BucketObjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BucketObject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketObjectList.
func (in *BucketObjectList) DeepCopy() *BucketObjectList {
	if in == nil {
		return nil
	}
	out := new(BucketObjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketObjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketObjectObservation) DeepCopyInto(out *BucketObjectObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.VersionID != nil {
		in, out := &in.VersionID, &out.VersionID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketObjectObservation.
func (in *BucketObjectObservation) DeepCopy() *BucketObjectObservation {
	if in == nil {
		return nil
	}
	out := new(BucketObjectObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketObjectParameters) DeepCopyInto(out *BucketObjectParameters) {
	*out = *in
	if in.ACL != nil {
		in, out := &in.ACL, &out.ACL
		*out = new(string)
		**out = **in
	}
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.CacheControl != nil {
		in, out := &in.CacheControl, &out.CacheControl
		*out = new(string)
		**out = **in
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.ContentDisposition != nil {
		in, out := &in.ContentDisposition, &out.ContentDisposition
		*out = new(string)
		**out = **in
	}
	if in.ContentEncoding != nil {
		in, out := &in.ContentEncoding, &out.ContentEncoding
		*out = new(string)
		**out = **in
	}
	if in.ContentLanguage != nil {
		in, out := &in.ContentLanguage, &out.ContentLanguage
		*out = new(string)
		**out = **in
	}
	if in.ContentType != nil {
		in, out := &in.ContentType, &out.ContentType
		*out = new(string)
		**out = **in
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.ServerSideEncryption != nil {
		in, out := &in.ServerSideEncryption, &out.ServerSideEncryption
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.SseKMSKeyID != nil {
		in, out := &in.SseKMSKeyID, &out.SseKMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.WebsiteRedirect != nil {
		in, out := &in.WebsiteRedirect, &out.WebsiteRedirect
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketObjectParameters.
func (in *BucketObjectParameters) DeepCopy() *BucketObjectParameters {
	if in == nil {
		return nil
	}
	out := new(BucketObjectParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketObjectSpec) DeepCopyInto(out *BucketObjectSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketObjectSpec.
func (in *BucketObjectSpec) DeepCopy() *BucketObjectSpec {
	if in == nil {
		return nil
	}
	out := new(BucketObjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketObjectStatus) DeepCopyInto(out *BucketObjectStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketObjectStatus.
func (in *BucketObjectStatus) DeepCopy() *BucketObjectStatus {
	if in == nil {
		return nil
	}
	out := new(BucketObjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketObservation) DeepCopyInto(out *BucketObservation) {
	*out = *in
	if in.BucketDomainName != nil {
		in, out := &in.BucketDomainName, &out.BucketDomainName
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketObservation.
func (in *BucketObservation) DeepCopy() *BucketObservation {
	if in == nil {
		return nil
	}
	out := new(BucketObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketParameters) DeepCopyInto(out *BucketParameters) {
	*out = *in
	if in.ACL != nil {
		in, out := &in.ACL, &out.ACL
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.BucketPrefix != nil {
		in, out := &in.BucketPrefix, &out.BucketPrefix
		*out = new(string)
		**out = **in
	}
	if in.CorsRule != nil {
		in, out := &in.CorsRule, &out.CorsRule
		*out = make([]CorsRuleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ForceDestroy != nil {
		in, out := &in.ForceDestroy, &out.ForceDestroy
		*out = new(bool)
		**out = **in
	}
	if in.HostedZoneID != nil {
		in, out := &in.HostedZoneID, &out.HostedZoneID
		*out = new(string)
		**out = **in
	}
	if in.LifecycleRule != nil {
		in, out := &in.LifecycleRule, &out.LifecycleRule
		*out = make([]LifecycleRuleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = make([]LoggingParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
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
	if in.Versioning != nil {
		in, out := &in.Versioning, &out.Versioning
		*out = make([]VersioningParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Website != nil {
		in, out := &in.Website, &out.Website
		*out = make([]WebsiteParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WebsiteDomain != nil {
		in, out := &in.WebsiteDomain, &out.WebsiteDomain
		*out = new(string)
		**out = **in
	}
	if in.WebsiteEndpoint != nil {
		in, out := &in.WebsiteEndpoint, &out.WebsiteEndpoint
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketParameters.
func (in *BucketParameters) DeepCopy() *BucketParameters {
	if in == nil {
		return nil
	}
	out := new(BucketParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketPolicy) DeepCopyInto(out *BucketPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketPolicy.
func (in *BucketPolicy) DeepCopy() *BucketPolicy {
	if in == nil {
		return nil
	}
	out := new(BucketPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketPolicyList) DeepCopyInto(out *BucketPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BucketPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketPolicyList.
func (in *BucketPolicyList) DeepCopy() *BucketPolicyList {
	if in == nil {
		return nil
	}
	out := new(BucketPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketPolicyObservation) DeepCopyInto(out *BucketPolicyObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketPolicyObservation.
func (in *BucketPolicyObservation) DeepCopy() *BucketPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(BucketPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketPolicyParameters) DeepCopyInto(out *BucketPolicyParameters) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketPolicyParameters.
func (in *BucketPolicyParameters) DeepCopy() *BucketPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(BucketPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketPolicySpec) DeepCopyInto(out *BucketPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketPolicySpec.
func (in *BucketPolicySpec) DeepCopy() *BucketPolicySpec {
	if in == nil {
		return nil
	}
	out := new(BucketPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketPolicyStatus) DeepCopyInto(out *BucketPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketPolicyStatus.
func (in *BucketPolicyStatus) DeepCopy() *BucketPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(BucketPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketSpec) DeepCopyInto(out *BucketSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketSpec.
func (in *BucketSpec) DeepCopy() *BucketSpec {
	if in == nil {
		return nil
	}
	out := new(BucketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketStatus) DeepCopyInto(out *BucketStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketStatus.
func (in *BucketStatus) DeepCopy() *BucketStatus {
	if in == nil {
		return nil
	}
	out := new(BucketStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CorsRuleObservation) DeepCopyInto(out *CorsRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CorsRuleObservation.
func (in *CorsRuleObservation) DeepCopy() *CorsRuleObservation {
	if in == nil {
		return nil
	}
	out := new(CorsRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CorsRuleParameters) DeepCopyInto(out *CorsRuleParameters) {
	*out = *in
	if in.AllowedHeaders != nil {
		in, out := &in.AllowedHeaders, &out.AllowedHeaders
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedMethods != nil {
		in, out := &in.AllowedMethods, &out.AllowedMethods
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedOrigins != nil {
		in, out := &in.AllowedOrigins, &out.AllowedOrigins
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ExposeHeaders != nil {
		in, out := &in.ExposeHeaders, &out.ExposeHeaders
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.MaxAgeSeconds != nil {
		in, out := &in.MaxAgeSeconds, &out.MaxAgeSeconds
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CorsRuleParameters.
func (in *CorsRuleParameters) DeepCopy() *CorsRuleParameters {
	if in == nil {
		return nil
	}
	out := new(CorsRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpirationObservation) DeepCopyInto(out *ExpirationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpirationObservation.
func (in *ExpirationObservation) DeepCopy() *ExpirationObservation {
	if in == nil {
		return nil
	}
	out := new(ExpirationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpirationParameters) DeepCopyInto(out *ExpirationParameters) {
	*out = *in
	if in.Date != nil {
		in, out := &in.Date, &out.Date
		*out = new(string)
		**out = **in
	}
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = new(int64)
		**out = **in
	}
	if in.ExpiredObjectDeleteMarker != nil {
		in, out := &in.ExpiredObjectDeleteMarker, &out.ExpiredObjectDeleteMarker
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpirationParameters.
func (in *ExpirationParameters) DeepCopy() *ExpirationParameters {
	if in == nil {
		return nil
	}
	out := new(ExpirationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LifecycleRuleObservation) DeepCopyInto(out *LifecycleRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LifecycleRuleObservation.
func (in *LifecycleRuleObservation) DeepCopy() *LifecycleRuleObservation {
	if in == nil {
		return nil
	}
	out := new(LifecycleRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LifecycleRuleParameters) DeepCopyInto(out *LifecycleRuleParameters) {
	*out = *in
	if in.AbortIncompleteMultipartUploadDays != nil {
		in, out := &in.AbortIncompleteMultipartUploadDays, &out.AbortIncompleteMultipartUploadDays
		*out = new(int64)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Expiration != nil {
		in, out := &in.Expiration, &out.Expiration
		*out = make([]ExpirationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.NoncurrentVersionExpiration != nil {
		in, out := &in.NoncurrentVersionExpiration, &out.NoncurrentVersionExpiration
		*out = make([]NoncurrentVersionExpirationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LifecycleRuleParameters.
func (in *LifecycleRuleParameters) DeepCopy() *LifecycleRuleParameters {
	if in == nil {
		return nil
	}
	out := new(LifecycleRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingObservation) DeepCopyInto(out *LoggingObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingObservation.
func (in *LoggingObservation) DeepCopy() *LoggingObservation {
	if in == nil {
		return nil
	}
	out := new(LoggingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingParameters) DeepCopyInto(out *LoggingParameters) {
	*out = *in
	if in.TargetBucket != nil {
		in, out := &in.TargetBucket, &out.TargetBucket
		*out = new(string)
		**out = **in
	}
	if in.TargetPrefix != nil {
		in, out := &in.TargetPrefix, &out.TargetPrefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingParameters.
func (in *LoggingParameters) DeepCopy() *LoggingParameters {
	if in == nil {
		return nil
	}
	out := new(LoggingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoncurrentVersionExpirationObservation) DeepCopyInto(out *NoncurrentVersionExpirationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoncurrentVersionExpirationObservation.
func (in *NoncurrentVersionExpirationObservation) DeepCopy() *NoncurrentVersionExpirationObservation {
	if in == nil {
		return nil
	}
	out := new(NoncurrentVersionExpirationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoncurrentVersionExpirationParameters) DeepCopyInto(out *NoncurrentVersionExpirationParameters) {
	*out = *in
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoncurrentVersionExpirationParameters.
func (in *NoncurrentVersionExpirationParameters) DeepCopy() *NoncurrentVersionExpirationParameters {
	if in == nil {
		return nil
	}
	out := new(NoncurrentVersionExpirationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersioningObservation) DeepCopyInto(out *VersioningObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersioningObservation.
func (in *VersioningObservation) DeepCopy() *VersioningObservation {
	if in == nil {
		return nil
	}
	out := new(VersioningObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersioningParameters) DeepCopyInto(out *VersioningParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.MfaDelete != nil {
		in, out := &in.MfaDelete, &out.MfaDelete
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersioningParameters.
func (in *VersioningParameters) DeepCopy() *VersioningParameters {
	if in == nil {
		return nil
	}
	out := new(VersioningParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebsiteObservation) DeepCopyInto(out *WebsiteObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebsiteObservation.
func (in *WebsiteObservation) DeepCopy() *WebsiteObservation {
	if in == nil {
		return nil
	}
	out := new(WebsiteObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebsiteParameters) DeepCopyInto(out *WebsiteParameters) {
	*out = *in
	if in.ErrorDocument != nil {
		in, out := &in.ErrorDocument, &out.ErrorDocument
		*out = new(string)
		**out = **in
	}
	if in.IndexDocument != nil {
		in, out := &in.IndexDocument, &out.IndexDocument
		*out = new(string)
		**out = **in
	}
	if in.RedirectAllRequestsTo != nil {
		in, out := &in.RedirectAllRequestsTo, &out.RedirectAllRequestsTo
		*out = new(string)
		**out = **in
	}
	if in.RoutingRules != nil {
		in, out := &in.RoutingRules, &out.RoutingRules
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebsiteParameters.
func (in *WebsiteParameters) DeepCopy() *WebsiteParameters {
	if in == nil {
		return nil
	}
	out := new(WebsiteParameters)
	in.DeepCopyInto(out)
	return out
}