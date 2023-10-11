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

type ClusterActivityStreamInitParameters struct {

	// Specifies whether the database activity stream includes engine-native audit fields. This option only applies to an Oracle DB instance. By default, no engine-native audit fields are included. Defaults false.
	EngineNativeAuditFieldsIncluded *bool `json:"engineNativeAuditFieldsIncluded,omitempty" tf:"engine_native_audit_fields_included,omitempty"`

	// Specifies the mode of the database activity stream. Database events such as a change or access generate an activity stream event. The database session can handle these events either synchronously or asynchronously. One of: sync, async.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type ClusterActivityStreamObservation struct {

	// Specifies whether the database activity stream includes engine-native audit fields. This option only applies to an Oracle DB instance. By default, no engine-native audit fields are included. Defaults false.
	EngineNativeAuditFieldsIncluded *bool `json:"engineNativeAuditFieldsIncluded,omitempty" tf:"engine_native_audit_fields_included,omitempty"`

	// The Amazon Resource Name (ARN) of the DB cluster.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The AWS KMS key identifier for encrypting messages in the database activity stream. The AWS KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The name of the Amazon Kinesis data stream to be used for the database activity stream.
	KinesisStreamName *string `json:"kinesisStreamName,omitempty" tf:"kinesis_stream_name,omitempty"`

	// Specifies the mode of the database activity stream. Database events such as a change or access generate an activity stream event. The database session can handle these events either synchronously or asynchronously. One of: sync, async.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The Amazon Resource Name (ARN) of the DB cluster.
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`
}

type ClusterActivityStreamParameters struct {

	// Specifies whether the database activity stream includes engine-native audit fields. This option only applies to an Oracle DB instance. By default, no engine-native audit fields are included. Defaults false.
	// +kubebuilder:validation:Optional
	EngineNativeAuditFieldsIncluded *bool `json:"engineNativeAuditFieldsIncluded,omitempty" tf:"engine_native_audit_fields_included,omitempty"`

	// The AWS KMS key identifier for encrypting messages in the database activity stream. The AWS KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// Specifies the mode of the database activity stream. Database events such as a change or access generate an activity stream event. The database session can handle these events either synchronously or asynchronously. One of: sync, async.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The Amazon Resource Name (ARN) of the DB cluster.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/rds/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// Reference to a Cluster in rds to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnRef *v1.Reference `json:"resourceArnRef,omitempty" tf:"-"`

	// Selector for a Cluster in rds to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnSelector *v1.Selector `json:"resourceArnSelector,omitempty" tf:"-"`
}

// ClusterActivityStreamSpec defines the desired state of ClusterActivityStream
type ClusterActivityStreamSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterActivityStreamParameters `json:"forProvider"`
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
	InitProvider ClusterActivityStreamInitParameters `json:"initProvider,omitempty"`
}

// ClusterActivityStreamStatus defines the observed state of ClusterActivityStream.
type ClusterActivityStreamStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterActivityStreamObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterActivityStream is the Schema for the ClusterActivityStreams API. Manages RDS Aurora Cluster Database Activity Streams
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ClusterActivityStream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.mode) || (has(self.initProvider) && has(self.initProvider.mode))",message="spec.forProvider.mode is a required parameter"
	Spec   ClusterActivityStreamSpec   `json:"spec"`
	Status ClusterActivityStreamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterActivityStreamList contains a list of ClusterActivityStreams
type ClusterActivityStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterActivityStream `json:"items"`
}

// Repository type metadata.
var (
	ClusterActivityStream_Kind             = "ClusterActivityStream"
	ClusterActivityStream_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClusterActivityStream_Kind}.String()
	ClusterActivityStream_KindAPIVersion   = ClusterActivityStream_Kind + "." + CRDGroupVersion.String()
	ClusterActivityStream_GroupVersionKind = CRDGroupVersion.WithKind(ClusterActivityStream_Kind)
)

func init() {
	SchemeBuilder.Register(&ClusterActivityStream{}, &ClusterActivityStreamList{})
}
