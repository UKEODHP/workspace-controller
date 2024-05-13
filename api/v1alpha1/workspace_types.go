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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type User struct {
	UID int64 `json:"uid,omitempty"`
	GID int64 `json:"gid,omitempty"`
}

type StorageSpec struct {
	// Persistent volume name
	PVName string `json:"pvName,omitempty"`
	// Persistent volume claim name
	PVCName string `json:"pvcName,omitempty"`
	// Kubernetes storage class to use
	StorageClass string `json:"storageClass,omitempty"`
	// Size of the storage
	Size string `json:"size,omitempty"`
	// User
	User User `json:"user,omitempty"`
	// Define the EFS storage
	AWSEFS EFSSpec `json:"awsEFS,omitempty"`
	AWSS3  S3Spec  `json:"awsS3,omitempty"`
}

type StorageStatus struct {
	AWSEFS EFSStatus `json:"awsEFS,omitempty"`
	AWSS3  S3Status  `json:"awsS3,omitempty"`
}

type ServiceAccountSpec struct {
	Name string `json:"name,omitempty"`
	// Service account annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// WorkspaceSpec defines the desired state of Workspace
type WorkspaceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// The username of the user
	Username string `json:"username,omitempty"`
	// Namespace to create for the workspace
	Namespace string `json:"namespace,omitempty"`
	// Storage parameters
	Storage StorageSpec `json:"storage,omitempty"`
	// Service account
	ServiceAccount ServiceAccountSpec `json:"serviceAccount,omitempty"`
}

// WorkspaceStatus defines the observed state of Workspace
type WorkspaceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Name of child namespace
	Namespace string `json:"namespace,omitempty"`
	// The AWS Role created for the user's workspace
	AWSRole string `json:"awsRole,omitempty"`
	// Storage parameters
	Storage StorageStatus `json:"storage,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Workspace is the Schema for the workspaces API
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkspaceSpec   `json:"spec,omitempty"`
	Status WorkspaceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// WorkspaceList contains a list of Workspace
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workspace `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Workspace{}, &WorkspaceList{})
}
