//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 The kubegems.io Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeCluster) DeepCopyInto(out *EdgeCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeCluster.
func (in *EdgeCluster) DeepCopy() *EdgeCluster {
	if in == nil {
		return nil
	}
	out := new(EdgeCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EdgeCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeClusterCridential) DeepCopyInto(out *EdgeClusterCridential) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeClusterCridential.
func (in *EdgeClusterCridential) DeepCopy() *EdgeClusterCridential {
	if in == nil {
		return nil
	}
	out := new(EdgeClusterCridential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EdgeClusterCridential) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeClusterCridentialList) DeepCopyInto(out *EdgeClusterCridentialList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EdgeCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeClusterCridentialList.
func (in *EdgeClusterCridentialList) DeepCopy() *EdgeClusterCridentialList {
	if in == nil {
		return nil
	}
	out := new(EdgeClusterCridentialList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EdgeClusterCridentialList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeClusterCridentialSpec) DeepCopyInto(out *EdgeClusterCridentialSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeClusterCridentialSpec.
func (in *EdgeClusterCridentialSpec) DeepCopy() *EdgeClusterCridentialSpec {
	if in == nil {
		return nil
	}
	out := new(EdgeClusterCridentialSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeClusterCridentialStatus) DeepCopyInto(out *EdgeClusterCridentialStatus) {
	*out = *in
	in.Expire.DeepCopyInto(&out.Expire)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeClusterCridentialStatus.
func (in *EdgeClusterCridentialStatus) DeepCopy() *EdgeClusterCridentialStatus {
	if in == nil {
		return nil
	}
	out := new(EdgeClusterCridentialStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeClusterList) DeepCopyInto(out *EdgeClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EdgeCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeClusterList.
func (in *EdgeClusterList) DeepCopy() *EdgeClusterList {
	if in == nil {
		return nil
	}
	out := new(EdgeClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EdgeClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeClusterSpec) DeepCopyInto(out *EdgeClusterSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeClusterSpec.
func (in *EdgeClusterSpec) DeepCopy() *EdgeClusterSpec {
	if in == nil {
		return nil
	}
	out := new(EdgeClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeClusterStatus) DeepCopyInto(out *EdgeClusterStatus) {
	*out = *in
	in.Register.DeepCopyInto(&out.Register)
	if in.Manufacture != nil {
		in, out := &in.Manufacture, &out.Manufacture
		*out = make(ManufactureStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeClusterStatus.
func (in *EdgeClusterStatus) DeepCopy() *EdgeClusterStatus {
	if in == nil {
		return nil
	}
	out := new(EdgeClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeClusterStatusRegister) DeepCopyInto(out *EdgeClusterStatusRegister) {
	*out = *in
	in.LastRegister.DeepCopyInto(&out.LastRegister)
	in.LastReporr.DeepCopyInto(&out.LastReporr)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeClusterStatusRegister.
func (in *EdgeClusterStatusRegister) DeepCopy() *EdgeClusterStatusRegister {
	if in == nil {
		return nil
	}
	out := new(EdgeClusterStatusRegister)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ManufactureStatus) DeepCopyInto(out *ManufactureStatus) {
	{
		in := &in
		*out = make(ManufactureStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManufactureStatus.
func (in ManufactureStatus) DeepCopy() ManufactureStatus {
	if in == nil {
		return nil
	}
	out := new(ManufactureStatus)
	in.DeepCopyInto(out)
	return *out
}
