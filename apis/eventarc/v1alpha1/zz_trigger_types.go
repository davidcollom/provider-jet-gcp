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

type CloudRunServiceObservation struct {
}

type CloudRunServiceParameters struct {

	// Optional. The relative path on the Cloud Run service the events should be sent to. The value must conform to the definition of URI path segment (section 3.3 of RFC2396). Examples: "/route", "route", "route/subroute".
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Required. The region the Cloud Run service is deployed in.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Required. The name of the Cloud Run service being addressed. See https://cloud.google.com/run/docs/reference/rest/v1/namespaces.services. Only services located in the same project of the trigger object can be addressed.
	// +kubebuilder:validation:Required
	Service *string `json:"service" tf:"service,omitempty"`
}

type DestinationObservation struct {
}

type DestinationParameters struct {

	// The Cloud Function resource name. Only Cloud Functions V2 is supported. Format: projects/{project}/locations/{location}/functions/{function}
	// +kubebuilder:validation:Optional
	CloudFunction *string `json:"cloudFunction,omitempty" tf:"cloud_function,omitempty"`

	// Cloud Run fully-managed service that receives the events. The service should be running in the same project of the trigger.
	// +kubebuilder:validation:Optional
	CloudRunService []CloudRunServiceParameters `json:"cloudRunService,omitempty" tf:"cloud_run_service,omitempty"`
}

type MatchingCriteriaObservation struct {
}

type MatchingCriteriaParameters struct {

	// Required. The name of a CloudEvents attribute. Currently, only a subset of attributes are supported for filtering. All triggers MUST provide a filter for the 'type' attribute.
	// +kubebuilder:validation:Required
	Attribute *string `json:"attribute" tf:"attribute,omitempty"`

	// Required. The value for the attribute.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type PubsubObservation struct {
	Subscription *string `json:"subscription,omitempty" tf:"subscription,omitempty"`
}

type PubsubParameters struct {

	// Optional. The name of the Pub/Sub topic created and managed by Eventarc system as a transport for the event delivery. Format: `projects/{PROJECT_ID}/topics/{TOPIC_NAME You may set an existing topic for triggers of the type google.cloud.pubsub.topic.v1.messagePublished` only. The topic you provide here will not be deleted by Eventarc at trigger deletion.
	// +kubebuilder:validation:Optional
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`
}

type TransportObservation struct {
}

type TransportParameters struct {

	// The Pub/Sub topic and subscription used by Eventarc as delivery intermediary.
	// +kubebuilder:validation:Optional
	Pubsub []PubsubParameters `json:"pubsub,omitempty" tf:"pubsub,omitempty"`
}

type TriggerObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type TriggerParameters struct {

	// Required. Destination specifies where the events should be sent to.
	// +kubebuilder:validation:Required
	Destination []DestinationParameters `json:"destination" tf:"destination,omitempty"`

	// Optional. User labels attached to the triggers that can be used to group resources.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The location for the resource
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Required. null The list of filters that applies to event attributes. Only events that match all the provided filters will be sent to the destination.
	// +kubebuilder:validation:Required
	MatchingCriteria []MatchingCriteriaParameters `json:"matchingCriteria" tf:"matching_criteria,omitempty"`

	// Required. The resource name of the trigger. Must be unique within the location on the project and must be in `projects/{project}/locations/{location}/triggers/{trigger}` format.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The project for the resource
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
	// +kubebuilder:validation:Optional
	ServiceAccount *string `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`

	// Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
	// +kubebuilder:validation:Optional
	Transport []TransportParameters `json:"transport,omitempty" tf:"transport,omitempty"`
}

// TriggerSpec defines the desired state of Trigger
type TriggerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TriggerParameters `json:"forProvider"`
}

// TriggerStatus defines the observed state of Trigger.
type TriggerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TriggerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Trigger is the Schema for the Triggers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Trigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TriggerSpec   `json:"spec"`
	Status            TriggerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TriggerList contains a list of Triggers
type TriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Trigger `json:"items"`
}

// Repository type metadata.
var (
	Trigger_Kind             = "Trigger"
	Trigger_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Trigger_Kind}.String()
	Trigger_KindAPIVersion   = Trigger_Kind + "." + CRDGroupVersion.String()
	Trigger_GroupVersionKind = CRDGroupVersion.WithKind(Trigger_Kind)
)

func init() {
	SchemeBuilder.Register(&Trigger{}, &TriggerList{})
}