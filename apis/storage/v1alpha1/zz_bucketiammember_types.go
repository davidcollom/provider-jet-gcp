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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BucketIAMMemberConditionObservation struct {
}

type BucketIAMMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type BucketIAMMemberObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BucketIAMMemberParameters struct {

	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Optional
	Condition []BucketIAMMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Member *string `json:"member" tf:"member,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

// BucketIAMMemberSpec defines the desired state of BucketIAMMember
type BucketIAMMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketIAMMemberParameters `json:"forProvider"`
}

// BucketIAMMemberStatus defines the observed state of BucketIAMMember.
type BucketIAMMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketIAMMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketIAMMember is the Schema for the BucketIAMMembers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type BucketIAMMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketIAMMemberSpec   `json:"spec"`
	Status            BucketIAMMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketIAMMemberList contains a list of BucketIAMMembers
type BucketIAMMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketIAMMember `json:"items"`
}

// Repository type metadata.
var (
	BucketIAMMember_Kind             = "BucketIAMMember"
	BucketIAMMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketIAMMember_Kind}.String()
	BucketIAMMember_KindAPIVersion   = BucketIAMMember_Kind + "." + CRDGroupVersion.String()
	BucketIAMMember_GroupVersionKind = CRDGroupVersion.WithKind(BucketIAMMember_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketIAMMember{}, &BucketIAMMemberList{})
}