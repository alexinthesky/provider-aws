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

type DefinitionInitParameters struct {

	// Specific log groups to use with the query.
	LogGroupNames []*string `json:"logGroupNames,omitempty" tf:"log_group_names,omitempty"`

	// The name of the query.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The query to save. You can read more about CloudWatch Logs Query Syntax in the documentation.
	QueryString *string `json:"queryString,omitempty" tf:"query_string,omitempty"`
}

type DefinitionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specific log groups to use with the query.
	LogGroupNames []*string `json:"logGroupNames,omitempty" tf:"log_group_names,omitempty"`

	// The name of the query.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The query definition ID.
	QueryDefinitionID *string `json:"queryDefinitionId,omitempty" tf:"query_definition_id,omitempty"`

	// The query to save. You can read more about CloudWatch Logs Query Syntax in the documentation.
	QueryString *string `json:"queryString,omitempty" tf:"query_string,omitempty"`
}

type DefinitionParameters struct {

	// Specific log groups to use with the query.
	// +kubebuilder:validation:Optional
	LogGroupNames []*string `json:"logGroupNames,omitempty" tf:"log_group_names,omitempty"`

	// The name of the query.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The query to save. You can read more about CloudWatch Logs Query Syntax in the documentation.
	// +kubebuilder:validation:Optional
	QueryString *string `json:"queryString,omitempty" tf:"query_string,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// DefinitionSpec defines the desired state of Definition
type DefinitionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DefinitionParameters `json:"forProvider"`
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
	InitProvider DefinitionInitParameters `json:"initProvider,omitempty"`
}

// DefinitionStatus defines the observed state of Definition.
type DefinitionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DefinitionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Definition is the Schema for the Definitions API. Provides a CloudWatch Logs query definition resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Definition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.queryString) || (has(self.initProvider) && has(self.initProvider.queryString))",message="spec.forProvider.queryString is a required parameter"
	Spec   DefinitionSpec   `json:"spec"`
	Status DefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefinitionList contains a list of Definitions
type DefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Definition `json:"items"`
}

// Repository type metadata.
var (
	Definition_Kind             = "Definition"
	Definition_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Definition_Kind}.String()
	Definition_KindAPIVersion   = Definition_Kind + "." + CRDGroupVersion.String()
	Definition_GroupVersionKind = CRDGroupVersion.WithKind(Definition_Kind)
)

func init() {
	SchemeBuilder.Register(&Definition{}, &DefinitionList{})
}
