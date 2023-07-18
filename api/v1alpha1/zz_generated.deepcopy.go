//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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
func (in *ApplicationDisruptionBudget) DeepCopyInto(out *ApplicationDisruptionBudget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationDisruptionBudget.
func (in *ApplicationDisruptionBudget) DeepCopy() *ApplicationDisruptionBudget {
	if in == nil {
		return nil
	}
	out := new(ApplicationDisruptionBudget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationDisruptionBudget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationDisruptionBudgetList) DeepCopyInto(out *ApplicationDisruptionBudgetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApplicationDisruptionBudget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationDisruptionBudgetList.
func (in *ApplicationDisruptionBudgetList) DeepCopy() *ApplicationDisruptionBudgetList {
	if in == nil {
		return nil
	}
	out := new(ApplicationDisruptionBudgetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationDisruptionBudgetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationDisruptionBudgetSpec) DeepCopyInto(out *ApplicationDisruptionBudgetSpec) {
	*out = *in
	in.PodSelector.DeepCopyInto(&out.PodSelector)
	in.PVCSelector.DeepCopyInto(&out.PVCSelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationDisruptionBudgetSpec.
func (in *ApplicationDisruptionBudgetSpec) DeepCopy() *ApplicationDisruptionBudgetSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationDisruptionBudgetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DisruptionBudgetStatus) DeepCopyInto(out *DisruptionBudgetStatus) {
	*out = *in
	if in.WatchedNodes != nil {
		in, out := &in.WatchedNodes, &out.WatchedNodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DisruptionBudgetStatus.
func (in *DisruptionBudgetStatus) DeepCopy() *DisruptionBudgetStatus {
	if in == nil {
		return nil
	}
	out := new(DisruptionBudgetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedName) DeepCopyInto(out *NamespacedName) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedName.
func (in *NamespacedName) DeepCopy() *NamespacedName {
	if in == nil {
		return nil
	}
	out := new(NamespacedName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeDisruption) DeepCopyInto(out *NodeDisruption) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeDisruption.
func (in *NodeDisruption) DeepCopy() *NodeDisruption {
	if in == nil {
		return nil
	}
	out := new(NodeDisruption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeDisruption) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeDisruptionBudget) DeepCopyInto(out *NodeDisruptionBudget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeDisruptionBudget.
func (in *NodeDisruptionBudget) DeepCopy() *NodeDisruptionBudget {
	if in == nil {
		return nil
	}
	out := new(NodeDisruptionBudget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeDisruptionBudget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeDisruptionBudgetList) DeepCopyInto(out *NodeDisruptionBudgetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeDisruptionBudget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeDisruptionBudgetList.
func (in *NodeDisruptionBudgetList) DeepCopy() *NodeDisruptionBudgetList {
	if in == nil {
		return nil
	}
	out := new(NodeDisruptionBudgetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeDisruptionBudgetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeDisruptionBudgetSpec) DeepCopyInto(out *NodeDisruptionBudgetSpec) {
	*out = *in
	in.NodeSelector.DeepCopyInto(&out.NodeSelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeDisruptionBudgetSpec.
func (in *NodeDisruptionBudgetSpec) DeepCopy() *NodeDisruptionBudgetSpec {
	if in == nil {
		return nil
	}
	out := new(NodeDisruptionBudgetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeDisruptionList) DeepCopyInto(out *NodeDisruptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeDisruption, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeDisruptionList.
func (in *NodeDisruptionList) DeepCopy() *NodeDisruptionList {
	if in == nil {
		return nil
	}
	out := new(NodeDisruptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeDisruptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeDisruptionSpec) DeepCopyInto(out *NodeDisruptionSpec) {
	*out = *in
	in.NodeSelector.DeepCopyInto(&out.NodeSelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeDisruptionSpec.
func (in *NodeDisruptionSpec) DeepCopy() *NodeDisruptionSpec {
	if in == nil {
		return nil
	}
	out := new(NodeDisruptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeDisruptionStatus) DeepCopyInto(out *NodeDisruptionStatus) {
	*out = *in
	if in.DisruptedADB != nil {
		in, out := &in.DisruptedADB, &out.DisruptedADB
		*out = make([]NamespacedName, len(*in))
		copy(*out, *in)
	}
	if in.DisruptedNDB != nil {
		in, out := &in.DisruptedNDB, &out.DisruptedNDB
		*out = make([]NamespacedName, len(*in))
		copy(*out, *in)
	}
	if in.DisruptedNodes != nil {
		in, out := &in.DisruptedNodes, &out.DisruptedNodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeDisruptionStatus.
func (in *NodeDisruptionStatus) DeepCopy() *NodeDisruptionStatus {
	if in == nil {
		return nil
	}
	out := new(NodeDisruptionStatus)
	in.DeepCopyInto(out)
	return out
}
