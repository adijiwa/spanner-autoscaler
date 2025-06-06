//go:build !ignore_autogenerated

/*
Copyright 2023.

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
func (in *ImpersonateConfig) DeepCopyInto(out *ImpersonateConfig) {
	*out = *in
	if in.Delegates != nil {
		in, out := &in.Delegates, &out.Delegates
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImpersonateConfig.
func (in *ImpersonateConfig) DeepCopy() *ImpersonateConfig {
	if in == nil {
		return nil
	}
	out := new(ImpersonateConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleTargetRef) DeepCopyInto(out *ScaleTargetRef) {
	*out = *in
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleTargetRef.
func (in *ScaleTargetRef) DeepCopy() *ScaleTargetRef {
	if in == nil {
		return nil
	}
	out := new(ScaleTargetRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountSecretRef) DeepCopyInto(out *ServiceAccountSecretRef) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountSecretRef.
func (in *ServiceAccountSecretRef) DeepCopy() *ServiceAccountSecretRef {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountSecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpannerAutoscaler) DeepCopyInto(out *SpannerAutoscaler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpannerAutoscaler.
func (in *SpannerAutoscaler) DeepCopy() *SpannerAutoscaler {
	if in == nil {
		return nil
	}
	out := new(SpannerAutoscaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpannerAutoscaler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpannerAutoscalerList) DeepCopyInto(out *SpannerAutoscalerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SpannerAutoscaler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpannerAutoscalerList.
func (in *SpannerAutoscalerList) DeepCopy() *SpannerAutoscalerList {
	if in == nil {
		return nil
	}
	out := new(SpannerAutoscalerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpannerAutoscalerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpannerAutoscalerSpec) DeepCopyInto(out *SpannerAutoscalerSpec) {
	*out = *in
	in.ScaleTargetRef.DeepCopyInto(&out.ScaleTargetRef)
	if in.ServiceAccountSecretRef != nil {
		in, out := &in.ServiceAccountSecretRef, &out.ServiceAccountSecretRef
		*out = new(ServiceAccountSecretRef)
		(*in).DeepCopyInto(*out)
	}
	if in.ImpersonateConfig != nil {
		in, out := &in.ImpersonateConfig, &out.ImpersonateConfig
		*out = new(ImpersonateConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.MinNodes != nil {
		in, out := &in.MinNodes, &out.MinNodes
		*out = new(int32)
		**out = **in
	}
	if in.MaxNodes != nil {
		in, out := &in.MaxNodes, &out.MaxNodes
		*out = new(int32)
		**out = **in
	}
	if in.MinProcessingUnits != nil {
		in, out := &in.MinProcessingUnits, &out.MinProcessingUnits
		*out = new(int32)
		**out = **in
	}
	if in.MaxProcessingUnits != nil {
		in, out := &in.MaxProcessingUnits, &out.MaxProcessingUnits
		*out = new(int32)
		**out = **in
	}
	if in.MaxScaleDownNodes != nil {
		in, out := &in.MaxScaleDownNodes, &out.MaxScaleDownNodes
		*out = new(int32)
		**out = **in
	}
	in.TargetCPUUtilization.DeepCopyInto(&out.TargetCPUUtilization)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpannerAutoscalerSpec.
func (in *SpannerAutoscalerSpec) DeepCopy() *SpannerAutoscalerSpec {
	if in == nil {
		return nil
	}
	out := new(SpannerAutoscalerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpannerAutoscalerStatus) DeepCopyInto(out *SpannerAutoscalerStatus) {
	*out = *in
	if in.LastScaleTime != nil {
		in, out := &in.LastScaleTime, &out.LastScaleTime
		*out = (*in).DeepCopy()
	}
	if in.LastSyncTime != nil {
		in, out := &in.LastSyncTime, &out.LastSyncTime
		*out = (*in).DeepCopy()
	}
	if in.CurrentNodes != nil {
		in, out := &in.CurrentNodes, &out.CurrentNodes
		*out = new(int32)
		**out = **in
	}
	if in.CurrentProcessingUnits != nil {
		in, out := &in.CurrentProcessingUnits, &out.CurrentProcessingUnits
		*out = new(int32)
		**out = **in
	}
	if in.DesiredNodes != nil {
		in, out := &in.DesiredNodes, &out.DesiredNodes
		*out = new(int32)
		**out = **in
	}
	if in.DesiredProcessingUnits != nil {
		in, out := &in.DesiredProcessingUnits, &out.DesiredProcessingUnits
		*out = new(int32)
		**out = **in
	}
	if in.CurrentHighPriorityCPUUtilization != nil {
		in, out := &in.CurrentHighPriorityCPUUtilization, &out.CurrentHighPriorityCPUUtilization
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpannerAutoscalerStatus.
func (in *SpannerAutoscalerStatus) DeepCopy() *SpannerAutoscalerStatus {
	if in == nil {
		return nil
	}
	out := new(SpannerAutoscalerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetCPUUtilization) DeepCopyInto(out *TargetCPUUtilization) {
	*out = *in
	if in.HighPriority != nil {
		in, out := &in.HighPriority, &out.HighPriority
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetCPUUtilization.
func (in *TargetCPUUtilization) DeepCopy() *TargetCPUUtilization {
	if in == nil {
		return nil
	}
	out := new(TargetCPUUtilization)
	in.DeepCopyInto(out)
	return out
}
