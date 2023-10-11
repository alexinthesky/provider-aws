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

type ContainerInitParameters struct {

	// The DNS host name for the container.
	ContainerHostname *string `json:"containerHostname,omitempty" tf:"container_hostname,omitempty"`

	// Environment variables for the Docker container.
	// A list of key value pairs.
	Environment map[string]*string `json:"environment,omitempty" tf:"environment,omitempty"`

	// The registry path where the inference code image is stored in Amazon ECR.
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC). For more information see Using a Private Docker Registry for Real-Time Inference Containers. see Image Config.
	ImageConfig []ImageConfigInitParameters `json:"imageConfig,omitempty" tf:"image_config,omitempty"`

	// The container hosts value SingleModel/MultiModel. The default value is SingleModel.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The URL for the S3 location where model artifacts are stored.
	ModelDataURL *string `json:"modelDataUrl,omitempty" tf:"model_data_url,omitempty"`
}

type ContainerObservation struct {

	// The DNS host name for the container.
	ContainerHostname *string `json:"containerHostname,omitempty" tf:"container_hostname,omitempty"`

	// Environment variables for the Docker container.
	// A list of key value pairs.
	Environment map[string]*string `json:"environment,omitempty" tf:"environment,omitempty"`

	// The registry path where the inference code image is stored in Amazon ECR.
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC). For more information see Using a Private Docker Registry for Real-Time Inference Containers. see Image Config.
	ImageConfig []ImageConfigObservation `json:"imageConfig,omitempty" tf:"image_config,omitempty"`

	// The container hosts value SingleModel/MultiModel. The default value is SingleModel.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The URL for the S3 location where model artifacts are stored.
	ModelDataURL *string `json:"modelDataUrl,omitempty" tf:"model_data_url,omitempty"`
}

type ContainerParameters struct {

	// The DNS host name for the container.
	// +kubebuilder:validation:Optional
	ContainerHostname *string `json:"containerHostname,omitempty" tf:"container_hostname,omitempty"`

	// Environment variables for the Docker container.
	// A list of key value pairs.
	// +kubebuilder:validation:Optional
	Environment map[string]*string `json:"environment,omitempty" tf:"environment,omitempty"`

	// The registry path where the inference code image is stored in Amazon ECR.
	// +kubebuilder:validation:Optional
	Image *string `json:"image" tf:"image,omitempty"`

	// Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC). For more information see Using a Private Docker Registry for Real-Time Inference Containers. see Image Config.
	// +kubebuilder:validation:Optional
	ImageConfig []ImageConfigParameters `json:"imageConfig,omitempty" tf:"image_config,omitempty"`

	// The container hosts value SingleModel/MultiModel. The default value is SingleModel.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The URL for the S3 location where model artifacts are stored.
	// +kubebuilder:validation:Optional
	ModelDataURL *string `json:"modelDataUrl,omitempty" tf:"model_data_url,omitempty"`
}

type ImageConfigInitParameters struct {

	// Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC). Allowed values are: Platform and Vpc.
	RepositoryAccessMode *string `json:"repositoryAccessMode,omitempty" tf:"repository_access_mode,omitempty"`

	// Specifies an authentication configuration for the private docker registry where your model image is hosted. Specify a value for this property only if you specified Vpc as the value for the RepositoryAccessMode field, and the private Docker registry where the model image is hosted requires authentication. see Repository Auth Config.
	RepositoryAuthConfig []RepositoryAuthConfigInitParameters `json:"repositoryAuthConfig,omitempty" tf:"repository_auth_config,omitempty"`
}

type ImageConfigObservation struct {

	// Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC). Allowed values are: Platform and Vpc.
	RepositoryAccessMode *string `json:"repositoryAccessMode,omitempty" tf:"repository_access_mode,omitempty"`

	// Specifies an authentication configuration for the private docker registry where your model image is hosted. Specify a value for this property only if you specified Vpc as the value for the RepositoryAccessMode field, and the private Docker registry where the model image is hosted requires authentication. see Repository Auth Config.
	RepositoryAuthConfig []RepositoryAuthConfigObservation `json:"repositoryAuthConfig,omitempty" tf:"repository_auth_config,omitempty"`
}

type ImageConfigParameters struct {

	// Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC). Allowed values are: Platform and Vpc.
	// +kubebuilder:validation:Optional
	RepositoryAccessMode *string `json:"repositoryAccessMode" tf:"repository_access_mode,omitempty"`

	// Specifies an authentication configuration for the private docker registry where your model image is hosted. Specify a value for this property only if you specified Vpc as the value for the RepositoryAccessMode field, and the private Docker registry where the model image is hosted requires authentication. see Repository Auth Config.
	// +kubebuilder:validation:Optional
	RepositoryAuthConfig []RepositoryAuthConfigParameters `json:"repositoryAuthConfig,omitempty" tf:"repository_auth_config,omitempty"`
}

type ImageConfigRepositoryAuthConfigInitParameters struct {

	// The Amazon Resource Name (ARN) of an AWS Lambda function that provides credentials to authenticate to the private Docker registry where your model image is hosted. For information about how to create an AWS Lambda function, see Create a Lambda function with the console in the AWS Lambda Developer Guide.
	RepositoryCredentialsProviderArn *string `json:"repositoryCredentialsProviderArn,omitempty" tf:"repository_credentials_provider_arn,omitempty"`
}

type ImageConfigRepositoryAuthConfigObservation struct {

	// The Amazon Resource Name (ARN) of an AWS Lambda function that provides credentials to authenticate to the private Docker registry where your model image is hosted. For information about how to create an AWS Lambda function, see Create a Lambda function with the console in the AWS Lambda Developer Guide.
	RepositoryCredentialsProviderArn *string `json:"repositoryCredentialsProviderArn,omitempty" tf:"repository_credentials_provider_arn,omitempty"`
}

type ImageConfigRepositoryAuthConfigParameters struct {

	// The Amazon Resource Name (ARN) of an AWS Lambda function that provides credentials to authenticate to the private Docker registry where your model image is hosted. For information about how to create an AWS Lambda function, see Create a Lambda function with the console in the AWS Lambda Developer Guide.
	// +kubebuilder:validation:Optional
	RepositoryCredentialsProviderArn *string `json:"repositoryCredentialsProviderArn" tf:"repository_credentials_provider_arn,omitempty"`
}

type InferenceExecutionConfigInitParameters struct {

	// The container hosts value SingleModel/MultiModel. The default value is SingleModel.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type InferenceExecutionConfigObservation struct {

	// The container hosts value SingleModel/MultiModel. The default value is SingleModel.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type InferenceExecutionConfigParameters struct {

	// The container hosts value SingleModel/MultiModel. The default value is SingleModel.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode" tf:"mode,omitempty"`
}

type ModelInitParameters struct {

	// Specifies containers in the inference pipeline. If not specified, the primary_container argument is required. Fields are documented below.
	Container []ContainerInitParameters `json:"container,omitempty" tf:"container,omitempty"`

	// Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
	EnableNetworkIsolation *bool `json:"enableNetworkIsolation,omitempty" tf:"enable_network_isolation,omitempty"`

	// Specifies details of how containers in a multi-container endpoint are called. see Inference Execution Config.
	InferenceExecutionConfig []InferenceExecutionConfigInitParameters `json:"inferenceExecutionConfig,omitempty" tf:"inference_execution_config,omitempty"`

	// The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the container argument is required. Fields are documented below.
	PrimaryContainer []PrimaryContainerInitParameters `json:"primaryContainer,omitempty" tf:"primary_container,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.
	VPCConfig []VPCConfigInitParameters `json:"vpcConfig,omitempty" tf:"vpc_config,omitempty"`
}

type ModelObservation struct {

	// The Amazon Resource Name (ARN) assigned by AWS to this model.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Specifies containers in the inference pipeline. If not specified, the primary_container argument is required. Fields are documented below.
	Container []ContainerObservation `json:"container,omitempty" tf:"container,omitempty"`

	// Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
	EnableNetworkIsolation *bool `json:"enableNetworkIsolation,omitempty" tf:"enable_network_isolation,omitempty"`

	// A role that SageMaker can assume to access model artifacts and docker images for deployment.
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" tf:"execution_role_arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies details of how containers in a multi-container endpoint are called. see Inference Execution Config.
	InferenceExecutionConfig []InferenceExecutionConfigObservation `json:"inferenceExecutionConfig,omitempty" tf:"inference_execution_config,omitempty"`

	// The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the container argument is required. Fields are documented below.
	PrimaryContainer []PrimaryContainerObservation `json:"primaryContainer,omitempty" tf:"primary_container,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.
	VPCConfig []VPCConfigObservation `json:"vpcConfig,omitempty" tf:"vpc_config,omitempty"`
}

type ModelParameters struct {

	// Specifies containers in the inference pipeline. If not specified, the primary_container argument is required. Fields are documented below.
	// +kubebuilder:validation:Optional
	Container []ContainerParameters `json:"container,omitempty" tf:"container,omitempty"`

	// Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
	// +kubebuilder:validation:Optional
	EnableNetworkIsolation *bool `json:"enableNetworkIsolation,omitempty" tf:"enable_network_isolation,omitempty"`

	// A role that SageMaker can assume to access model artifacts and docker images for deployment.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" tf:"execution_role_arn,omitempty"`

	// Reference to a Role in iam to populate executionRoleArn.
	// +kubebuilder:validation:Optional
	ExecutionRoleArnRef *v1.Reference `json:"executionRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate executionRoleArn.
	// +kubebuilder:validation:Optional
	ExecutionRoleArnSelector *v1.Selector `json:"executionRoleArnSelector,omitempty" tf:"-"`

	// Specifies details of how containers in a multi-container endpoint are called. see Inference Execution Config.
	// +kubebuilder:validation:Optional
	InferenceExecutionConfig []InferenceExecutionConfigParameters `json:"inferenceExecutionConfig,omitempty" tf:"inference_execution_config,omitempty"`

	// The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the container argument is required. Fields are documented below.
	// +kubebuilder:validation:Optional
	PrimaryContainer []PrimaryContainerParameters `json:"primaryContainer,omitempty" tf:"primary_container,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.
	// +kubebuilder:validation:Optional
	VPCConfig []VPCConfigParameters `json:"vpcConfig,omitempty" tf:"vpc_config,omitempty"`
}

type PrimaryContainerImageConfigInitParameters struct {

	// Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC). Allowed values are: Platform and Vpc.
	RepositoryAccessMode *string `json:"repositoryAccessMode,omitempty" tf:"repository_access_mode,omitempty"`

	// Specifies an authentication configuration for the private docker registry where your model image is hosted. Specify a value for this property only if you specified Vpc as the value for the RepositoryAccessMode field, and the private Docker registry where the model image is hosted requires authentication. see Repository Auth Config.
	RepositoryAuthConfig []ImageConfigRepositoryAuthConfigInitParameters `json:"repositoryAuthConfig,omitempty" tf:"repository_auth_config,omitempty"`
}

type PrimaryContainerImageConfigObservation struct {

	// Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC). Allowed values are: Platform and Vpc.
	RepositoryAccessMode *string `json:"repositoryAccessMode,omitempty" tf:"repository_access_mode,omitempty"`

	// Specifies an authentication configuration for the private docker registry where your model image is hosted. Specify a value for this property only if you specified Vpc as the value for the RepositoryAccessMode field, and the private Docker registry where the model image is hosted requires authentication. see Repository Auth Config.
	RepositoryAuthConfig []ImageConfigRepositoryAuthConfigObservation `json:"repositoryAuthConfig,omitempty" tf:"repository_auth_config,omitempty"`
}

type PrimaryContainerImageConfigParameters struct {

	// Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC). Allowed values are: Platform and Vpc.
	// +kubebuilder:validation:Optional
	RepositoryAccessMode *string `json:"repositoryAccessMode" tf:"repository_access_mode,omitempty"`

	// Specifies an authentication configuration for the private docker registry where your model image is hosted. Specify a value for this property only if you specified Vpc as the value for the RepositoryAccessMode field, and the private Docker registry where the model image is hosted requires authentication. see Repository Auth Config.
	// +kubebuilder:validation:Optional
	RepositoryAuthConfig []ImageConfigRepositoryAuthConfigParameters `json:"repositoryAuthConfig,omitempty" tf:"repository_auth_config,omitempty"`
}

type PrimaryContainerInitParameters struct {

	// The DNS host name for the container.
	ContainerHostname *string `json:"containerHostname,omitempty" tf:"container_hostname,omitempty"`

	// Environment variables for the Docker container.
	// A list of key value pairs.
	Environment map[string]*string `json:"environment,omitempty" tf:"environment,omitempty"`

	// The registry path where the inference code image is stored in Amazon ECR.
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC). For more information see Using a Private Docker Registry for Real-Time Inference Containers. see Image Config.
	ImageConfig []PrimaryContainerImageConfigInitParameters `json:"imageConfig,omitempty" tf:"image_config,omitempty"`

	// The container hosts value SingleModel/MultiModel. The default value is SingleModel.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The URL for the S3 location where model artifacts are stored.
	ModelDataURL *string `json:"modelDataUrl,omitempty" tf:"model_data_url,omitempty"`
}

type PrimaryContainerObservation struct {

	// The DNS host name for the container.
	ContainerHostname *string `json:"containerHostname,omitempty" tf:"container_hostname,omitempty"`

	// Environment variables for the Docker container.
	// A list of key value pairs.
	Environment map[string]*string `json:"environment,omitempty" tf:"environment,omitempty"`

	// The registry path where the inference code image is stored in Amazon ECR.
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC). For more information see Using a Private Docker Registry for Real-Time Inference Containers. see Image Config.
	ImageConfig []PrimaryContainerImageConfigObservation `json:"imageConfig,omitempty" tf:"image_config,omitempty"`

	// The container hosts value SingleModel/MultiModel. The default value is SingleModel.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The URL for the S3 location where model artifacts are stored.
	ModelDataURL *string `json:"modelDataUrl,omitempty" tf:"model_data_url,omitempty"`
}

type PrimaryContainerParameters struct {

	// The DNS host name for the container.
	// +kubebuilder:validation:Optional
	ContainerHostname *string `json:"containerHostname,omitempty" tf:"container_hostname,omitempty"`

	// Environment variables for the Docker container.
	// A list of key value pairs.
	// +kubebuilder:validation:Optional
	Environment map[string]*string `json:"environment,omitempty" tf:"environment,omitempty"`

	// The registry path where the inference code image is stored in Amazon ECR.
	// +kubebuilder:validation:Optional
	Image *string `json:"image" tf:"image,omitempty"`

	// Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC). For more information see Using a Private Docker Registry for Real-Time Inference Containers. see Image Config.
	// +kubebuilder:validation:Optional
	ImageConfig []PrimaryContainerImageConfigParameters `json:"imageConfig,omitempty" tf:"image_config,omitempty"`

	// The container hosts value SingleModel/MultiModel. The default value is SingleModel.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The URL for the S3 location where model artifacts are stored.
	// +kubebuilder:validation:Optional
	ModelDataURL *string `json:"modelDataUrl,omitempty" tf:"model_data_url,omitempty"`
}

type RepositoryAuthConfigInitParameters struct {

	// The Amazon Resource Name (ARN) of an AWS Lambda function that provides credentials to authenticate to the private Docker registry where your model image is hosted. For information about how to create an AWS Lambda function, see Create a Lambda function with the console in the AWS Lambda Developer Guide.
	RepositoryCredentialsProviderArn *string `json:"repositoryCredentialsProviderArn,omitempty" tf:"repository_credentials_provider_arn,omitempty"`
}

type RepositoryAuthConfigObservation struct {

	// The Amazon Resource Name (ARN) of an AWS Lambda function that provides credentials to authenticate to the private Docker registry where your model image is hosted. For information about how to create an AWS Lambda function, see Create a Lambda function with the console in the AWS Lambda Developer Guide.
	RepositoryCredentialsProviderArn *string `json:"repositoryCredentialsProviderArn,omitempty" tf:"repository_credentials_provider_arn,omitempty"`
}

type RepositoryAuthConfigParameters struct {

	// The Amazon Resource Name (ARN) of an AWS Lambda function that provides credentials to authenticate to the private Docker registry where your model image is hosted. For information about how to create an AWS Lambda function, see Create a Lambda function with the console in the AWS Lambda Developer Guide.
	// +kubebuilder:validation:Optional
	RepositoryCredentialsProviderArn *string `json:"repositoryCredentialsProviderArn" tf:"repository_credentials_provider_arn,omitempty"`
}

type VPCConfigInitParameters struct {
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`
}

type VPCConfigObservation struct {
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`
}

type VPCConfigParameters struct {

	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds" tf:"security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	Subnets []*string `json:"subnets" tf:"subnets,omitempty"`
}

// ModelSpec defines the desired state of Model
type ModelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ModelParameters `json:"forProvider"`
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
	InitProvider ModelInitParameters `json:"initProvider,omitempty"`
}

// ModelStatus defines the observed state of Model.
type ModelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ModelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Model is the Schema for the Models API. Provides a SageMaker model resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Model struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ModelSpec   `json:"spec"`
	Status            ModelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ModelList contains a list of Models
type ModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Model `json:"items"`
}

// Repository type metadata.
var (
	Model_Kind             = "Model"
	Model_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Model_Kind}.String()
	Model_KindAPIVersion   = Model_Kind + "." + CRDGroupVersion.String()
	Model_GroupVersionKind = CRDGroupVersion.WithKind(Model_Kind)
)

func init() {
	SchemeBuilder.Register(&Model{}, &ModelList{})
}
