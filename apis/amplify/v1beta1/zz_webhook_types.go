// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type WebhookInitParameters struct {

	// Description for a webhook.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`
}

type WebhookObservation struct {

	// Unique ID for an Amplify app.
	AppID *string `json:"appId,omitempty" tf:"app_id,omitempty"`

	// ARN for the webhook.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Name for a branch that is part of the Amplify app.
	BranchName *string `json:"branchName,omitempty" tf:"branch_name,omitempty"`

	// Description for a webhook.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// URL of the webhook.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type WebhookParameters struct {

	// Unique ID for an Amplify app.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/amplify/v1beta1.App
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	AppID *string `json:"appId,omitempty" tf:"app_id,omitempty"`

	// Reference to a App in amplify to populate appId.
	// +kubebuilder:validation:Optional
	AppIDRef *v1.Reference `json:"appIdRef,omitempty" tf:"-"`

	// Selector for a App in amplify to populate appId.
	// +kubebuilder:validation:Optional
	AppIDSelector *v1.Selector `json:"appIdSelector,omitempty" tf:"-"`

	// Name for a branch that is part of the Amplify app.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/amplify/v1beta1.Branch
	// +kubebuilder:validation:Optional
	BranchName *string `json:"branchName,omitempty" tf:"branch_name,omitempty"`

	// Reference to a Branch in amplify to populate branchName.
	// +kubebuilder:validation:Optional
	BranchNameRef *v1.Reference `json:"branchNameRef,omitempty" tf:"-"`

	// Selector for a Branch in amplify to populate branchName.
	// +kubebuilder:validation:Optional
	BranchNameSelector *v1.Selector `json:"branchNameSelector,omitempty" tf:"-"`

	// Description for a webhook.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// WebhookSpec defines the desired state of Webhook
type WebhookSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WebhookParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider WebhookInitParameters `json:"initProvider,omitempty"`
}

// WebhookStatus defines the observed state of Webhook.
type WebhookStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WebhookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Webhook is the Schema for the Webhooks API. Provides an Amplify Webhook resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Webhook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WebhookSpec   `json:"spec"`
	Status            WebhookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebhookList contains a list of Webhooks
type WebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Webhook `json:"items"`
}

// Repository type metadata.
var (
	Webhook_Kind             = "Webhook"
	Webhook_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Webhook_Kind}.String()
	Webhook_KindAPIVersion   = Webhook_Kind + "." + CRDGroupVersion.String()
	Webhook_GroupVersionKind = CRDGroupVersion.WithKind(Webhook_Kind)
)

func init() {
	SchemeBuilder.Register(&Webhook{}, &WebhookList{})
}
