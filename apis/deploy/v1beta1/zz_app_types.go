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

type AppInitParameters struct {

	// The compute platform can either be ECS, Lambda, or Server. Default is Server.
	ComputePlatform *string `json:"computePlatform,omitempty" tf:"compute_platform,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AppObservation struct {

	// The application ID.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// The ARN of the CodeDeploy application.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The compute platform can either be ECS, Lambda, or Server. Default is Server.
	ComputePlatform *string `json:"computePlatform,omitempty" tf:"compute_platform,omitempty"`

	// The name for a connection to a GitHub account.
	GithubAccountName *string `json:"githubAccountName,omitempty" tf:"github_account_name,omitempty"`

	// Amazon's assigned ID for the application.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether the user has authenticated with GitHub for the specified application.
	LinkedToGithub *bool `json:"linkedToGithub,omitempty" tf:"linked_to_github,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type AppParameters struct {

	// The compute platform can either be ECS, Lambda, or Server. Default is Server.
	// +kubebuilder:validation:Optional
	ComputePlatform *string `json:"computePlatform,omitempty" tf:"compute_platform,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// AppSpec defines the desired state of App
type AppSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppParameters `json:"forProvider"`
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
	InitProvider AppInitParameters `json:"initProvider,omitempty"`
}

// AppStatus defines the observed state of App.
type AppStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// App is the Schema for the Apps API. Provides a CodeDeploy application.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type App struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppSpec   `json:"spec"`
	Status            AppStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppList contains a list of Apps
type AppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []App `json:"items"`
}

// Repository type metadata.
var (
	App_Kind             = "App"
	App_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: App_Kind}.String()
	App_KindAPIVersion   = App_Kind + "." + CRDGroupVersion.String()
	App_GroupVersionKind = CRDGroupVersion.WithKind(App_Kind)
)

func init() {
	SchemeBuilder.Register(&App{}, &AppList{})
}
