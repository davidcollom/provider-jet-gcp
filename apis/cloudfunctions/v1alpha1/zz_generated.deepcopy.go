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
func (in *ConditionObservation) DeepCopyInto(out *ConditionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionObservation.
func (in *ConditionObservation) DeepCopy() *ConditionObservation {
	if in == nil {
		return nil
	}
	out := new(ConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionParameters) DeepCopyInto(out *ConditionParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionParameters.
func (in *ConditionParameters) DeepCopy() *ConditionParameters {
	if in == nil {
		return nil
	}
	out := new(ConditionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventTriggerObservation) DeepCopyInto(out *EventTriggerObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventTriggerObservation.
func (in *EventTriggerObservation) DeepCopy() *EventTriggerObservation {
	if in == nil {
		return nil
	}
	out := new(EventTriggerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventTriggerParameters) DeepCopyInto(out *EventTriggerParameters) {
	*out = *in
	if in.EventType != nil {
		in, out := &in.EventType, &out.EventType
		*out = new(string)
		**out = **in
	}
	if in.FailurePolicy != nil {
		in, out := &in.FailurePolicy, &out.FailurePolicy
		*out = make([]FailurePolicyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventTriggerParameters.
func (in *EventTriggerParameters) DeepCopy() *EventTriggerParameters {
	if in == nil {
		return nil
	}
	out := new(EventTriggerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FailurePolicyObservation) DeepCopyInto(out *FailurePolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FailurePolicyObservation.
func (in *FailurePolicyObservation) DeepCopy() *FailurePolicyObservation {
	if in == nil {
		return nil
	}
	out := new(FailurePolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FailurePolicyParameters) DeepCopyInto(out *FailurePolicyParameters) {
	*out = *in
	if in.Retry != nil {
		in, out := &in.Retry, &out.Retry
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FailurePolicyParameters.
func (in *FailurePolicyParameters) DeepCopy() *FailurePolicyParameters {
	if in == nil {
		return nil
	}
	out := new(FailurePolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Function) DeepCopyInto(out *Function) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Function.
func (in *Function) DeepCopy() *Function {
	if in == nil {
		return nil
	}
	out := new(Function)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Function) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionIAMBinding) DeepCopyInto(out *FunctionIAMBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionIAMBinding.
func (in *FunctionIAMBinding) DeepCopy() *FunctionIAMBinding {
	if in == nil {
		return nil
	}
	out := new(FunctionIAMBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionIAMBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionIAMBindingList) DeepCopyInto(out *FunctionIAMBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FunctionIAMBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionIAMBindingList.
func (in *FunctionIAMBindingList) DeepCopy() *FunctionIAMBindingList {
	if in == nil {
		return nil
	}
	out := new(FunctionIAMBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionIAMBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionIAMBindingObservation) DeepCopyInto(out *FunctionIAMBindingObservation) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionIAMBindingObservation.
func (in *FunctionIAMBindingObservation) DeepCopy() *FunctionIAMBindingObservation {
	if in == nil {
		return nil
	}
	out := new(FunctionIAMBindingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionIAMBindingParameters) DeepCopyInto(out *FunctionIAMBindingParameters) {
	*out = *in
	if in.CloudFunction != nil {
		in, out := &in.CloudFunction, &out.CloudFunction
		*out = new(string)
		**out = **in
	}
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = make([]ConditionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionIAMBindingParameters.
func (in *FunctionIAMBindingParameters) DeepCopy() *FunctionIAMBindingParameters {
	if in == nil {
		return nil
	}
	out := new(FunctionIAMBindingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionIAMBindingSpec) DeepCopyInto(out *FunctionIAMBindingSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionIAMBindingSpec.
func (in *FunctionIAMBindingSpec) DeepCopy() *FunctionIAMBindingSpec {
	if in == nil {
		return nil
	}
	out := new(FunctionIAMBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionIAMBindingStatus) DeepCopyInto(out *FunctionIAMBindingStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionIAMBindingStatus.
func (in *FunctionIAMBindingStatus) DeepCopy() *FunctionIAMBindingStatus {
	if in == nil {
		return nil
	}
	out := new(FunctionIAMBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionIAMMember) DeepCopyInto(out *FunctionIAMMember) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionIAMMember.
func (in *FunctionIAMMember) DeepCopy() *FunctionIAMMember {
	if in == nil {
		return nil
	}
	out := new(FunctionIAMMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionIAMMember) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionIAMMemberConditionObservation) DeepCopyInto(out *FunctionIAMMemberConditionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionIAMMemberConditionObservation.
func (in *FunctionIAMMemberConditionObservation) DeepCopy() *FunctionIAMMemberConditionObservation {
	if in == nil {
		return nil
	}
	out := new(FunctionIAMMemberConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionIAMMemberConditionParameters) DeepCopyInto(out *FunctionIAMMemberConditionParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionIAMMemberConditionParameters.
func (in *FunctionIAMMemberConditionParameters) DeepCopy() *FunctionIAMMemberConditionParameters {
	if in == nil {
		return nil
	}
	out := new(FunctionIAMMemberConditionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionIAMMemberList) DeepCopyInto(out *FunctionIAMMemberList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FunctionIAMMember, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionIAMMemberList.
func (in *FunctionIAMMemberList) DeepCopy() *FunctionIAMMemberList {
	if in == nil {
		return nil
	}
	out := new(FunctionIAMMemberList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionIAMMemberList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionIAMMemberObservation) DeepCopyInto(out *FunctionIAMMemberObservation) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionIAMMemberObservation.
func (in *FunctionIAMMemberObservation) DeepCopy() *FunctionIAMMemberObservation {
	if in == nil {
		return nil
	}
	out := new(FunctionIAMMemberObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionIAMMemberParameters) DeepCopyInto(out *FunctionIAMMemberParameters) {
	*out = *in
	if in.CloudFunction != nil {
		in, out := &in.CloudFunction, &out.CloudFunction
		*out = new(string)
		**out = **in
	}
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = make([]FunctionIAMMemberConditionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionIAMMemberParameters.
func (in *FunctionIAMMemberParameters) DeepCopy() *FunctionIAMMemberParameters {
	if in == nil {
		return nil
	}
	out := new(FunctionIAMMemberParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionIAMMemberSpec) DeepCopyInto(out *FunctionIAMMemberSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionIAMMemberSpec.
func (in *FunctionIAMMemberSpec) DeepCopy() *FunctionIAMMemberSpec {
	if in == nil {
		return nil
	}
	out := new(FunctionIAMMemberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionIAMMemberStatus) DeepCopyInto(out *FunctionIAMMemberStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionIAMMemberStatus.
func (in *FunctionIAMMemberStatus) DeepCopy() *FunctionIAMMemberStatus {
	if in == nil {
		return nil
	}
	out := new(FunctionIAMMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionIAMPolicy) DeepCopyInto(out *FunctionIAMPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionIAMPolicy.
func (in *FunctionIAMPolicy) DeepCopy() *FunctionIAMPolicy {
	if in == nil {
		return nil
	}
	out := new(FunctionIAMPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionIAMPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionIAMPolicyList) DeepCopyInto(out *FunctionIAMPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FunctionIAMPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionIAMPolicyList.
func (in *FunctionIAMPolicyList) DeepCopy() *FunctionIAMPolicyList {
	if in == nil {
		return nil
	}
	out := new(FunctionIAMPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionIAMPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionIAMPolicyObservation) DeepCopyInto(out *FunctionIAMPolicyObservation) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionIAMPolicyObservation.
func (in *FunctionIAMPolicyObservation) DeepCopy() *FunctionIAMPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(FunctionIAMPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionIAMPolicyParameters) DeepCopyInto(out *FunctionIAMPolicyParameters) {
	*out = *in
	if in.CloudFunction != nil {
		in, out := &in.CloudFunction, &out.CloudFunction
		*out = new(string)
		**out = **in
	}
	if in.PolicyData != nil {
		in, out := &in.PolicyData, &out.PolicyData
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionIAMPolicyParameters.
func (in *FunctionIAMPolicyParameters) DeepCopy() *FunctionIAMPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(FunctionIAMPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionIAMPolicySpec) DeepCopyInto(out *FunctionIAMPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionIAMPolicySpec.
func (in *FunctionIAMPolicySpec) DeepCopy() *FunctionIAMPolicySpec {
	if in == nil {
		return nil
	}
	out := new(FunctionIAMPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionIAMPolicyStatus) DeepCopyInto(out *FunctionIAMPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionIAMPolicyStatus.
func (in *FunctionIAMPolicyStatus) DeepCopy() *FunctionIAMPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(FunctionIAMPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionList) DeepCopyInto(out *FunctionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Function, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionList.
func (in *FunctionList) DeepCopy() *FunctionList {
	if in == nil {
		return nil
	}
	out := new(FunctionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionObservation) DeepCopyInto(out *FunctionObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionObservation.
func (in *FunctionObservation) DeepCopy() *FunctionObservation {
	if in == nil {
		return nil
	}
	out := new(FunctionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionParameters) DeepCopyInto(out *FunctionParameters) {
	*out = *in
	if in.AvailableMemoryMb != nil {
		in, out := &in.AvailableMemoryMb, &out.AvailableMemoryMb
		*out = new(int64)
		**out = **in
	}
	if in.BuildEnvironmentVariables != nil {
		in, out := &in.BuildEnvironmentVariables, &out.BuildEnvironmentVariables
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EntryPoint != nil {
		in, out := &in.EntryPoint, &out.EntryPoint
		*out = new(string)
		**out = **in
	}
	if in.EnvironmentVariables != nil {
		in, out := &in.EnvironmentVariables, &out.EnvironmentVariables
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EventTrigger != nil {
		in, out := &in.EventTrigger, &out.EventTrigger
		*out = make([]EventTriggerParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HTTPSTriggerURL != nil {
		in, out := &in.HTTPSTriggerURL, &out.HTTPSTriggerURL
		*out = new(string)
		**out = **in
	}
	if in.IngressSettings != nil {
		in, out := &in.IngressSettings, &out.IngressSettings
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MaxInstances != nil {
		in, out := &in.MaxInstances, &out.MaxInstances
		*out = new(int64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Runtime != nil {
		in, out := &in.Runtime, &out.Runtime
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountEmail != nil {
		in, out := &in.ServiceAccountEmail, &out.ServiceAccountEmail
		*out = new(string)
		**out = **in
	}
	if in.SourceArchiveBucket != nil {
		in, out := &in.SourceArchiveBucket, &out.SourceArchiveBucket
		*out = new(string)
		**out = **in
	}
	if in.SourceArchiveObject != nil {
		in, out := &in.SourceArchiveObject, &out.SourceArchiveObject
		*out = new(string)
		**out = **in
	}
	if in.SourceRepository != nil {
		in, out := &in.SourceRepository, &out.SourceRepository
		*out = make([]SourceRepositoryParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int64)
		**out = **in
	}
	if in.TriggerHTTP != nil {
		in, out := &in.TriggerHTTP, &out.TriggerHTTP
		*out = new(bool)
		**out = **in
	}
	if in.VPCConnector != nil {
		in, out := &in.VPCConnector, &out.VPCConnector
		*out = new(string)
		**out = **in
	}
	if in.VPCConnectorEgressSettings != nil {
		in, out := &in.VPCConnectorEgressSettings, &out.VPCConnectorEgressSettings
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionParameters.
func (in *FunctionParameters) DeepCopy() *FunctionParameters {
	if in == nil {
		return nil
	}
	out := new(FunctionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionSpec) DeepCopyInto(out *FunctionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionSpec.
func (in *FunctionSpec) DeepCopy() *FunctionSpec {
	if in == nil {
		return nil
	}
	out := new(FunctionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionStatus) DeepCopyInto(out *FunctionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionStatus.
func (in *FunctionStatus) DeepCopy() *FunctionStatus {
	if in == nil {
		return nil
	}
	out := new(FunctionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceRepositoryObservation) DeepCopyInto(out *SourceRepositoryObservation) {
	*out = *in
	if in.DeployedURL != nil {
		in, out := &in.DeployedURL, &out.DeployedURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceRepositoryObservation.
func (in *SourceRepositoryObservation) DeepCopy() *SourceRepositoryObservation {
	if in == nil {
		return nil
	}
	out := new(SourceRepositoryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceRepositoryParameters) DeepCopyInto(out *SourceRepositoryParameters) {
	*out = *in
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceRepositoryParameters.
func (in *SourceRepositoryParameters) DeepCopy() *SourceRepositoryParameters {
	if in == nil {
		return nil
	}
	out := new(SourceRepositoryParameters)
	in.DeepCopyInto(out)
	return out
}