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

type ContentSourceConfigurationInitParameters struct {

	// The identifiers of the data sources you want to use for your Amazon Kendra experience. Maximum number of 100 items.
	DataSourceIds []*string `json:"dataSourceIds,omitempty" tf:"data_source_ids,omitempty"`

	// Whether to use documents you indexed directly using the BatchPutDocument API. Defaults to false.
	DirectPutContent *bool `json:"directPutContent,omitempty" tf:"direct_put_content,omitempty"`

	// The identifier of the FAQs that you want to use for your Amazon Kendra experience. Maximum number of 100 items.
	FaqIds []*string `json:"faqIds,omitempty" tf:"faq_ids,omitempty"`
}

type ContentSourceConfigurationObservation struct {

	// The identifiers of the data sources you want to use for your Amazon Kendra experience. Maximum number of 100 items.
	DataSourceIds []*string `json:"dataSourceIds,omitempty" tf:"data_source_ids,omitempty"`

	// Whether to use documents you indexed directly using the BatchPutDocument API. Defaults to false.
	DirectPutContent *bool `json:"directPutContent,omitempty" tf:"direct_put_content,omitempty"`

	// The identifier of the FAQs that you want to use for your Amazon Kendra experience. Maximum number of 100 items.
	FaqIds []*string `json:"faqIds,omitempty" tf:"faq_ids,omitempty"`
}

type ContentSourceConfigurationParameters struct {

	// The identifiers of the data sources you want to use for your Amazon Kendra experience. Maximum number of 100 items.
	// +kubebuilder:validation:Optional
	DataSourceIds []*string `json:"dataSourceIds,omitempty" tf:"data_source_ids,omitempty"`

	// Whether to use documents you indexed directly using the BatchPutDocument API. Defaults to false.
	// +kubebuilder:validation:Optional
	DirectPutContent *bool `json:"directPutContent,omitempty" tf:"direct_put_content,omitempty"`

	// The identifier of the FAQs that you want to use for your Amazon Kendra experience. Maximum number of 100 items.
	// +kubebuilder:validation:Optional
	FaqIds []*string `json:"faqIds,omitempty" tf:"faq_ids,omitempty"`
}

type EndpointsInitParameters struct {
}

type EndpointsObservation struct {

	// The endpoint of your Amazon Kendra experience.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The type of endpoint for your Amazon Kendra experience.
	EndpointType *string `json:"endpointType,omitempty" tf:"endpoint_type,omitempty"`
}

type EndpointsParameters struct {
}

type ExperienceConfigurationInitParameters struct {

	// The identifiers of your data sources and FAQs. Or, you can specify that you want to use documents indexed via the BatchPutDocument API. Detailed below.
	ContentSourceConfiguration []ContentSourceConfigurationInitParameters `json:"contentSourceConfiguration,omitempty" tf:"content_source_configuration,omitempty"`

	// The AWS SSO field name that contains the identifiers of your users, such as their emails. Detailed below.
	UserIdentityConfiguration []UserIdentityConfigurationInitParameters `json:"userIdentityConfiguration,omitempty" tf:"user_identity_configuration,omitempty"`
}

type ExperienceConfigurationObservation struct {

	// The identifiers of your data sources and FAQs. Or, you can specify that you want to use documents indexed via the BatchPutDocument API. Detailed below.
	ContentSourceConfiguration []ContentSourceConfigurationObservation `json:"contentSourceConfiguration,omitempty" tf:"content_source_configuration,omitempty"`

	// The AWS SSO field name that contains the identifiers of your users, such as their emails. Detailed below.
	UserIdentityConfiguration []UserIdentityConfigurationObservation `json:"userIdentityConfiguration,omitempty" tf:"user_identity_configuration,omitempty"`
}

type ExperienceConfigurationParameters struct {

	// The identifiers of your data sources and FAQs. Or, you can specify that you want to use documents indexed via the BatchPutDocument API. Detailed below.
	// +kubebuilder:validation:Optional
	ContentSourceConfiguration []ContentSourceConfigurationParameters `json:"contentSourceConfiguration,omitempty" tf:"content_source_configuration,omitempty"`

	// The AWS SSO field name that contains the identifiers of your users, such as their emails. Detailed below.
	// +kubebuilder:validation:Optional
	UserIdentityConfiguration []UserIdentityConfigurationParameters `json:"userIdentityConfiguration,omitempty" tf:"user_identity_configuration,omitempty"`
}

type ExperienceInitParameters struct {

	// Configuration information for your Amazon Kendra experience. Detailed below.
	Configuration []ExperienceConfigurationInitParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// A description for your Amazon Kendra experience.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A name for your Amazon Kendra experience.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ExperienceObservation struct {

	// ARN of the Experience.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Configuration information for your Amazon Kendra experience. Detailed below.
	Configuration []ExperienceConfigurationObservation `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// A description for your Amazon Kendra experience.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Shows the endpoint URLs for your Amazon Kendra experiences. The URLs are unique and fully hosted by AWS.
	Endpoints []EndpointsObservation `json:"endpoints,omitempty" tf:"endpoints,omitempty"`

	// The unique identifier of the experience.
	ExperienceID *string `json:"experienceId,omitempty" tf:"experience_id,omitempty"`

	// The unique identifiers of the experience and index separated by a slash (/).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The identifier of the index for your Amazon Kendra experience.
	IndexID *string `json:"indexId,omitempty" tf:"index_id,omitempty"`

	// A name for your Amazon Kendra experience.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Amazon Resource Name (ARN) of a role with permission to access Query API, QuerySuggestions API, SubmitFeedback API, and AWS SSO that stores your user and group information. For more information, see IAM roles for Amazon Kendra.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// The current processing status of your Amazon Kendra experience.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ExperienceParameters struct {

	// Configuration information for your Amazon Kendra experience. Detailed below.
	// +kubebuilder:validation:Optional
	Configuration []ExperienceConfigurationParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// A description for your Amazon Kendra experience.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The identifier of the index for your Amazon Kendra experience.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kendra/v1beta1.Index
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	IndexID *string `json:"indexId,omitempty" tf:"index_id,omitempty"`

	// Reference to a Index in kendra to populate indexId.
	// +kubebuilder:validation:Optional
	IndexIDRef *v1.Reference `json:"indexIdRef,omitempty" tf:"-"`

	// Selector for a Index in kendra to populate indexId.
	// +kubebuilder:validation:Optional
	IndexIDSelector *v1.Selector `json:"indexIdSelector,omitempty" tf:"-"`

	// A name for your Amazon Kendra experience.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The Amazon Resource Name (ARN) of a role with permission to access Query API, QuerySuggestions API, SubmitFeedback API, and AWS SSO that stores your user and group information. For more information, see IAM roles for Amazon Kendra.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`
}

type UserIdentityConfigurationInitParameters struct {

	// The AWS SSO field name that contains the identifiers of your users, such as their emails.
	IdentityAttributeName *string `json:"identityAttributeName,omitempty" tf:"identity_attribute_name,omitempty"`
}

type UserIdentityConfigurationObservation struct {

	// The AWS SSO field name that contains the identifiers of your users, such as their emails.
	IdentityAttributeName *string `json:"identityAttributeName,omitempty" tf:"identity_attribute_name,omitempty"`
}

type UserIdentityConfigurationParameters struct {

	// The AWS SSO field name that contains the identifiers of your users, such as their emails.
	// +kubebuilder:validation:Optional
	IdentityAttributeName *string `json:"identityAttributeName" tf:"identity_attribute_name,omitempty"`
}

// ExperienceSpec defines the desired state of Experience
type ExperienceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExperienceParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ExperienceInitParameters `json:"initProvider,omitempty"`
}

// ExperienceStatus defines the observed state of Experience.
type ExperienceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExperienceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Experience is the Schema for the Experiences API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Experience struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ExperienceSpec   `json:"spec"`
	Status ExperienceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExperienceList contains a list of Experiences
type ExperienceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Experience `json:"items"`
}

// Repository type metadata.
var (
	Experience_Kind             = "Experience"
	Experience_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Experience_Kind}.String()
	Experience_KindAPIVersion   = Experience_Kind + "." + CRDGroupVersion.String()
	Experience_GroupVersionKind = CRDGroupVersion.WithKind(Experience_Kind)
)

func init() {
	SchemeBuilder.Register(&Experience{}, &ExperienceList{})
}
