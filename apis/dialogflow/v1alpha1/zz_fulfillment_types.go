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

type FeaturesObservation struct {
}

type FeaturesParameters struct {

	// The type of the feature that enabled for fulfillment.
	// * SMALLTALK: Fulfillment is enabled for SmallTalk. Possible values: ["SMALLTALK"]
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type FulfillmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type FulfillmentParameters struct {

	// The human-readable name of the fulfillment, unique within the agent.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// Whether fulfillment is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The field defines whether the fulfillment is enabled for certain features.
	// +kubebuilder:validation:Optional
	Features []FeaturesParameters `json:"features,omitempty" tf:"features,omitempty"`

	// Represents configuration for a generic web service. Dialogflow supports two mechanisms for authentications: - Basic authentication with username and password. - Authentication with additional authentication headers.
	// +kubebuilder:validation:Optional
	GenericWebService []GenericWebServiceParameters `json:"genericWebService,omitempty" tf:"generic_web_service,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type GenericWebServiceObservation struct {
}

type GenericWebServiceParameters struct {

	// The password for HTTP Basic authentication.
	// +kubebuilder:validation:Optional
	Password *string `json:"password,omitempty" tf:"password,omitempty"`

	// The HTTP request headers to send together with fulfillment requests.
	// +kubebuilder:validation:Optional
	RequestHeaders map[string]*string `json:"requestHeaders,omitempty" tf:"request_headers,omitempty"`

	// The fulfillment URI for receiving POST requests. It must use https protocol.
	// +kubebuilder:validation:Required
	URI *string `json:"uri" tf:"uri,omitempty"`

	// The user name for HTTP Basic authentication.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// FulfillmentSpec defines the desired state of Fulfillment
type FulfillmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FulfillmentParameters `json:"forProvider"`
}

// FulfillmentStatus defines the observed state of Fulfillment.
type FulfillmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FulfillmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Fulfillment is the Schema for the Fulfillments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Fulfillment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FulfillmentSpec   `json:"spec"`
	Status            FulfillmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FulfillmentList contains a list of Fulfillments
type FulfillmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Fulfillment `json:"items"`
}

// Repository type metadata.
var (
	Fulfillment_Kind             = "Fulfillment"
	Fulfillment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Fulfillment_Kind}.String()
	Fulfillment_KindAPIVersion   = Fulfillment_Kind + "." + CRDGroupVersion.String()
	Fulfillment_GroupVersionKind = CRDGroupVersion.WithKind(Fulfillment_Kind)
)

func init() {
	SchemeBuilder.Register(&Fulfillment{}, &FulfillmentList{})
}