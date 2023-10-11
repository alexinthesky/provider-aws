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

type LBStickinessPolicyInitParameters struct {

	// The cookie duration in seconds. This determines the length of the session stickiness.
	CookieDuration *float64 `json:"cookieDuration,omitempty" tf:"cookie_duration,omitempty"`

	// - The Session Stickiness state of the load balancer. true to activate session stickiness or false to deactivate session stickiness.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type LBStickinessPolicyObservation struct {

	// The cookie duration in seconds. This determines the length of the session stickiness.
	CookieDuration *float64 `json:"cookieDuration,omitempty" tf:"cookie_duration,omitempty"`

	// - The Session Stickiness state of the load balancer. true to activate session stickiness or false to deactivate session stickiness.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The name used for this load balancer (matches lb_name).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LBStickinessPolicyParameters struct {

	// The cookie duration in seconds. This determines the length of the session stickiness.
	// +kubebuilder:validation:Optional
	CookieDuration *float64 `json:"cookieDuration,omitempty" tf:"cookie_duration,omitempty"`

	// - The Session Stickiness state of the load balancer. true to activate session stickiness or false to deactivate session stickiness.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// LBStickinessPolicySpec defines the desired state of LBStickinessPolicy
type LBStickinessPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LBStickinessPolicyParameters `json:"forProvider"`
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
	InitProvider LBStickinessPolicyInitParameters `json:"initProvider,omitempty"`
}

// LBStickinessPolicyStatus defines the observed state of LBStickinessPolicy.
type LBStickinessPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LBStickinessPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LBStickinessPolicy is the Schema for the LBStickinessPolicys API. Configures Session Stickiness for a Lightsail Load Balancer
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LBStickinessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cookieDuration) || (has(self.initProvider) && has(self.initProvider.cookieDuration))",message="spec.forProvider.cookieDuration is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enabled) || (has(self.initProvider) && has(self.initProvider.enabled))",message="spec.forProvider.enabled is a required parameter"
	Spec   LBStickinessPolicySpec   `json:"spec"`
	Status LBStickinessPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LBStickinessPolicyList contains a list of LBStickinessPolicys
type LBStickinessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LBStickinessPolicy `json:"items"`
}

// Repository type metadata.
var (
	LBStickinessPolicy_Kind             = "LBStickinessPolicy"
	LBStickinessPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LBStickinessPolicy_Kind}.String()
	LBStickinessPolicy_KindAPIVersion   = LBStickinessPolicy_Kind + "." + CRDGroupVersion.String()
	LBStickinessPolicy_GroupVersionKind = CRDGroupVersion.WithKind(LBStickinessPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&LBStickinessPolicy{}, &LBStickinessPolicyList{})
}
