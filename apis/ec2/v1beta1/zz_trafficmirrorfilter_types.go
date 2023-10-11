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

type TrafficMirrorFilterInitParameters struct {

	// A description of the filter.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// List of amazon network services that should be mirrored. Valid values: amazon-dns.
	NetworkServices []*string `json:"networkServices,omitempty" tf:"network_services,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type TrafficMirrorFilterObservation struct {

	// The ARN of the traffic mirror filter.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A description of the filter.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the filter.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of amazon network services that should be mirrored. Valid values: amazon-dns.
	NetworkServices []*string `json:"networkServices,omitempty" tf:"network_services,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type TrafficMirrorFilterParameters struct {

	// A description of the filter.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// List of amazon network services that should be mirrored. Valid values: amazon-dns.
	// +kubebuilder:validation:Optional
	NetworkServices []*string `json:"networkServices,omitempty" tf:"network_services,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// TrafficMirrorFilterSpec defines the desired state of TrafficMirrorFilter
type TrafficMirrorFilterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TrafficMirrorFilterParameters `json:"forProvider"`
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
	InitProvider TrafficMirrorFilterInitParameters `json:"initProvider,omitempty"`
}

// TrafficMirrorFilterStatus defines the observed state of TrafficMirrorFilter.
type TrafficMirrorFilterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TrafficMirrorFilterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TrafficMirrorFilter is the Schema for the TrafficMirrorFilters API. Provides an Traffic mirror filter
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TrafficMirrorFilter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TrafficMirrorFilterSpec   `json:"spec"`
	Status            TrafficMirrorFilterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrafficMirrorFilterList contains a list of TrafficMirrorFilters
type TrafficMirrorFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficMirrorFilter `json:"items"`
}

// Repository type metadata.
var (
	TrafficMirrorFilter_Kind             = "TrafficMirrorFilter"
	TrafficMirrorFilter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TrafficMirrorFilter_Kind}.String()
	TrafficMirrorFilter_KindAPIVersion   = TrafficMirrorFilter_Kind + "." + CRDGroupVersion.String()
	TrafficMirrorFilter_GroupVersionKind = CRDGroupVersion.WithKind(TrafficMirrorFilter_Kind)
)

func init() {
	SchemeBuilder.Register(&TrafficMirrorFilter{}, &TrafficMirrorFilterList{})
}
