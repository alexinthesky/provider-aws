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

type AddOnInitParameters struct {

	// The daily time when an automatic snapshot will be created. Must be in HH:00 format, and in an hourly increment and specified in Coordinated Universal Time (UTC). The snapshot will be automatically created between the time specified and up to 45 minutes after.
	SnapshotTime *string `json:"snapshotTime,omitempty" tf:"snapshot_time,omitempty"`

	// The status of the add on. Valid Values: Enabled, Disabled.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The add-on type. There is currently only one valid type AutoSnapshot.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AddOnObservation struct {

	// The daily time when an automatic snapshot will be created. Must be in HH:00 format, and in an hourly increment and specified in Coordinated Universal Time (UTC). The snapshot will be automatically created between the time specified and up to 45 minutes after.
	SnapshotTime *string `json:"snapshotTime,omitempty" tf:"snapshot_time,omitempty"`

	// The status of the add on. Valid Values: Enabled, Disabled.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The add-on type. There is currently only one valid type AutoSnapshot.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AddOnParameters struct {

	// The daily time when an automatic snapshot will be created. Must be in HH:00 format, and in an hourly increment and specified in Coordinated Universal Time (UTC). The snapshot will be automatically created between the time specified and up to 45 minutes after.
	// +kubebuilder:validation:Optional
	SnapshotTime *string `json:"snapshotTime" tf:"snapshot_time,omitempty"`

	// The status of the add on. Valid Values: Enabled, Disabled.
	// +kubebuilder:validation:Optional
	Status *string `json:"status" tf:"status,omitempty"`

	// The add-on type. There is currently only one valid type AutoSnapshot.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type InstanceInitParameters struct {

	// The add on configuration for the instance. Detailed below.
	AddOn []AddOnInitParameters `json:"addOn,omitempty" tf:"add_on,omitempty"`

	// The Availability Zone in which to create your
	// instance (see list below)
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The ID for a virtual private server image. A list of available blueprint IDs can be obtained using the AWS CLI command: aws lightsail get-blueprints
	BlueprintID *string `json:"blueprintId,omitempty" tf:"blueprint_id,omitempty"`

	// The bundle of specification information (see list below)
	BundleID *string `json:"bundleId,omitempty" tf:"bundle_id,omitempty"`

	// The IP address type of the Lightsail Instance. Valid Values: dualstack | ipv4.
	IPAddressType *string `json:"ipAddressType,omitempty" tf:"ip_address_type,omitempty"`

	// The name of your key pair. Created in the
	// Lightsail console (cannot use aws_key_pair at this time)
	KeyPairName *string `json:"keyPairName,omitempty" tf:"key_pair_name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Single lined launch script as a string to configure server with additional user data
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`
}

type InstanceObservation struct {

	// The add on configuration for the instance. Detailed below.
	AddOn []AddOnObservation `json:"addOn,omitempty" tf:"add_on,omitempty"`

	// The ARN of the Lightsail instance (matches id).
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The Availability Zone in which to create your
	// instance (see list below)
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The ID for a virtual private server image. A list of available blueprint IDs can be obtained using the AWS CLI command: aws lightsail get-blueprints
	BlueprintID *string `json:"blueprintId,omitempty" tf:"blueprint_id,omitempty"`

	// The bundle of specification information (see list below)
	BundleID *string `json:"bundleId,omitempty" tf:"bundle_id,omitempty"`

	// The number of vCPUs the instance has.
	CPUCount *float64 `json:"cpuCount,omitempty" tf:"cpu_count,omitempty"`

	// The timestamp when the instance was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The ARN of the Lightsail instance (matches arn).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IP address type of the Lightsail Instance. Valid Values: dualstack | ipv4.
	IPAddressType *string `json:"ipAddressType,omitempty" tf:"ip_address_type,omitempty"`

	// (Deprecated) The first IPv6 address of the Lightsail instance. Use ipv6_addresses attribute instead.
	IPv6Address *string `json:"ipv6Address,omitempty" tf:"ipv6_address,omitempty"`

	// List of IPv6 addresses for the Lightsail instance.
	IPv6Addresses []*string `json:"ipv6Addresses,omitempty" tf:"ipv6_addresses,omitempty"`

	// A Boolean value indicating whether this instance has a static IP assigned to it.
	IsStaticIP *bool `json:"isStaticIp,omitempty" tf:"is_static_ip,omitempty"`

	// The name of your key pair. Created in the
	// Lightsail console (cannot use aws_key_pair at this time)
	KeyPairName *string `json:"keyPairName,omitempty" tf:"key_pair_name,omitempty"`

	// The private IP address of the instance.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The public IP address of the instance.
	PublicIPAddress *string `json:"publicIpAddress,omitempty" tf:"public_ip_address,omitempty"`

	// The amount of RAM in GB on the instance (e.g., 1.0).
	RAMSize *float64 `json:"ramSize,omitempty" tf:"ram_size,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Single lined launch script as a string to configure server with additional user data
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// The user name for connecting to the instance (e.g., ec2-user).
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type InstanceParameters struct {

	// The add on configuration for the instance. Detailed below.
	// +kubebuilder:validation:Optional
	AddOn []AddOnParameters `json:"addOn,omitempty" tf:"add_on,omitempty"`

	// The Availability Zone in which to create your
	// instance (see list below)
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The ID for a virtual private server image. A list of available blueprint IDs can be obtained using the AWS CLI command: aws lightsail get-blueprints
	// +kubebuilder:validation:Optional
	BlueprintID *string `json:"blueprintId,omitempty" tf:"blueprint_id,omitempty"`

	// The bundle of specification information (see list below)
	// +kubebuilder:validation:Optional
	BundleID *string `json:"bundleId,omitempty" tf:"bundle_id,omitempty"`

	// The IP address type of the Lightsail Instance. Valid Values: dualstack | ipv4.
	// +kubebuilder:validation:Optional
	IPAddressType *string `json:"ipAddressType,omitempty" tf:"ip_address_type,omitempty"`

	// The name of your key pair. Created in the
	// Lightsail console (cannot use aws_key_pair at this time)
	// +kubebuilder:validation:Optional
	KeyPairName *string `json:"keyPairName,omitempty" tf:"key_pair_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Single lined launch script as a string to configure server with additional user data
	// +kubebuilder:validation:Optional
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
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
	InitProvider InstanceInitParameters `json:"initProvider,omitempty"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API. Provides an Lightsail Instance
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.availabilityZone) || (has(self.initProvider) && has(self.initProvider.availabilityZone))",message="spec.forProvider.availabilityZone is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.blueprintId) || (has(self.initProvider) && has(self.initProvider.blueprintId))",message="spec.forProvider.blueprintId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bundleId) || (has(self.initProvider) && has(self.initProvider.bundleId))",message="spec.forProvider.bundleId is a required parameter"
	Spec   InstanceSpec   `json:"spec"`
	Status InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
