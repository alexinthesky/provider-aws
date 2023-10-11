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

type ArchiveRuleInitParameters struct {

	// Filter criteria for the archive rule. See Filter for more details.
	Filter []FilterInitParameters `json:"filter,omitempty" tf:"filter,omitempty"`
}

type ArchiveRuleObservation struct {

	// Analyzer name.
	AnalyzerName *string `json:"analyzerName,omitempty" tf:"analyzer_name,omitempty"`

	// Filter criteria for the archive rule. See Filter for more details.
	Filter []FilterObservation `json:"filter,omitempty" tf:"filter,omitempty"`

	// Resource ID in the format: analyzer_name/rule_name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ArchiveRuleParameters struct {

	// Analyzer name.
	// +kubebuilder:validation:Required
	AnalyzerName *string `json:"analyzerName" tf:"analyzer_name,omitempty"`

	// Filter criteria for the archive rule. See Filter for more details.
	// +kubebuilder:validation:Optional
	Filter []FilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type FilterInitParameters struct {

	// Contains comparator.
	Contains []*string `json:"contains,omitempty" tf:"contains,omitempty"`

	// Filter criteria.
	Criteria *string `json:"criteria,omitempty" tf:"criteria,omitempty"`

	// Equals comparator.
	Eq []*string `json:"eq,omitempty" tf:"eq,omitempty"`

	// Boolean comparator.
	Exists *string `json:"exists,omitempty" tf:"exists,omitempty"`

	// Not Equals comparator.
	Neq []*string `json:"neq,omitempty" tf:"neq,omitempty"`
}

type FilterObservation struct {

	// Contains comparator.
	Contains []*string `json:"contains,omitempty" tf:"contains,omitempty"`

	// Filter criteria.
	Criteria *string `json:"criteria,omitempty" tf:"criteria,omitempty"`

	// Equals comparator.
	Eq []*string `json:"eq,omitempty" tf:"eq,omitempty"`

	// Boolean comparator.
	Exists *string `json:"exists,omitempty" tf:"exists,omitempty"`

	// Not Equals comparator.
	Neq []*string `json:"neq,omitempty" tf:"neq,omitempty"`
}

type FilterParameters struct {

	// Contains comparator.
	// +kubebuilder:validation:Optional
	Contains []*string `json:"contains,omitempty" tf:"contains,omitempty"`

	// Filter criteria.
	// +kubebuilder:validation:Optional
	Criteria *string `json:"criteria" tf:"criteria,omitempty"`

	// Equals comparator.
	// +kubebuilder:validation:Optional
	Eq []*string `json:"eq,omitempty" tf:"eq,omitempty"`

	// Boolean comparator.
	// +kubebuilder:validation:Optional
	Exists *string `json:"exists,omitempty" tf:"exists,omitempty"`

	// Not Equals comparator.
	// +kubebuilder:validation:Optional
	Neq []*string `json:"neq,omitempty" tf:"neq,omitempty"`
}

// ArchiveRuleSpec defines the desired state of ArchiveRule
type ArchiveRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ArchiveRuleParameters `json:"forProvider"`
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
	InitProvider ArchiveRuleInitParameters `json:"initProvider,omitempty"`
}

// ArchiveRuleStatus defines the observed state of ArchiveRule.
type ArchiveRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ArchiveRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ArchiveRule is the Schema for the ArchiveRules API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ArchiveRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.filter) || (has(self.initProvider) && has(self.initProvider.filter))",message="spec.forProvider.filter is a required parameter"
	Spec   ArchiveRuleSpec   `json:"spec"`
	Status ArchiveRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ArchiveRuleList contains a list of ArchiveRules
type ArchiveRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ArchiveRule `json:"items"`
}

// Repository type metadata.
var (
	ArchiveRule_Kind             = "ArchiveRule"
	ArchiveRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ArchiveRule_Kind}.String()
	ArchiveRule_KindAPIVersion   = ArchiveRule_Kind + "." + CRDGroupVersion.String()
	ArchiveRule_GroupVersionKind = CRDGroupVersion.WithKind(ArchiveRule_Kind)
)

func init() {
	SchemeBuilder.Register(&ArchiveRule{}, &ArchiveRuleList{})
}
