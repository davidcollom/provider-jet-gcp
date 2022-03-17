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

type HTTPHealthCheckObservation_2 struct {
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type HTTPHealthCheckParameters_2 struct {

	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	// +kubebuilder:validation:Optional
	CheckIntervalSec *int64 `json:"checkIntervalSec,omitempty" tf:"check_interval_sec,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	// +kubebuilder:validation:Optional
	HealthyThreshold *int64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// The value of the host header in the HTTP health check request. If
	// left empty (default value), the public IP on behalf of which this
	// health check is performed will be used.
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The TCP port number for the HTTP health check request.
	// The default value is 80.
	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The request path of the HTTP health check request.
	// The default value is /.
	// +kubebuilder:validation:Optional
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	// +kubebuilder:validation:Optional
	TimeoutSec *int64 `json:"timeoutSec,omitempty" tf:"timeout_sec,omitempty"`

	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	// +kubebuilder:validation:Optional
	UnhealthyThreshold *int64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

// HTTPHealthCheckSpec defines the desired state of HTTPHealthCheck
type HTTPHealthCheckSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HTTPHealthCheckParameters_2 `json:"forProvider"`
}

// HTTPHealthCheckStatus defines the observed state of HTTPHealthCheck.
type HTTPHealthCheckStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HTTPHealthCheckObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HTTPHealthCheck is the Schema for the HTTPHealthChecks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type HTTPHealthCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HTTPHealthCheckSpec   `json:"spec"`
	Status            HTTPHealthCheckStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HTTPHealthCheckList contains a list of HTTPHealthChecks
type HTTPHealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HTTPHealthCheck `json:"items"`
}

// Repository type metadata.
var (
	HTTPHealthCheck_Kind             = "HTTPHealthCheck"
	HTTPHealthCheck_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HTTPHealthCheck_Kind}.String()
	HTTPHealthCheck_KindAPIVersion   = HTTPHealthCheck_Kind + "." + CRDGroupVersion.String()
	HTTPHealthCheck_GroupVersionKind = CRDGroupVersion.WithKind(HTTPHealthCheck_Kind)
)

func init() {
	SchemeBuilder.Register(&HTTPHealthCheck{}, &HTTPHealthCheckList{})
}