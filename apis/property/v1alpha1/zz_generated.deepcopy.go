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
func (in *CertStatusObservation) DeepCopyInto(out *CertStatusObservation) {
	*out = *in
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.ProductionStatus != nil {
		in, out := &in.ProductionStatus, &out.ProductionStatus
		*out = new(string)
		**out = **in
	}
	if in.StagingStatus != nil {
		in, out := &in.StagingStatus, &out.StagingStatus
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertStatusObservation.
func (in *CertStatusObservation) DeepCopy() *CertStatusObservation {
	if in == nil {
		return nil
	}
	out := new(CertStatusObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertStatusParameters) DeepCopyInto(out *CertStatusParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertStatusParameters.
func (in *CertStatusParameters) DeepCopy() *CertStatusParameters {
	if in == nil {
		return nil
	}
	out := new(CertStatusParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostnamesObservation) DeepCopyInto(out *HostnamesObservation) {
	*out = *in
	if in.EdgeHostnameID != nil {
		in, out := &in.EdgeHostnameID, &out.EdgeHostnameID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostnamesObservation.
func (in *HostnamesObservation) DeepCopy() *HostnamesObservation {
	if in == nil {
		return nil
	}
	out := new(HostnamesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostnamesParameters) DeepCopyInto(out *HostnamesParameters) {
	*out = *in
	if in.CertProvisioningType != nil {
		in, out := &in.CertProvisioningType, &out.CertProvisioningType
		*out = new(string)
		**out = **in
	}
	if in.CertStatus != nil {
		in, out := &in.CertStatus, &out.CertStatus
		*out = make([]CertStatusParameters, len(*in))
		copy(*out, *in)
	}
	if in.CnameFrom != nil {
		in, out := &in.CnameFrom, &out.CnameFrom
		*out = new(string)
		**out = **in
	}
	if in.CnameTo != nil {
		in, out := &in.CnameTo, &out.CnameTo
		*out = new(string)
		**out = **in
	}
	if in.CnameType != nil {
		in, out := &in.CnameType, &out.CnameType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostnamesParameters.
func (in *HostnamesParameters) DeepCopy() *HostnamesParameters {
	if in == nil {
		return nil
	}
	out := new(HostnamesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginObservation) DeepCopyInto(out *OriginObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginObservation.
func (in *OriginObservation) DeepCopy() *OriginObservation {
	if in == nil {
		return nil
	}
	out := new(OriginObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginParameters) DeepCopyInto(out *OriginParameters) {
	*out = *in
	if in.CacheKeyHostname != nil {
		in, out := &in.CacheKeyHostname, &out.CacheKeyHostname
		*out = new(string)
		**out = **in
	}
	if in.Compress != nil {
		in, out := &in.Compress, &out.Compress
		*out = new(bool)
		**out = **in
	}
	if in.EnableTrueClientIP != nil {
		in, out := &in.EnableTrueClientIP, &out.EnableTrueClientIP
		*out = new(bool)
		**out = **in
	}
	if in.ForwardHostname != nil {
		in, out := &in.ForwardHostname, &out.ForwardHostname
		*out = new(string)
		**out = **in
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginParameters.
func (in *OriginParameters) DeepCopy() *OriginParameters {
	if in == nil {
		return nil
	}
	out := new(OriginParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Property) DeepCopyInto(out *Property) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Property.
func (in *Property) DeepCopy() *Property {
	if in == nil {
		return nil
	}
	out := new(Property)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Property) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PropertyList) DeepCopyInto(out *PropertyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Property, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PropertyList.
func (in *PropertyList) DeepCopy() *PropertyList {
	if in == nil {
		return nil
	}
	out := new(PropertyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PropertyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PropertyObservation) DeepCopyInto(out *PropertyObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LatestVersion != nil {
		in, out := &in.LatestVersion, &out.LatestVersion
		*out = new(int64)
		**out = **in
	}
	if in.ProductionVersion != nil {
		in, out := &in.ProductionVersion, &out.ProductionVersion
		*out = new(int64)
		**out = **in
	}
	if in.ReadVersion != nil {
		in, out := &in.ReadVersion, &out.ReadVersion
		*out = new(int64)
		**out = **in
	}
	if in.RuleErrors != nil {
		in, out := &in.RuleErrors, &out.RuleErrors
		*out = make([]RuleErrorsObservation, len(*in))
		copy(*out, *in)
	}
	if in.StagingVersion != nil {
		in, out := &in.StagingVersion, &out.StagingVersion
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PropertyObservation.
func (in *PropertyObservation) DeepCopy() *PropertyObservation {
	if in == nil {
		return nil
	}
	out := new(PropertyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PropertyParameters) DeepCopyInto(out *PropertyParameters) {
	*out = *in
	if in.Contact != nil {
		in, out := &in.Contact, &out.Contact
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Contract != nil {
		in, out := &in.Contract, &out.Contract
		*out = new(string)
		**out = **in
	}
	if in.ContractID != nil {
		in, out := &in.ContractID, &out.ContractID
		*out = new(string)
		**out = **in
	}
	if in.CpCode != nil {
		in, out := &in.CpCode, &out.CpCode
		*out = new(string)
		**out = **in
	}
	if in.Group != nil {
		in, out := &in.Group, &out.Group
		*out = new(string)
		**out = **in
	}
	if in.GroupID != nil {
		in, out := &in.GroupID, &out.GroupID
		*out = new(string)
		**out = **in
	}
	if in.Hostnames != nil {
		in, out := &in.Hostnames, &out.Hostnames
		*out = make([]HostnamesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IsSecure != nil {
		in, out := &in.IsSecure, &out.IsSecure
		*out = new(bool)
		**out = **in
	}
	if in.Origin != nil {
		in, out := &in.Origin, &out.Origin
		*out = make([]OriginParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Product != nil {
		in, out := &in.Product, &out.Product
		*out = new(string)
		**out = **in
	}
	if in.ProductID != nil {
		in, out := &in.ProductID, &out.ProductID
		*out = new(string)
		**out = **in
	}
	if in.RuleFormat != nil {
		in, out := &in.RuleFormat, &out.RuleFormat
		*out = new(string)
		**out = **in
	}
	if in.RuleWarnings != nil {
		in, out := &in.RuleWarnings, &out.RuleWarnings
		*out = make([]RuleWarningsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = new(string)
		**out = **in
	}
	if in.Variables != nil {
		in, out := &in.Variables, &out.Variables
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PropertyParameters.
func (in *PropertyParameters) DeepCopy() *PropertyParameters {
	if in == nil {
		return nil
	}
	out := new(PropertyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PropertySpec) DeepCopyInto(out *PropertySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PropertySpec.
func (in *PropertySpec) DeepCopy() *PropertySpec {
	if in == nil {
		return nil
	}
	out := new(PropertySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PropertyStatus) DeepCopyInto(out *PropertyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PropertyStatus.
func (in *PropertyStatus) DeepCopy() *PropertyStatus {
	if in == nil {
		return nil
	}
	out := new(PropertyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleErrorsObservation) DeepCopyInto(out *RuleErrorsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleErrorsObservation.
func (in *RuleErrorsObservation) DeepCopy() *RuleErrorsObservation {
	if in == nil {
		return nil
	}
	out := new(RuleErrorsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleErrorsParameters) DeepCopyInto(out *RuleErrorsParameters) {
	*out = *in
	if in.BehaviorName != nil {
		in, out := &in.BehaviorName, &out.BehaviorName
		*out = new(string)
		**out = **in
	}
	if in.Detail != nil {
		in, out := &in.Detail, &out.Detail
		*out = new(string)
		**out = **in
	}
	if in.ErrorLocation != nil {
		in, out := &in.ErrorLocation, &out.ErrorLocation
		*out = new(string)
		**out = **in
	}
	if in.Instance != nil {
		in, out := &in.Instance, &out.Instance
		*out = new(string)
		**out = **in
	}
	if in.StatusCode != nil {
		in, out := &in.StatusCode, &out.StatusCode
		*out = new(int64)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleErrorsParameters.
func (in *RuleErrorsParameters) DeepCopy() *RuleErrorsParameters {
	if in == nil {
		return nil
	}
	out := new(RuleErrorsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleWarningsObservation) DeepCopyInto(out *RuleWarningsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleWarningsObservation.
func (in *RuleWarningsObservation) DeepCopy() *RuleWarningsObservation {
	if in == nil {
		return nil
	}
	out := new(RuleWarningsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleWarningsParameters) DeepCopyInto(out *RuleWarningsParameters) {
	*out = *in
	if in.BehaviorName != nil {
		in, out := &in.BehaviorName, &out.BehaviorName
		*out = new(string)
		**out = **in
	}
	if in.Detail != nil {
		in, out := &in.Detail, &out.Detail
		*out = new(string)
		**out = **in
	}
	if in.ErrorLocation != nil {
		in, out := &in.ErrorLocation, &out.ErrorLocation
		*out = new(string)
		**out = **in
	}
	if in.Instance != nil {
		in, out := &in.Instance, &out.Instance
		*out = new(string)
		**out = **in
	}
	if in.StatusCode != nil {
		in, out := &in.StatusCode, &out.StatusCode
		*out = new(int64)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleWarningsParameters.
func (in *RuleWarningsParameters) DeepCopy() *RuleWarningsParameters {
	if in == nil {
		return nil
	}
	out := new(RuleWarningsParameters)
	in.DeepCopyInto(out)
	return out
}
