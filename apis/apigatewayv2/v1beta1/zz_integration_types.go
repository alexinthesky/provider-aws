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

type IntegrationInitParameters struct {

	// Type of the network connection to the integration endpoint. Valid values: INTERNET, VPC_LINK. Default is INTERNET.
	ConnectionType *string `json:"connectionType,omitempty" tf:"connection_type,omitempty"`

	// How to handle response payload content type conversions. Valid values: CONVERT_TO_BINARY, CONVERT_TO_TEXT. Supported only for WebSocket APIs.
	ContentHandlingStrategy *string `json:"contentHandlingStrategy,omitempty" tf:"content_handling_strategy,omitempty"`

	// Description of the integration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Integration's HTTP method. Must be specified if integration_type is not MOCK.
	IntegrationMethod *string `json:"integrationMethod,omitempty" tf:"integration_method,omitempty"`

	// AWS service action to invoke. Supported only for HTTP APIs when integration_type is AWS_PROXY. See the AWS service integration reference documentation for supported values. Must be between 1 and 128 characters in length.
	IntegrationSubtype *string `json:"integrationSubtype,omitempty" tf:"integration_subtype,omitempty"`

	// Integration type of an integration.
	// Valid values: AWS (supported only for WebSocket APIs), AWS_PROXY, HTTP (supported only for WebSocket APIs), HTTP_PROXY, MOCK (supported only for WebSocket APIs). For an HTTP API private integration, use HTTP_PROXY.
	IntegrationType *string `json:"integrationType,omitempty" tf:"integration_type,omitempty"`

	// Pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the request_templates attribute.
	// Valid values: WHEN_NO_MATCH, WHEN_NO_TEMPLATES, NEVER. Default is WHEN_NO_MATCH. Supported only for WebSocket APIs.
	PassthroughBehavior *string `json:"passthroughBehavior,omitempty" tf:"passthrough_behavior,omitempty"`

	// The format of the payload sent to an integration. Valid values: 1.0, 2.0. Default is 1.0.
	PayloadFormatVersion *string `json:"payloadFormatVersion,omitempty" tf:"payload_format_version,omitempty"`

	// For WebSocket APIs, a key-value map specifying request parameters that are passed from the method request to the backend.
	// For HTTP APIs with a specified integration_subtype, a key-value map specifying parameters that are passed to AWS_PROXY integrations.
	// For HTTP APIs without a specified integration_subtype, a key-value map specifying how to transform HTTP requests before sending them to the backend.
	// See the Amazon API Gateway Developer Guide for details.
	RequestParameters map[string]*string `json:"requestParameters,omitempty" tf:"request_parameters,omitempty"`

	// Map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client. Supported only for WebSocket APIs.
	RequestTemplates map[string]*string `json:"requestTemplates,omitempty" tf:"request_templates,omitempty"`

	// Mappings to transform the HTTP response from a backend integration before returning the response to clients. Supported only for HTTP APIs.
	ResponseParameters []ResponseParametersInitParameters `json:"responseParameters,omitempty" tf:"response_parameters,omitempty"`

	// TLS configuration for a private integration. Supported only for HTTP APIs.
	TLSConfig []TLSConfigInitParameters `json:"tlsConfig,omitempty" tf:"tls_config,omitempty"`

	// The template selection expression for the integration.
	TemplateSelectionExpression *string `json:"templateSelectionExpression,omitempty" tf:"template_selection_expression,omitempty"`

	// Custom timeout between 50 and 29,000 milliseconds for WebSocket APIs and between 50 and 30,000 milliseconds for HTTP APIs.
	// The default timeout is 29 seconds for WebSocket APIs and 30 seconds for HTTP APIs.
	TimeoutMilliseconds *float64 `json:"timeoutMilliseconds,omitempty" tf:"timeout_milliseconds,omitempty"`
}

type IntegrationObservation struct {

	// API identifier.
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// ID of the VPC link for a private integration. Supported only for HTTP APIs. Must be between 1 and 1024 characters in length.
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// Type of the network connection to the integration endpoint. Valid values: INTERNET, VPC_LINK. Default is INTERNET.
	ConnectionType *string `json:"connectionType,omitempty" tf:"connection_type,omitempty"`

	// How to handle response payload content type conversions. Valid values: CONVERT_TO_BINARY, CONVERT_TO_TEXT. Supported only for WebSocket APIs.
	ContentHandlingStrategy *string `json:"contentHandlingStrategy,omitempty" tf:"content_handling_strategy,omitempty"`

	// Credentials required for the integration, if any.
	CredentialsArn *string `json:"credentialsArn,omitempty" tf:"credentials_arn,omitempty"`

	// Description of the integration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Integration identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Integration's HTTP method. Must be specified if integration_type is not MOCK.
	IntegrationMethod *string `json:"integrationMethod,omitempty" tf:"integration_method,omitempty"`

	// The integration response selection expression for the integration.
	IntegrationResponseSelectionExpression *string `json:"integrationResponseSelectionExpression,omitempty" tf:"integration_response_selection_expression,omitempty"`

	// AWS service action to invoke. Supported only for HTTP APIs when integration_type is AWS_PROXY. See the AWS service integration reference documentation for supported values. Must be between 1 and 128 characters in length.
	IntegrationSubtype *string `json:"integrationSubtype,omitempty" tf:"integration_subtype,omitempty"`

	// Integration type of an integration.
	// Valid values: AWS (supported only for WebSocket APIs), AWS_PROXY, HTTP (supported only for WebSocket APIs), HTTP_PROXY, MOCK (supported only for WebSocket APIs). For an HTTP API private integration, use HTTP_PROXY.
	IntegrationType *string `json:"integrationType,omitempty" tf:"integration_type,omitempty"`

	// URI of the Lambda function for a Lambda proxy integration, when integration_type is AWS_PROXY.
	// For an HTTP integration, specify a fully-qualified URL. For an HTTP API private integration, specify the ARN of an Application Load Balancer listener, Network Load Balancer listener, or AWS Cloud Map service.
	IntegrationURI *string `json:"integrationUri,omitempty" tf:"integration_uri,omitempty"`

	// Pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the request_templates attribute.
	// Valid values: WHEN_NO_MATCH, WHEN_NO_TEMPLATES, NEVER. Default is WHEN_NO_MATCH. Supported only for WebSocket APIs.
	PassthroughBehavior *string `json:"passthroughBehavior,omitempty" tf:"passthrough_behavior,omitempty"`

	// The format of the payload sent to an integration. Valid values: 1.0, 2.0. Default is 1.0.
	PayloadFormatVersion *string `json:"payloadFormatVersion,omitempty" tf:"payload_format_version,omitempty"`

	// For WebSocket APIs, a key-value map specifying request parameters that are passed from the method request to the backend.
	// For HTTP APIs with a specified integration_subtype, a key-value map specifying parameters that are passed to AWS_PROXY integrations.
	// For HTTP APIs without a specified integration_subtype, a key-value map specifying how to transform HTTP requests before sending them to the backend.
	// See the Amazon API Gateway Developer Guide for details.
	RequestParameters map[string]*string `json:"requestParameters,omitempty" tf:"request_parameters,omitempty"`

	// Map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client. Supported only for WebSocket APIs.
	RequestTemplates map[string]*string `json:"requestTemplates,omitempty" tf:"request_templates,omitempty"`

	// Mappings to transform the HTTP response from a backend integration before returning the response to clients. Supported only for HTTP APIs.
	ResponseParameters []ResponseParametersObservation `json:"responseParameters,omitempty" tf:"response_parameters,omitempty"`

	// TLS configuration for a private integration. Supported only for HTTP APIs.
	TLSConfig []TLSConfigObservation `json:"tlsConfig,omitempty" tf:"tls_config,omitempty"`

	// The template selection expression for the integration.
	TemplateSelectionExpression *string `json:"templateSelectionExpression,omitempty" tf:"template_selection_expression,omitempty"`

	// Custom timeout between 50 and 29,000 milliseconds for WebSocket APIs and between 50 and 30,000 milliseconds for HTTP APIs.
	// The default timeout is 29 seconds for WebSocket APIs and 30 seconds for HTTP APIs.
	TimeoutMilliseconds *float64 `json:"timeoutMilliseconds,omitempty" tf:"timeout_milliseconds,omitempty"`
}

type IntegrationParameters struct {

	// API identifier.
	// +crossplane:generate:reference:type=API
	// +kubebuilder:validation:Optional
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Reference to a API to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// Selector for a API to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// ID of the VPC link for a private integration. Supported only for HTTP APIs. Must be between 1 and 1024 characters in length.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigatewayv2/v1beta1.VPCLink
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// Reference to a VPCLink in apigatewayv2 to populate connectionId.
	// +kubebuilder:validation:Optional
	ConnectionIDRef *v1.Reference `json:"connectionIdRef,omitempty" tf:"-"`

	// Selector for a VPCLink in apigatewayv2 to populate connectionId.
	// +kubebuilder:validation:Optional
	ConnectionIDSelector *v1.Selector `json:"connectionIdSelector,omitempty" tf:"-"`

	// Type of the network connection to the integration endpoint. Valid values: INTERNET, VPC_LINK. Default is INTERNET.
	// +kubebuilder:validation:Optional
	ConnectionType *string `json:"connectionType,omitempty" tf:"connection_type,omitempty"`

	// How to handle response payload content type conversions. Valid values: CONVERT_TO_BINARY, CONVERT_TO_TEXT. Supported only for WebSocket APIs.
	// +kubebuilder:validation:Optional
	ContentHandlingStrategy *string `json:"contentHandlingStrategy,omitempty" tf:"content_handling_strategy,omitempty"`

	// Credentials required for the integration, if any.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	CredentialsArn *string `json:"credentialsArn,omitempty" tf:"credentials_arn,omitempty"`

	// Reference to a Role in iam to populate credentialsArn.
	// +kubebuilder:validation:Optional
	CredentialsArnRef *v1.Reference `json:"credentialsArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate credentialsArn.
	// +kubebuilder:validation:Optional
	CredentialsArnSelector *v1.Selector `json:"credentialsArnSelector,omitempty" tf:"-"`

	// Description of the integration.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Integration's HTTP method. Must be specified if integration_type is not MOCK.
	// +kubebuilder:validation:Optional
	IntegrationMethod *string `json:"integrationMethod,omitempty" tf:"integration_method,omitempty"`

	// AWS service action to invoke. Supported only for HTTP APIs when integration_type is AWS_PROXY. See the AWS service integration reference documentation for supported values. Must be between 1 and 128 characters in length.
	// +kubebuilder:validation:Optional
	IntegrationSubtype *string `json:"integrationSubtype,omitempty" tf:"integration_subtype,omitempty"`

	// Integration type of an integration.
	// Valid values: AWS (supported only for WebSocket APIs), AWS_PROXY, HTTP (supported only for WebSocket APIs), HTTP_PROXY, MOCK (supported only for WebSocket APIs). For an HTTP API private integration, use HTTP_PROXY.
	// +kubebuilder:validation:Optional
	IntegrationType *string `json:"integrationType,omitempty" tf:"integration_type,omitempty"`

	// URI of the Lambda function for a Lambda proxy integration, when integration_type is AWS_PROXY.
	// For an HTTP integration, specify a fully-qualified URL. For an HTTP API private integration, specify the ARN of an Application Load Balancer listener, Network Load Balancer listener, or AWS Cloud Map service.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lambda/v1beta1.Function
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("invoke_arn",true)
	// +kubebuilder:validation:Optional
	IntegrationURI *string `json:"integrationUri,omitempty" tf:"integration_uri,omitempty"`

	// Reference to a Function in lambda to populate integrationUri.
	// +kubebuilder:validation:Optional
	IntegrationURIRef *v1.Reference `json:"integrationUriRef,omitempty" tf:"-"`

	// Selector for a Function in lambda to populate integrationUri.
	// +kubebuilder:validation:Optional
	IntegrationURISelector *v1.Selector `json:"integrationUriSelector,omitempty" tf:"-"`

	// Pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the request_templates attribute.
	// Valid values: WHEN_NO_MATCH, WHEN_NO_TEMPLATES, NEVER. Default is WHEN_NO_MATCH. Supported only for WebSocket APIs.
	// +kubebuilder:validation:Optional
	PassthroughBehavior *string `json:"passthroughBehavior,omitempty" tf:"passthrough_behavior,omitempty"`

	// The format of the payload sent to an integration. Valid values: 1.0, 2.0. Default is 1.0.
	// +kubebuilder:validation:Optional
	PayloadFormatVersion *string `json:"payloadFormatVersion,omitempty" tf:"payload_format_version,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// For WebSocket APIs, a key-value map specifying request parameters that are passed from the method request to the backend.
	// For HTTP APIs with a specified integration_subtype, a key-value map specifying parameters that are passed to AWS_PROXY integrations.
	// For HTTP APIs without a specified integration_subtype, a key-value map specifying how to transform HTTP requests before sending them to the backend.
	// See the Amazon API Gateway Developer Guide for details.
	// +kubebuilder:validation:Optional
	RequestParameters map[string]*string `json:"requestParameters,omitempty" tf:"request_parameters,omitempty"`

	// Map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client. Supported only for WebSocket APIs.
	// +kubebuilder:validation:Optional
	RequestTemplates map[string]*string `json:"requestTemplates,omitempty" tf:"request_templates,omitempty"`

	// Mappings to transform the HTTP response from a backend integration before returning the response to clients. Supported only for HTTP APIs.
	// +kubebuilder:validation:Optional
	ResponseParameters []ResponseParametersParameters `json:"responseParameters,omitempty" tf:"response_parameters,omitempty"`

	// TLS configuration for a private integration. Supported only for HTTP APIs.
	// +kubebuilder:validation:Optional
	TLSConfig []TLSConfigParameters `json:"tlsConfig,omitempty" tf:"tls_config,omitempty"`

	// The template selection expression for the integration.
	// +kubebuilder:validation:Optional
	TemplateSelectionExpression *string `json:"templateSelectionExpression,omitempty" tf:"template_selection_expression,omitempty"`

	// Custom timeout between 50 and 29,000 milliseconds for WebSocket APIs and between 50 and 30,000 milliseconds for HTTP APIs.
	// The default timeout is 29 seconds for WebSocket APIs and 30 seconds for HTTP APIs.
	// +kubebuilder:validation:Optional
	TimeoutMilliseconds *float64 `json:"timeoutMilliseconds,omitempty" tf:"timeout_milliseconds,omitempty"`
}

type ResponseParametersInitParameters struct {

	// Key-value map. The key of this map identifies the location of the request parameter to change, and how to change it. The corresponding value specifies the new data for the parameter.
	// See the Amazon API Gateway Developer Guide for details.
	Mappings map[string]*string `json:"mappings,omitempty" tf:"mappings,omitempty"`

	// HTTP status code in the range 200-599.
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

type ResponseParametersObservation struct {

	// Key-value map. The key of this map identifies the location of the request parameter to change, and how to change it. The corresponding value specifies the new data for the parameter.
	// See the Amazon API Gateway Developer Guide for details.
	Mappings map[string]*string `json:"mappings,omitempty" tf:"mappings,omitempty"`

	// HTTP status code in the range 200-599.
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

type ResponseParametersParameters struct {

	// Key-value map. The key of this map identifies the location of the request parameter to change, and how to change it. The corresponding value specifies the new data for the parameter.
	// See the Amazon API Gateway Developer Guide for details.
	// +kubebuilder:validation:Optional
	Mappings map[string]*string `json:"mappings" tf:"mappings,omitempty"`

	// HTTP status code in the range 200-599.
	// +kubebuilder:validation:Optional
	StatusCode *string `json:"statusCode" tf:"status_code,omitempty"`
}

type TLSConfigInitParameters struct {

	// If you specify a server name, API Gateway uses it to verify the hostname on the integration's certificate. The server name is also included in the TLS handshake to support Server Name Indication (SNI) or virtual hosting.
	ServerNameToVerify *string `json:"serverNameToVerify,omitempty" tf:"server_name_to_verify,omitempty"`
}

type TLSConfigObservation struct {

	// If you specify a server name, API Gateway uses it to verify the hostname on the integration's certificate. The server name is also included in the TLS handshake to support Server Name Indication (SNI) or virtual hosting.
	ServerNameToVerify *string `json:"serverNameToVerify,omitempty" tf:"server_name_to_verify,omitempty"`
}

type TLSConfigParameters struct {

	// If you specify a server name, API Gateway uses it to verify the hostname on the integration's certificate. The server name is also included in the TLS handshake to support Server Name Indication (SNI) or virtual hosting.
	// +kubebuilder:validation:Optional
	ServerNameToVerify *string `json:"serverNameToVerify,omitempty" tf:"server_name_to_verify,omitempty"`
}

// IntegrationSpec defines the desired state of Integration
type IntegrationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IntegrationParameters `json:"forProvider"`
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
	InitProvider IntegrationInitParameters `json:"initProvider,omitempty"`
}

// IntegrationStatus defines the observed state of Integration.
type IntegrationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IntegrationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Integration is the Schema for the Integrations API. Manages an Amazon API Gateway Version 2 integration.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Integration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.integrationType) || (has(self.initProvider) && has(self.initProvider.integrationType))",message="spec.forProvider.integrationType is a required parameter"
	Spec   IntegrationSpec   `json:"spec"`
	Status IntegrationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IntegrationList contains a list of Integrations
type IntegrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Integration `json:"items"`
}

// Repository type metadata.
var (
	Integration_Kind             = "Integration"
	Integration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Integration_Kind}.String()
	Integration_KindAPIVersion   = Integration_Kind + "." + CRDGroupVersion.String()
	Integration_GroupVersionKind = CRDGroupVersion.WithKind(Integration_Kind)
)

func init() {
	SchemeBuilder.Register(&Integration{}, &IntegrationList{})
}
