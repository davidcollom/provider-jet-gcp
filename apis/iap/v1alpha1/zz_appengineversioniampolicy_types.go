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

type AppEngineVersionIAMPolicyObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AppEngineVersionIAMPolicyParameters struct {

	// +kubebuilder:validation:Required
	AppID *string `json:"appId" tf:"app_id,omitempty"`

	// +kubebuilder:validation:Required
	PolicyData *string `json:"policyData" tf:"policy_data,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Service *string `json:"service" tf:"service,omitempty"`

	// +kubebuilder:validation:Required
	VersionID *string `json:"versionId" tf:"version_id,omitempty"`
}

// AppEngineVersionIAMPolicySpec defines the desired state of AppEngineVersionIAMPolicy
type AppEngineVersionIAMPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppEngineVersionIAMPolicyParameters `json:"forProvider"`
}

// AppEngineVersionIAMPolicyStatus defines the observed state of AppEngineVersionIAMPolicy.
type AppEngineVersionIAMPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppEngineVersionIAMPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppEngineVersionIAMPolicy is the Schema for the AppEngineVersionIAMPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type AppEngineVersionIAMPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppEngineVersionIAMPolicySpec   `json:"spec"`
	Status            AppEngineVersionIAMPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppEngineVersionIAMPolicyList contains a list of AppEngineVersionIAMPolicys
type AppEngineVersionIAMPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppEngineVersionIAMPolicy `json:"items"`
}

// Repository type metadata.
var (
	AppEngineVersionIAMPolicy_Kind             = "AppEngineVersionIAMPolicy"
	AppEngineVersionIAMPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppEngineVersionIAMPolicy_Kind}.String()
	AppEngineVersionIAMPolicy_KindAPIVersion   = AppEngineVersionIAMPolicy_Kind + "." + CRDGroupVersion.String()
	AppEngineVersionIAMPolicy_GroupVersionKind = CRDGroupVersion.WithKind(AppEngineVersionIAMPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&AppEngineVersionIAMPolicy{}, &AppEngineVersionIAMPolicyList{})
}