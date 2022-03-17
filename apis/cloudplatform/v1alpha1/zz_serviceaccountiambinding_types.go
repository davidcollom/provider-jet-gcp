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

type ServiceAccountIAMBindingConditionObservation struct {
}

type ServiceAccountIAMBindingConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type ServiceAccountIAMBindingObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ServiceAccountIAMBindingParameters struct {

	// +kubebuilder:validation:Optional
	Condition []ServiceAccountIAMBindingConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Members []*string `json:"members" tf:"members,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Required
	ServiceAccountID *string `json:"serviceAccountId" tf:"service_account_id,omitempty"`
}

// ServiceAccountIAMBindingSpec defines the desired state of ServiceAccountIAMBinding
type ServiceAccountIAMBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceAccountIAMBindingParameters `json:"forProvider"`
}

// ServiceAccountIAMBindingStatus defines the observed state of ServiceAccountIAMBinding.
type ServiceAccountIAMBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceAccountIAMBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccountIAMBinding is the Schema for the ServiceAccountIAMBindings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type ServiceAccountIAMBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceAccountIAMBindingSpec   `json:"spec"`
	Status            ServiceAccountIAMBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccountIAMBindingList contains a list of ServiceAccountIAMBindings
type ServiceAccountIAMBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceAccountIAMBinding `json:"items"`
}

// Repository type metadata.
var (
	ServiceAccountIAMBinding_Kind             = "ServiceAccountIAMBinding"
	ServiceAccountIAMBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceAccountIAMBinding_Kind}.String()
	ServiceAccountIAMBinding_KindAPIVersion   = ServiceAccountIAMBinding_Kind + "." + CRDGroupVersion.String()
	ServiceAccountIAMBinding_GroupVersionKind = CRDGroupVersion.WithKind(ServiceAccountIAMBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceAccountIAMBinding{}, &ServiceAccountIAMBindingList{})
}