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

type BotAliasInitParameters struct {

	// The name of the bot.
	BotName *string `json:"botName,omitempty" tf:"bot_name,omitempty"`

	// The name of the bot.
	BotVersion *string `json:"botVersion,omitempty" tf:"bot_version,omitempty"`

	// The settings that determine how Amazon Lex uses conversation logs for the alias. Attributes are documented under conversation_logs.
	ConversationLogs []ConversationLogsInitParameters `json:"conversationLogs,omitempty" tf:"conversation_logs,omitempty"`

	// A description of the alias. Must be less than or equal to 200 characters in length.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`
}

type BotAliasObservation struct {

	// The ARN of the bot alias.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The name of the bot.
	BotName *string `json:"botName,omitempty" tf:"bot_name,omitempty"`

	// The name of the bot.
	BotVersion *string `json:"botVersion,omitempty" tf:"bot_version,omitempty"`

	// Checksum of the bot alias.
	Checksum *string `json:"checksum,omitempty" tf:"checksum,omitempty"`

	// The settings that determine how Amazon Lex uses conversation logs for the alias. Attributes are documented under conversation_logs.
	ConversationLogs []ConversationLogsObservation `json:"conversationLogs,omitempty" tf:"conversation_logs,omitempty"`

	// The date that the bot alias was created.
	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date,omitempty"`

	// A description of the alias. Must be less than or equal to 200 characters in length.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The date that the bot alias was updated. When you create a resource, the creation date and the last updated date are the same.
	LastUpdatedDate *string `json:"lastUpdatedDate,omitempty" tf:"last_updated_date,omitempty"`
}

type BotAliasParameters struct {

	// The name of the bot.
	// +kubebuilder:validation:Optional
	BotName *string `json:"botName,omitempty" tf:"bot_name,omitempty"`

	// The name of the bot.
	// +kubebuilder:validation:Optional
	BotVersion *string `json:"botVersion,omitempty" tf:"bot_version,omitempty"`

	// The settings that determine how Amazon Lex uses conversation logs for the alias. Attributes are documented under conversation_logs.
	// +kubebuilder:validation:Optional
	ConversationLogs []ConversationLogsParameters `json:"conversationLogs,omitempty" tf:"conversation_logs,omitempty"`

	// A description of the alias. Must be less than or equal to 200 characters in length.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type ConversationLogsInitParameters struct {

	// The Amazon Resource Name (ARN) of the IAM role used to write your logs to CloudWatch Logs or an S3 bucket. Must be between 20 and 2048 characters in length.
	IAMRoleArn *string `json:"iamRoleArn,omitempty" tf:"iam_role_arn,omitempty"`

	// The settings for your conversation logs. You can log text, audio, or both. Attributes are documented under log_settings.
	LogSettings []LogSettingsInitParameters `json:"logSettings,omitempty" tf:"log_settings,omitempty"`
}

type ConversationLogsObservation struct {

	// The Amazon Resource Name (ARN) of the IAM role used to write your logs to CloudWatch Logs or an S3 bucket. Must be between 20 and 2048 characters in length.
	IAMRoleArn *string `json:"iamRoleArn,omitempty" tf:"iam_role_arn,omitempty"`

	// The settings for your conversation logs. You can log text, audio, or both. Attributes are documented under log_settings.
	LogSettings []LogSettingsObservation `json:"logSettings,omitempty" tf:"log_settings,omitempty"`
}

type ConversationLogsParameters struct {

	// The Amazon Resource Name (ARN) of the IAM role used to write your logs to CloudWatch Logs or an S3 bucket. Must be between 20 and 2048 characters in length.
	// +kubebuilder:validation:Optional
	IAMRoleArn *string `json:"iamRoleArn" tf:"iam_role_arn,omitempty"`

	// The settings for your conversation logs. You can log text, audio, or both. Attributes are documented under log_settings.
	// +kubebuilder:validation:Optional
	LogSettings []LogSettingsParameters `json:"logSettings,omitempty" tf:"log_settings,omitempty"`
}

type LogSettingsInitParameters struct {

	// The destination where logs are delivered. Options are CLOUDWATCH_LOGS or S3.
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// The Amazon Resource Name (ARN) of the key used to encrypt audio logs in an S3 bucket. This can only be specified when destination is set to S3. Must be between 20 and 2048 characters in length.
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// The type of logging that is enabled. Options are AUDIO or TEXT.
	LogType *string `json:"logType,omitempty" tf:"log_type,omitempty"`

	// The Amazon Resource Name (ARN) of the CloudWatch Logs log group or S3 bucket where the logs are delivered. Must be less than or equal to 2048 characters in length.
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`
}

type LogSettingsObservation struct {

	// The destination where logs are delivered. Options are CLOUDWATCH_LOGS or S3.
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// The Amazon Resource Name (ARN) of the key used to encrypt audio logs in an S3 bucket. This can only be specified when destination is set to S3. Must be between 20 and 2048 characters in length.
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// The type of logging that is enabled. Options are AUDIO or TEXT.
	LogType *string `json:"logType,omitempty" tf:"log_type,omitempty"`

	// The Amazon Resource Name (ARN) of the CloudWatch Logs log group or S3 bucket where the logs are delivered. Must be less than or equal to 2048 characters in length.
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// (Computed) The prefix of the S3 object key for AUDIO logs or the log stream name for TEXT logs.
	ResourcePrefix *string `json:"resourcePrefix,omitempty" tf:"resource_prefix,omitempty"`
}

type LogSettingsParameters struct {

	// The destination where logs are delivered. Options are CLOUDWATCH_LOGS or S3.
	// +kubebuilder:validation:Optional
	Destination *string `json:"destination" tf:"destination,omitempty"`

	// The Amazon Resource Name (ARN) of the key used to encrypt audio logs in an S3 bucket. This can only be specified when destination is set to S3. Must be between 20 and 2048 characters in length.
	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// The type of logging that is enabled. Options are AUDIO or TEXT.
	// +kubebuilder:validation:Optional
	LogType *string `json:"logType" tf:"log_type,omitempty"`

	// The Amazon Resource Name (ARN) of the CloudWatch Logs log group or S3 bucket where the logs are delivered. Must be less than or equal to 2048 characters in length.
	// +kubebuilder:validation:Optional
	ResourceArn *string `json:"resourceArn" tf:"resource_arn,omitempty"`
}

// BotAliasSpec defines the desired state of BotAlias
type BotAliasSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BotAliasParameters `json:"forProvider"`
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
	InitProvider BotAliasInitParameters `json:"initProvider,omitempty"`
}

// BotAliasStatus defines the observed state of BotAlias.
type BotAliasStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BotAliasObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BotAlias is the Schema for the BotAliass API. Provides an Amazon Lex Bot Alias resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BotAlias struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.botName) || (has(self.initProvider) && has(self.initProvider.botName))",message="spec.forProvider.botName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.botVersion) || (has(self.initProvider) && has(self.initProvider.botVersion))",message="spec.forProvider.botVersion is a required parameter"
	Spec   BotAliasSpec   `json:"spec"`
	Status BotAliasStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BotAliasList contains a list of BotAliass
type BotAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BotAlias `json:"items"`
}

// Repository type metadata.
var (
	BotAlias_Kind             = "BotAlias"
	BotAlias_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BotAlias_Kind}.String()
	BotAlias_KindAPIVersion   = BotAlias_Kind + "." + CRDGroupVersion.String()
	BotAlias_GroupVersionKind = CRDGroupVersion.WithKind(BotAlias_Kind)
)

func init() {
	SchemeBuilder.Register(&BotAlias{}, &BotAliasList{})
}
