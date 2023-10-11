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

type ConfigurationSetInitParameters struct {

	// Whether messages that use the configuration set are required to use TLS. See below.
	DeliveryOptions []DeliveryOptionsInitParameters `json:"deliveryOptions,omitempty" tf:"delivery_options,omitempty"`

	// Whether or not Amazon SES publishes reputation metrics for the configuration set, such as bounce and complaint rates, to Amazon CloudWatch. The default value is false.
	ReputationMetricsEnabled *bool `json:"reputationMetricsEnabled,omitempty" tf:"reputation_metrics_enabled,omitempty"`

	// Whether email sending is enabled or disabled for the configuration set. The default value is true.
	SendingEnabled *bool `json:"sendingEnabled,omitempty" tf:"sending_enabled,omitempty"`

	// Domain that is used to redirect email recipients to an Amazon SES-operated domain. See below. NOTE: This functionality is best effort.
	TrackingOptions []TrackingOptionsInitParameters `json:"trackingOptions,omitempty" tf:"tracking_options,omitempty"`
}

type ConfigurationSetObservation struct {

	// SES configuration set ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Whether messages that use the configuration set are required to use TLS. See below.
	DeliveryOptions []DeliveryOptionsObservation `json:"deliveryOptions,omitempty" tf:"delivery_options,omitempty"`

	// SES configuration set name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Date and time at which the reputation metrics for the configuration set were last reset. Resetting these metrics is known as a fresh start.
	LastFreshStart *string `json:"lastFreshStart,omitempty" tf:"last_fresh_start,omitempty"`

	// Whether or not Amazon SES publishes reputation metrics for the configuration set, such as bounce and complaint rates, to Amazon CloudWatch. The default value is false.
	ReputationMetricsEnabled *bool `json:"reputationMetricsEnabled,omitempty" tf:"reputation_metrics_enabled,omitempty"`

	// Whether email sending is enabled or disabled for the configuration set. The default value is true.
	SendingEnabled *bool `json:"sendingEnabled,omitempty" tf:"sending_enabled,omitempty"`

	// Domain that is used to redirect email recipients to an Amazon SES-operated domain. See below. NOTE: This functionality is best effort.
	TrackingOptions []TrackingOptionsObservation `json:"trackingOptions,omitempty" tf:"tracking_options,omitempty"`
}

type ConfigurationSetParameters struct {

	// Whether messages that use the configuration set are required to use TLS. See below.
	// +kubebuilder:validation:Optional
	DeliveryOptions []DeliveryOptionsParameters `json:"deliveryOptions,omitempty" tf:"delivery_options,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Whether or not Amazon SES publishes reputation metrics for the configuration set, such as bounce and complaint rates, to Amazon CloudWatch. The default value is false.
	// +kubebuilder:validation:Optional
	ReputationMetricsEnabled *bool `json:"reputationMetricsEnabled,omitempty" tf:"reputation_metrics_enabled,omitempty"`

	// Whether email sending is enabled or disabled for the configuration set. The default value is true.
	// +kubebuilder:validation:Optional
	SendingEnabled *bool `json:"sendingEnabled,omitempty" tf:"sending_enabled,omitempty"`

	// Domain that is used to redirect email recipients to an Amazon SES-operated domain. See below. NOTE: This functionality is best effort.
	// +kubebuilder:validation:Optional
	TrackingOptions []TrackingOptionsParameters `json:"trackingOptions,omitempty" tf:"tracking_options,omitempty"`
}

type DeliveryOptionsInitParameters struct {

	// Whether messages that use the configuration set are required to use Transport Layer Security (TLS). If the value is Require, messages are only delivered if a TLS connection can be established. If the value is Optional, messages can be delivered in plain text if a TLS connection can't be established. Valid values: Require or Optional. Defaults to Optional.
	TLSPolicy *string `json:"tlsPolicy,omitempty" tf:"tls_policy,omitempty"`
}

type DeliveryOptionsObservation struct {

	// Whether messages that use the configuration set are required to use Transport Layer Security (TLS). If the value is Require, messages are only delivered if a TLS connection can be established. If the value is Optional, messages can be delivered in plain text if a TLS connection can't be established. Valid values: Require or Optional. Defaults to Optional.
	TLSPolicy *string `json:"tlsPolicy,omitempty" tf:"tls_policy,omitempty"`
}

type DeliveryOptionsParameters struct {

	// Whether messages that use the configuration set are required to use Transport Layer Security (TLS). If the value is Require, messages are only delivered if a TLS connection can be established. If the value is Optional, messages can be delivered in plain text if a TLS connection can't be established. Valid values: Require or Optional. Defaults to Optional.
	// +kubebuilder:validation:Optional
	TLSPolicy *string `json:"tlsPolicy,omitempty" tf:"tls_policy,omitempty"`
}

type TrackingOptionsInitParameters struct {

	// Custom subdomain that is used to redirect email recipients to the Amazon SES event tracking domain.
	CustomRedirectDomain *string `json:"customRedirectDomain,omitempty" tf:"custom_redirect_domain,omitempty"`
}

type TrackingOptionsObservation struct {

	// Custom subdomain that is used to redirect email recipients to the Amazon SES event tracking domain.
	CustomRedirectDomain *string `json:"customRedirectDomain,omitempty" tf:"custom_redirect_domain,omitempty"`
}

type TrackingOptionsParameters struct {

	// Custom subdomain that is used to redirect email recipients to the Amazon SES event tracking domain.
	// +kubebuilder:validation:Optional
	CustomRedirectDomain *string `json:"customRedirectDomain,omitempty" tf:"custom_redirect_domain,omitempty"`
}

// ConfigurationSetSpec defines the desired state of ConfigurationSet
type ConfigurationSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConfigurationSetParameters `json:"forProvider"`
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
	InitProvider ConfigurationSetInitParameters `json:"initProvider,omitempty"`
}

// ConfigurationSetStatus defines the observed state of ConfigurationSet.
type ConfigurationSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConfigurationSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigurationSet is the Schema for the ConfigurationSets API. Provides an SES configuration set
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ConfigurationSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigurationSetSpec   `json:"spec"`
	Status            ConfigurationSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigurationSetList contains a list of ConfigurationSets
type ConfigurationSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigurationSet `json:"items"`
}

// Repository type metadata.
var (
	ConfigurationSet_Kind             = "ConfigurationSet"
	ConfigurationSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConfigurationSet_Kind}.String()
	ConfigurationSet_KindAPIVersion   = ConfigurationSet_Kind + "." + CRDGroupVersion.String()
	ConfigurationSet_GroupVersionKind = CRDGroupVersion.WithKind(ConfigurationSet_Kind)
)

func init() {
	SchemeBuilder.Register(&ConfigurationSet{}, &ConfigurationSetList{})
}
