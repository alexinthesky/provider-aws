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

type ContactFlowModuleInitParameters struct {

	// Specifies the content of the Contact Flow Module, provided as a JSON string, written in Amazon Connect Contact Flow Language. If defined, the filename argument cannot be used.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the Contact Flow Module source specified with filename. The usual way to set this is filebase64sha256("contact_flow_module.11.12 and later) or base64sha256(file("contact_flow_module.11.11 and earlier), where "contact_flow_module.json" is the local filename of the Contact Flow Module source.
	ContentHash *string `json:"contentHash,omitempty" tf:"content_hash,omitempty"`

	// Specifies the description of the Contact Flow Module.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The path to the Contact Flow Module source within the local filesystem. Conflicts with content.
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// Specifies the name of the Contact Flow Module.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ContactFlowModuleObservation struct {

	// The Amazon Resource Name (ARN) of the Contact Flow Module.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The identifier of the Contact Flow Module.
	ContactFlowModuleID *string `json:"contactFlowModuleId,omitempty" tf:"contact_flow_module_id,omitempty"`

	// Specifies the content of the Contact Flow Module, provided as a JSON string, written in Amazon Connect Contact Flow Language. If defined, the filename argument cannot be used.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the Contact Flow Module source specified with filename. The usual way to set this is filebase64sha256("contact_flow_module.11.12 and later) or base64sha256(file("contact_flow_module.11.11 and earlier), where "contact_flow_module.json" is the local filename of the Contact Flow Module source.
	ContentHash *string `json:"contentHash,omitempty" tf:"content_hash,omitempty"`

	// Specifies the description of the Contact Flow Module.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The path to the Contact Flow Module source within the local filesystem. Conflicts with content.
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// The identifier of the hosting Amazon Connect Instance and identifier of the Contact Flow Module separated by a colon (:).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Specifies the name of the Contact Flow Module.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ContactFlowModuleParameters struct {

	// Specifies the content of the Contact Flow Module, provided as a JSON string, written in Amazon Connect Contact Flow Language. If defined, the filename argument cannot be used.
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the Contact Flow Module source specified with filename. The usual way to set this is filebase64sha256("contact_flow_module.11.12 and later) or base64sha256(file("contact_flow_module.11.11 and earlier), where "contact_flow_module.json" is the local filename of the Contact Flow Module source.
	// +kubebuilder:validation:Optional
	ContentHash *string `json:"contentHash,omitempty" tf:"content_hash,omitempty"`

	// Specifies the description of the Contact Flow Module.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The path to the Contact Flow Module source within the local filesystem. Conflicts with content.
	// +kubebuilder:validation:Optional
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/connect/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies the name of the Contact Flow Module.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ContactFlowModuleSpec defines the desired state of ContactFlowModule
type ContactFlowModuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ContactFlowModuleParameters `json:"forProvider"`
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
	InitProvider ContactFlowModuleInitParameters `json:"initProvider,omitempty"`
}

// ContactFlowModuleStatus defines the observed state of ContactFlowModule.
type ContactFlowModuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ContactFlowModuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ContactFlowModule is the Schema for the ContactFlowModules API. Provides details about a specific Amazon Connect Contact Flow Module.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ContactFlowModule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ContactFlowModuleSpec   `json:"spec"`
	Status ContactFlowModuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContactFlowModuleList contains a list of ContactFlowModules
type ContactFlowModuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContactFlowModule `json:"items"`
}

// Repository type metadata.
var (
	ContactFlowModule_Kind             = "ContactFlowModule"
	ContactFlowModule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ContactFlowModule_Kind}.String()
	ContactFlowModule_KindAPIVersion   = ContactFlowModule_Kind + "." + CRDGroupVersion.String()
	ContactFlowModule_GroupVersionKind = CRDGroupVersion.WithKind(ContactFlowModule_Kind)
)

func init() {
	SchemeBuilder.Register(&ContactFlowModule{}, &ContactFlowModuleList{})
}
