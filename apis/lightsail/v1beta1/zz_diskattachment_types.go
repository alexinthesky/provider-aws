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

type DiskAttachmentInitParameters struct {

	// The disk path to expose to the instance.
	DiskPath *string `json:"diskPath,omitempty" tf:"disk_path,omitempty"`
}

type DiskAttachmentObservation struct {

	// The name of the Lightsail Disk.
	DiskName *string `json:"diskName,omitempty" tf:"disk_name,omitempty"`

	// The disk path to expose to the instance.
	DiskPath *string `json:"diskPath,omitempty" tf:"disk_path,omitempty"`

	// A combination of attributes to create a unique id: disk_name,instance_name
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Lightsail Instance to attach to.
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`
}

type DiskAttachmentParameters struct {

	// The name of the Lightsail Disk.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lightsail/v1beta1.Disk
	// +kubebuilder:validation:Optional
	DiskName *string `json:"diskName,omitempty" tf:"disk_name,omitempty"`

	// Reference to a Disk in lightsail to populate diskName.
	// +kubebuilder:validation:Optional
	DiskNameRef *v1.Reference `json:"diskNameRef,omitempty" tf:"-"`

	// Selector for a Disk in lightsail to populate diskName.
	// +kubebuilder:validation:Optional
	DiskNameSelector *v1.Selector `json:"diskNameSelector,omitempty" tf:"-"`

	// The disk path to expose to the instance.
	// +kubebuilder:validation:Optional
	DiskPath *string `json:"diskPath,omitempty" tf:"disk_path,omitempty"`

	// The name of the Lightsail Instance to attach to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lightsail/v1beta1.Instance
	// +kubebuilder:validation:Optional
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// Reference to a Instance in lightsail to populate instanceName.
	// +kubebuilder:validation:Optional
	InstanceNameRef *v1.Reference `json:"instanceNameRef,omitempty" tf:"-"`

	// Selector for a Instance in lightsail to populate instanceName.
	// +kubebuilder:validation:Optional
	InstanceNameSelector *v1.Selector `json:"instanceNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// DiskAttachmentSpec defines the desired state of DiskAttachment
type DiskAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiskAttachmentParameters `json:"forProvider"`
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
	InitProvider DiskAttachmentInitParameters `json:"initProvider,omitempty"`
}

// DiskAttachmentStatus defines the observed state of DiskAttachment.
type DiskAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiskAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DiskAttachment is the Schema for the DiskAttachments API. Attaches a Lightsail disk to a Lightsail Instance
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DiskAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.diskPath) || (has(self.initProvider) && has(self.initProvider.diskPath))",message="spec.forProvider.diskPath is a required parameter"
	Spec   DiskAttachmentSpec   `json:"spec"`
	Status DiskAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiskAttachmentList contains a list of DiskAttachments
type DiskAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiskAttachment `json:"items"`
}

// Repository type metadata.
var (
	DiskAttachment_Kind             = "DiskAttachment"
	DiskAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DiskAttachment_Kind}.String()
	DiskAttachment_KindAPIVersion   = DiskAttachment_Kind + "." + CRDGroupVersion.String()
	DiskAttachment_GroupVersionKind = CRDGroupVersion.WithKind(DiskAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&DiskAttachment{}, &DiskAttachmentList{})
}
