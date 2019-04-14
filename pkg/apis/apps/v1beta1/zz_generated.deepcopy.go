// +build !ignore_autogenerated

/*
Copyright 2019 FoundationDB project authors.

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
// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionString) DeepCopyInto(out *ConnectionString) {
	*out = *in
	if in.Coordinators != nil {
		in, out := &in.Coordinators, &out.Coordinators
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionString.
func (in *ConnectionString) DeepCopy() *ConnectionString {
	if in == nil {
		return nil
	}
	out := new(ConnectionString)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBCluster) DeepCopyInto(out *FoundationDBCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBCluster.
func (in *FoundationDBCluster) DeepCopy() *FoundationDBCluster {
	if in == nil {
		return nil
	}
	out := new(FoundationDBCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FoundationDBCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBClusterList) DeepCopyInto(out *FoundationDBClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FoundationDBCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBClusterList.
func (in *FoundationDBClusterList) DeepCopy() *FoundationDBClusterList {
	if in == nil {
		return nil
	}
	out := new(FoundationDBClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FoundationDBClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBClusterSpec) DeepCopyInto(out *FoundationDBClusterSpec) {
	*out = *in
	out.RoleCounts = in.RoleCounts
	out.ProcessCounts = in.ProcessCounts
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
	if in.PendingRemovals != nil {
		in, out := &in.PendingRemovals, &out.PendingRemovals
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CustomParameters != nil {
		in, out := &in.CustomParameters, &out.CustomParameters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBClusterSpec.
func (in *FoundationDBClusterSpec) DeepCopy() *FoundationDBClusterSpec {
	if in == nil {
		return nil
	}
	out := new(FoundationDBClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBClusterStatus) DeepCopyInto(out *FoundationDBClusterStatus) {
	*out = *in
	out.ProcessCounts = in.ProcessCounts
	if in.IncorrectProcesses != nil {
		in, out := &in.IncorrectProcesses, &out.IncorrectProcesses
		*out = make(map[string]int64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MissingProcesses != nil {
		in, out := &in.MissingProcesses, &out.MissingProcesses
		*out = make(map[string]int64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBClusterStatus.
func (in *FoundationDBClusterStatus) DeepCopy() *FoundationDBClusterStatus {
	if in == nil {
		return nil
	}
	out := new(FoundationDBClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBStatus) DeepCopyInto(out *FoundationDBStatus) {
	*out = *in
	in.Client.DeepCopyInto(&out.Client)
	in.Cluster.DeepCopyInto(&out.Cluster)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBStatus.
func (in *FoundationDBStatus) DeepCopy() *FoundationDBStatus {
	if in == nil {
		return nil
	}
	out := new(FoundationDBStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBStatusClientInfo) DeepCopyInto(out *FoundationDBStatusClientInfo) {
	*out = *in
	in.Coordinators.DeepCopyInto(&out.Coordinators)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBStatusClientInfo.
func (in *FoundationDBStatusClientInfo) DeepCopy() *FoundationDBStatusClientInfo {
	if in == nil {
		return nil
	}
	out := new(FoundationDBStatusClientInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBStatusClusterInfo) DeepCopyInto(out *FoundationDBStatusClusterInfo) {
	*out = *in
	if in.Processes != nil {
		in, out := &in.Processes, &out.Processes
		*out = make(map[string]FoundationDBStatusProcessInfo, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBStatusClusterInfo.
func (in *FoundationDBStatusClusterInfo) DeepCopy() *FoundationDBStatusClusterInfo {
	if in == nil {
		return nil
	}
	out := new(FoundationDBStatusClusterInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBStatusCoordinator) DeepCopyInto(out *FoundationDBStatusCoordinator) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBStatusCoordinator.
func (in *FoundationDBStatusCoordinator) DeepCopy() *FoundationDBStatusCoordinator {
	if in == nil {
		return nil
	}
	out := new(FoundationDBStatusCoordinator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBStatusCoordinatorInfo) DeepCopyInto(out *FoundationDBStatusCoordinatorInfo) {
	*out = *in
	if in.Coordinators != nil {
		in, out := &in.Coordinators, &out.Coordinators
		*out = make([]FoundationDBStatusCoordinator, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBStatusCoordinatorInfo.
func (in *FoundationDBStatusCoordinatorInfo) DeepCopy() *FoundationDBStatusCoordinatorInfo {
	if in == nil {
		return nil
	}
	out := new(FoundationDBStatusCoordinatorInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBStatusProcessInfo) DeepCopyInto(out *FoundationDBStatusProcessInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBStatusProcessInfo.
func (in *FoundationDBStatusProcessInfo) DeepCopy() *FoundationDBStatusProcessInfo {
	if in == nil {
		return nil
	}
	out := new(FoundationDBStatusProcessInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessCounts) DeepCopyInto(out *ProcessCounts) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessCounts.
func (in *ProcessCounts) DeepCopy() *ProcessCounts {
	if in == nil {
		return nil
	}
	out := new(ProcessCounts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleCounts) DeepCopyInto(out *RoleCounts) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleCounts.
func (in *RoleCounts) DeepCopy() *RoleCounts {
	if in == nil {
		return nil
	}
	out := new(RoleCounts)
	in.DeepCopyInto(out)
	return out
}
