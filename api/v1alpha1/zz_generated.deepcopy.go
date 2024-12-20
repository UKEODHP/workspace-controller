//go:build !ignore_autogenerated

/*
Copyright 2024 Telespazio UK.

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
func (in *AWSRoleStatus) DeepCopyInto(out *AWSRoleStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSRoleStatus.
func (in *AWSRoleStatus) DeepCopy() *AWSRoleStatus {
	if in == nil {
		return nil
	}
	out := new(AWSRoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSSpec) DeepCopyInto(out *AWSSpec) {
	*out = *in
	in.EFS.DeepCopyInto(&out.EFS)
	in.S3.DeepCopyInto(&out.S3)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSSpec.
func (in *AWSSpec) DeepCopy() *AWSSpec {
	if in == nil {
		return nil
	}
	out := new(AWSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSStatus) DeepCopyInto(out *AWSStatus) {
	*out = *in
	out.Role = in.Role
	in.EFS.DeepCopyInto(&out.EFS)
	in.S3.DeepCopyInto(&out.S3)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSStatus.
func (in *AWSStatus) DeepCopy() *AWSStatus {
	if in == nil {
		return nil
	}
	out := new(AWSStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthorizationSpec) DeepCopyInto(out *AuthorizationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizationSpec.
func (in *AuthorizationSpec) DeepCopy() *AuthorizationSpec {
	if in == nil {
		return nil
	}
	out := new(AuthorizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EFSAccess) DeepCopyInto(out *EFSAccess) {
	*out = *in
	out.User = in.User
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EFSAccess.
func (in *EFSAccess) DeepCopy() *EFSAccess {
	if in == nil {
		return nil
	}
	out := new(EFSAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EFSAccessStatus) DeepCopyInto(out *EFSAccessStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EFSAccessStatus.
func (in *EFSAccessStatus) DeepCopy() *EFSAccessStatus {
	if in == nil {
		return nil
	}
	out := new(EFSAccessStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EFSSpec) DeepCopyInto(out *EFSSpec) {
	*out = *in
	if in.AccessPoints != nil {
		in, out := &in.AccessPoints, &out.AccessPoints
		*out = make([]EFSAccess, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EFSSpec.
func (in *EFSSpec) DeepCopy() *EFSSpec {
	if in == nil {
		return nil
	}
	out := new(EFSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EFSStatus) DeepCopyInto(out *EFSStatus) {
	*out = *in
	if in.AccessPoints != nil {
		in, out := &in.AccessPoints, &out.AccessPoints
		*out = make([]EFSAccessStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EFSStatus.
func (in *EFSStatus) DeepCopy() *EFSStatus {
	if in == nil {
		return nil
	}
	out := new(EFSStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PVCSpec) DeepCopyInto(out *PVCSpec) {
	*out = *in
	in.PVSpec.DeepCopyInto(&out.PVSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PVCSpec.
func (in *PVCSpec) DeepCopy() *PVCSpec {
	if in == nil {
		return nil
	}
	out := new(PVCSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PVSpec) DeepCopyInto(out *PVSpec) {
	*out = *in
	if in.VolumeSource != nil {
		in, out := &in.VolumeSource, &out.VolumeSource
		*out = new(VolumeSource)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PVSpec.
func (in *PVSpec) DeepCopy() *PVSpec {
	if in == nil {
		return nil
	}
	out := new(PVSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Bucket) DeepCopyInto(out *S3Bucket) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Bucket.
func (in *S3Bucket) DeepCopy() *S3Bucket {
	if in == nil {
		return nil
	}
	out := new(S3Bucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketStatus) DeepCopyInto(out *S3BucketStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketStatus.
func (in *S3BucketStatus) DeepCopy() *S3BucketStatus {
	if in == nil {
		return nil
	}
	out := new(S3BucketStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Spec) DeepCopyInto(out *S3Spec) {
	*out = *in
	if in.Buckets != nil {
		in, out := &in.Buckets, &out.Buckets
		*out = make([]S3Bucket, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Spec.
func (in *S3Spec) DeepCopy() *S3Spec {
	if in == nil {
		return nil
	}
	out := new(S3Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Status) DeepCopyInto(out *S3Status) {
	*out = *in
	if in.Buckets != nil {
		in, out := &in.Buckets, &out.Buckets
		*out = make([]S3BucketStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Status.
func (in *S3Status) DeepCopy() *S3Status {
	if in == nil {
		return nil
	}
	out := new(S3Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountSpec) DeepCopyInto(out *ServiceAccountSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountSpec.
func (in *ServiceAccountSpec) DeepCopy() *ServiceAccountSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageSpec) DeepCopyInto(out *StorageSpec) {
	*out = *in
	if in.PersistentVolumes != nil {
		in, out := &in.PersistentVolumes, &out.PersistentVolumes
		*out = make([]PVSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PersistentVolumeClaims != nil {
		in, out := &in.PersistentVolumeClaims, &out.PersistentVolumeClaims
		*out = make([]PVCSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageSpec.
func (in *StorageSpec) DeepCopy() *StorageSpec {
	if in == nil {
		return nil
	}
	out := new(StorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *User) DeepCopyInto(out *User) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new User.
func (in *User) DeepCopy() *User {
	if in == nil {
		return nil
	}
	out := new(User)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSource) DeepCopyInto(out *VolumeSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeSource.
func (in *VolumeSource) DeepCopy() *VolumeSource {
	if in == nil {
		return nil
	}
	out := new(VolumeSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workspace) DeepCopyInto(out *Workspace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workspace.
func (in *Workspace) DeepCopy() *Workspace {
	if in == nil {
		return nil
	}
	out := new(Workspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Workspace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceList) DeepCopyInto(out *WorkspaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Workspace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceList.
func (in *WorkspaceList) DeepCopy() *WorkspaceList {
	if in == nil {
		return nil
	}
	out := new(WorkspaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceSpec) DeepCopyInto(out *WorkspaceSpec) {
	*out = *in
	in.AWS.DeepCopyInto(&out.AWS)
	in.ServiceAccount.DeepCopyInto(&out.ServiceAccount)
	in.Storage.DeepCopyInto(&out.Storage)
	out.Authorization = in.Authorization
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceSpec.
func (in *WorkspaceSpec) DeepCopy() *WorkspaceSpec {
	if in == nil {
		return nil
	}
	out := new(WorkspaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceStatus) DeepCopyInto(out *WorkspaceStatus) {
	*out = *in
	in.AWS.DeepCopyInto(&out.AWS)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceStatus.
func (in *WorkspaceStatus) DeepCopy() *WorkspaceStatus {
	if in == nil {
		return nil
	}
	out := new(WorkspaceStatus)
	in.DeepCopyInto(out)
	return out
}
