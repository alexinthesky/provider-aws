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

type NotebookInstanceLifecycleConfigurationInitParameters struct {

	// A shell script (base64-encoded) that runs only once when the SageMaker Notebook Instance is created.
	OnCreate *string `json:"onCreate,omitempty" tf:"on_create,omitempty"`

	// A shell script (base64-encoded) that runs every time the SageMaker Notebook Instance is started including the time it's created.
	OnStart *string `json:"onStart,omitempty" tf:"on_start,omitempty"`
}

type NotebookInstanceLifecycleConfigurationObservation struct {

	// The Amazon Resource Name (ARN) assigned by AWS to this lifecycle configuration.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A shell script (base64-encoded) that runs only once when the SageMaker Notebook Instance is created.
	OnCreate *string `json:"onCreate,omitempty" tf:"on_create,omitempty"`

	// A shell script (base64-encoded) that runs every time the SageMaker Notebook Instance is started including the time it's created.
	OnStart *string `json:"onStart,omitempty" tf:"on_start,omitempty"`
}

type NotebookInstanceLifecycleConfigurationParameters struct {

	// A shell script (base64-encoded) that runs only once when the SageMaker Notebook Instance is created.
	// +kubebuilder:validation:Optional
	OnCreate *string `json:"onCreate,omitempty" tf:"on_create,omitempty"`

	// A shell script (base64-encoded) that runs every time the SageMaker Notebook Instance is started including the time it's created.
	// +kubebuilder:validation:Optional
	OnStart *string `json:"onStart,omitempty" tf:"on_start,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// NotebookInstanceLifecycleConfigurationSpec defines the desired state of NotebookInstanceLifecycleConfiguration
type NotebookInstanceLifecycleConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NotebookInstanceLifecycleConfigurationParameters `json:"forProvider"`
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
	InitProvider NotebookInstanceLifecycleConfigurationInitParameters `json:"initProvider,omitempty"`
}

// NotebookInstanceLifecycleConfigurationStatus defines the observed state of NotebookInstanceLifecycleConfiguration.
type NotebookInstanceLifecycleConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NotebookInstanceLifecycleConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NotebookInstanceLifecycleConfiguration is the Schema for the NotebookInstanceLifecycleConfigurations API. Provides a lifecycle configuration for SageMaker Notebook Instances.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type NotebookInstanceLifecycleConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NotebookInstanceLifecycleConfigurationSpec   `json:"spec"`
	Status            NotebookInstanceLifecycleConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NotebookInstanceLifecycleConfigurationList contains a list of NotebookInstanceLifecycleConfigurations
type NotebookInstanceLifecycleConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotebookInstanceLifecycleConfiguration `json:"items"`
}

// Repository type metadata.
var (
	NotebookInstanceLifecycleConfiguration_Kind             = "NotebookInstanceLifecycleConfiguration"
	NotebookInstanceLifecycleConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NotebookInstanceLifecycleConfiguration_Kind}.String()
	NotebookInstanceLifecycleConfiguration_KindAPIVersion   = NotebookInstanceLifecycleConfiguration_Kind + "." + CRDGroupVersion.String()
	NotebookInstanceLifecycleConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(NotebookInstanceLifecycleConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&NotebookInstanceLifecycleConfiguration{}, &NotebookInstanceLifecycleConfigurationList{})
}
