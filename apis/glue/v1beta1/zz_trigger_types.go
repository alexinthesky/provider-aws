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

type ActionsInitParameters struct {

	// Arguments to be passed to the job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes.
	Arguments map[string]*string `json:"arguments,omitempty" tf:"arguments,omitempty"`

	// Specifies configuration properties of a job run notification. See Notification Property details below.
	NotificationProperty []ActionsNotificationPropertyInitParameters `json:"notificationProperty,omitempty" tf:"notification_property,omitempty"`

	// The name of the Security Configuration structure to be used with this action.
	SecurityConfiguration *string `json:"securityConfiguration,omitempty" tf:"security_configuration,omitempty"`

	// The job run timeout in minutes. It overrides the timeout value of the job.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type ActionsNotificationPropertyInitParameters struct {

	// After a job run starts, the number of minutes to wait before sending a job run delay notification.
	NotifyDelayAfter *float64 `json:"notifyDelayAfter,omitempty" tf:"notify_delay_after,omitempty"`
}

type ActionsNotificationPropertyObservation struct {

	// After a job run starts, the number of minutes to wait before sending a job run delay notification.
	NotifyDelayAfter *float64 `json:"notifyDelayAfter,omitempty" tf:"notify_delay_after,omitempty"`
}

type ActionsNotificationPropertyParameters struct {

	// After a job run starts, the number of minutes to wait before sending a job run delay notification.
	// +kubebuilder:validation:Optional
	NotifyDelayAfter *float64 `json:"notifyDelayAfter,omitempty" tf:"notify_delay_after,omitempty"`
}

type ActionsObservation struct {

	// Arguments to be passed to the job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes.
	Arguments map[string]*string `json:"arguments,omitempty" tf:"arguments,omitempty"`

	// The name of the crawler to be executed. Conflicts with job_name.
	CrawlerName *string `json:"crawlerName,omitempty" tf:"crawler_name,omitempty"`

	// The name of a job to be executed. Conflicts with crawler_name.
	JobName *string `json:"jobName,omitempty" tf:"job_name,omitempty"`

	// Specifies configuration properties of a job run notification. See Notification Property details below.
	NotificationProperty []ActionsNotificationPropertyObservation `json:"notificationProperty,omitempty" tf:"notification_property,omitempty"`

	// The name of the Security Configuration structure to be used with this action.
	SecurityConfiguration *string `json:"securityConfiguration,omitempty" tf:"security_configuration,omitempty"`

	// The job run timeout in minutes. It overrides the timeout value of the job.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type ActionsParameters struct {

	// Arguments to be passed to the job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes.
	// +kubebuilder:validation:Optional
	Arguments map[string]*string `json:"arguments,omitempty" tf:"arguments,omitempty"`

	// The name of the crawler to be executed. Conflicts with job_name.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/glue/v1beta1.Crawler
	// +kubebuilder:validation:Optional
	CrawlerName *string `json:"crawlerName,omitempty" tf:"crawler_name,omitempty"`

	// Reference to a Crawler in glue to populate crawlerName.
	// +kubebuilder:validation:Optional
	CrawlerNameRef *v1.Reference `json:"crawlerNameRef,omitempty" tf:"-"`

	// Selector for a Crawler in glue to populate crawlerName.
	// +kubebuilder:validation:Optional
	CrawlerNameSelector *v1.Selector `json:"crawlerNameSelector,omitempty" tf:"-"`

	// The name of a job to be executed. Conflicts with crawler_name.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/glue/v1beta1.Job
	// +kubebuilder:validation:Optional
	JobName *string `json:"jobName,omitempty" tf:"job_name,omitempty"`

	// Reference to a Job in glue to populate jobName.
	// +kubebuilder:validation:Optional
	JobNameRef *v1.Reference `json:"jobNameRef,omitempty" tf:"-"`

	// Selector for a Job in glue to populate jobName.
	// +kubebuilder:validation:Optional
	JobNameSelector *v1.Selector `json:"jobNameSelector,omitempty" tf:"-"`

	// Specifies configuration properties of a job run notification. See Notification Property details below.
	// +kubebuilder:validation:Optional
	NotificationProperty []ActionsNotificationPropertyParameters `json:"notificationProperty,omitempty" tf:"notification_property,omitempty"`

	// The name of the Security Configuration structure to be used with this action.
	// +kubebuilder:validation:Optional
	SecurityConfiguration *string `json:"securityConfiguration,omitempty" tf:"security_configuration,omitempty"`

	// The job run timeout in minutes. It overrides the timeout value of the job.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type ConditionsInitParameters struct {

	// The condition crawl state. Currently, the values supported are RUNNING, SUCCEEDED, CANCELLED, and FAILED. If this is specified, crawler_name must also be specified. Conflicts with state.
	CrawlState *string `json:"crawlState,omitempty" tf:"crawl_state,omitempty"`

	// A logical operator. Defaults to EQUALS.
	LogicalOperator *string `json:"logicalOperator,omitempty" tf:"logical_operator,omitempty"`

	// The condition job state. Currently, the values supported are SUCCEEDED, STOPPED, TIMEOUT and FAILED. If this is specified, job_name must also be specified. Conflicts with crawler_state.
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type ConditionsObservation struct {

	// The condition crawl state. Currently, the values supported are RUNNING, SUCCEEDED, CANCELLED, and FAILED. If this is specified, crawler_name must also be specified. Conflicts with state.
	CrawlState *string `json:"crawlState,omitempty" tf:"crawl_state,omitempty"`

	// The name of the crawler to be executed. Conflicts with job_name.
	CrawlerName *string `json:"crawlerName,omitempty" tf:"crawler_name,omitempty"`

	// The name of a job to be executed. Conflicts with crawler_name.
	JobName *string `json:"jobName,omitempty" tf:"job_name,omitempty"`

	// A logical operator. Defaults to EQUALS.
	LogicalOperator *string `json:"logicalOperator,omitempty" tf:"logical_operator,omitempty"`

	// The condition job state. Currently, the values supported are SUCCEEDED, STOPPED, TIMEOUT and FAILED. If this is specified, job_name must also be specified. Conflicts with crawler_state.
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type ConditionsParameters struct {

	// The condition crawl state. Currently, the values supported are RUNNING, SUCCEEDED, CANCELLED, and FAILED. If this is specified, crawler_name must also be specified. Conflicts with state.
	// +kubebuilder:validation:Optional
	CrawlState *string `json:"crawlState,omitempty" tf:"crawl_state,omitempty"`

	// The name of the crawler to be executed. Conflicts with job_name.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/glue/v1beta1.Crawler
	// +kubebuilder:validation:Optional
	CrawlerName *string `json:"crawlerName,omitempty" tf:"crawler_name,omitempty"`

	// Reference to a Crawler in glue to populate crawlerName.
	// +kubebuilder:validation:Optional
	CrawlerNameRef *v1.Reference `json:"crawlerNameRef,omitempty" tf:"-"`

	// Selector for a Crawler in glue to populate crawlerName.
	// +kubebuilder:validation:Optional
	CrawlerNameSelector *v1.Selector `json:"crawlerNameSelector,omitempty" tf:"-"`

	// The name of a job to be executed. Conflicts with crawler_name.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/glue/v1beta1.Job
	// +kubebuilder:validation:Optional
	JobName *string `json:"jobName,omitempty" tf:"job_name,omitempty"`

	// Reference to a Job in glue to populate jobName.
	// +kubebuilder:validation:Optional
	JobNameRef *v1.Reference `json:"jobNameRef,omitempty" tf:"-"`

	// Selector for a Job in glue to populate jobName.
	// +kubebuilder:validation:Optional
	JobNameSelector *v1.Selector `json:"jobNameSelector,omitempty" tf:"-"`

	// A logical operator. Defaults to EQUALS.
	// +kubebuilder:validation:Optional
	LogicalOperator *string `json:"logicalOperator,omitempty" tf:"logical_operator,omitempty"`

	// The condition job state. Currently, the values supported are SUCCEEDED, STOPPED, TIMEOUT and FAILED. If this is specified, job_name must also be specified. Conflicts with crawler_state.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type EventBatchingConditionInitParameters struct {

	// Number of events that must be received from Amazon EventBridge before EventBridge  event trigger fires.
	BatchSize *float64 `json:"batchSize,omitempty" tf:"batch_size,omitempty"`

	// Window of time in seconds after which EventBridge event trigger fires. Window starts when first event is received. Default value is 900.
	BatchWindow *float64 `json:"batchWindow,omitempty" tf:"batch_window,omitempty"`
}

type EventBatchingConditionObservation struct {

	// Number of events that must be received from Amazon EventBridge before EventBridge  event trigger fires.
	BatchSize *float64 `json:"batchSize,omitempty" tf:"batch_size,omitempty"`

	// Window of time in seconds after which EventBridge event trigger fires. Window starts when first event is received. Default value is 900.
	BatchWindow *float64 `json:"batchWindow,omitempty" tf:"batch_window,omitempty"`
}

type EventBatchingConditionParameters struct {

	// Number of events that must be received from Amazon EventBridge before EventBridge  event trigger fires.
	// +kubebuilder:validation:Optional
	BatchSize *float64 `json:"batchSize" tf:"batch_size,omitempty"`

	// Window of time in seconds after which EventBridge event trigger fires. Window starts when first event is received. Default value is 900.
	// +kubebuilder:validation:Optional
	BatchWindow *float64 `json:"batchWindow,omitempty" tf:"batch_window,omitempty"`
}

type PredicateInitParameters struct {

	// A list of the conditions that determine when the trigger will fire. See Conditions.
	Conditions []ConditionsInitParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// How to handle multiple conditions. Defaults to AND. Valid values are AND or ANY.
	Logical *string `json:"logical,omitempty" tf:"logical,omitempty"`
}

type PredicateObservation struct {

	// A list of the conditions that determine when the trigger will fire. See Conditions.
	Conditions []ConditionsObservation `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// How to handle multiple conditions. Defaults to AND. Valid values are AND or ANY.
	Logical *string `json:"logical,omitempty" tf:"logical,omitempty"`
}

type PredicateParameters struct {

	// A list of the conditions that determine when the trigger will fire. See Conditions.
	// +kubebuilder:validation:Optional
	Conditions []ConditionsParameters `json:"conditions" tf:"conditions,omitempty"`

	// How to handle multiple conditions. Defaults to AND. Valid values are AND or ANY.
	// +kubebuilder:validation:Optional
	Logical *string `json:"logical,omitempty" tf:"logical,omitempty"`
}

type TriggerInitParameters struct {

	// –  List of actions initiated by this trigger when it fires. See Actions Below.
	Actions []ActionsInitParameters `json:"actions,omitempty" tf:"actions,omitempty"`

	// –  A description of the new trigger.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// –  Start the trigger. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Batch condition that must be met (specified number of events received or batch time window expired) before EventBridge event trigger fires. See Event Batching Condition.
	EventBatchingCondition []EventBatchingConditionInitParameters `json:"eventBatchingCondition,omitempty" tf:"event_batching_condition,omitempty"`

	// –  A predicate to specify when the new trigger should fire. Required when trigger type is CONDITIONAL. See Predicate Below.
	Predicate []PredicateInitParameters `json:"predicate,omitempty" tf:"predicate,omitempty"`

	// Based Schedules for Jobs and Crawlers
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// –  Set to true to start SCHEDULED and CONDITIONAL triggers when created. True is not supported for ON_DEMAND triggers.
	StartOnCreation *bool `json:"startOnCreation,omitempty" tf:"start_on_creation,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// –  The type of trigger. Valid values are CONDITIONAL, EVENT, ON_DEMAND, and SCHEDULED.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// A workflow to which the trigger should be associated to. Every workflow graph (DAG) needs a starting trigger (ON_DEMAND or SCHEDULED type) and can contain multiple additional CONDITIONAL triggers.
	WorkflowName *string `json:"workflowName,omitempty" tf:"workflow_name,omitempty"`
}

type TriggerObservation struct {

	// –  List of actions initiated by this trigger when it fires. See Actions Below.
	Actions []ActionsObservation `json:"actions,omitempty" tf:"actions,omitempty"`

	// Amazon Resource Name (ARN) of Glue Trigger
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// –  A description of the new trigger.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// –  Start the trigger. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Batch condition that must be met (specified number of events received or batch time window expired) before EventBridge event trigger fires. See Event Batching Condition.
	EventBatchingCondition []EventBatchingConditionObservation `json:"eventBatchingCondition,omitempty" tf:"event_batching_condition,omitempty"`

	// Trigger name
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// –  A predicate to specify when the new trigger should fire. Required when trigger type is CONDITIONAL. See Predicate Below.
	Predicate []PredicateObservation `json:"predicate,omitempty" tf:"predicate,omitempty"`

	// Based Schedules for Jobs and Crawlers
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// –  Set to true to start SCHEDULED and CONDITIONAL triggers when created. True is not supported for ON_DEMAND triggers.
	StartOnCreation *bool `json:"startOnCreation,omitempty" tf:"start_on_creation,omitempty"`

	// The condition job state. Currently, the values supported are SUCCEEDED, STOPPED, TIMEOUT and FAILED. If this is specified, job_name must also be specified. Conflicts with crawler_state.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// –  The type of trigger. Valid values are CONDITIONAL, EVENT, ON_DEMAND, and SCHEDULED.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// A workflow to which the trigger should be associated to. Every workflow graph (DAG) needs a starting trigger (ON_DEMAND or SCHEDULED type) and can contain multiple additional CONDITIONAL triggers.
	WorkflowName *string `json:"workflowName,omitempty" tf:"workflow_name,omitempty"`
}

type TriggerParameters struct {

	// –  List of actions initiated by this trigger when it fires. See Actions Below.
	// +kubebuilder:validation:Optional
	Actions []ActionsParameters `json:"actions,omitempty" tf:"actions,omitempty"`

	// –  A description of the new trigger.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// –  Start the trigger. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Batch condition that must be met (specified number of events received or batch time window expired) before EventBridge event trigger fires. See Event Batching Condition.
	// +kubebuilder:validation:Optional
	EventBatchingCondition []EventBatchingConditionParameters `json:"eventBatchingCondition,omitempty" tf:"event_batching_condition,omitempty"`

	// –  A predicate to specify when the new trigger should fire. Required when trigger type is CONDITIONAL. See Predicate Below.
	// +kubebuilder:validation:Optional
	Predicate []PredicateParameters `json:"predicate,omitempty" tf:"predicate,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Based Schedules for Jobs and Crawlers
	// +kubebuilder:validation:Optional
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// –  Set to true to start SCHEDULED and CONDITIONAL triggers when created. True is not supported for ON_DEMAND triggers.
	// +kubebuilder:validation:Optional
	StartOnCreation *bool `json:"startOnCreation,omitempty" tf:"start_on_creation,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// –  The type of trigger. Valid values are CONDITIONAL, EVENT, ON_DEMAND, and SCHEDULED.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// A workflow to which the trigger should be associated to. Every workflow graph (DAG) needs a starting trigger (ON_DEMAND or SCHEDULED type) and can contain multiple additional CONDITIONAL triggers.
	// +kubebuilder:validation:Optional
	WorkflowName *string `json:"workflowName,omitempty" tf:"workflow_name,omitempty"`
}

// TriggerSpec defines the desired state of Trigger
type TriggerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TriggerParameters `json:"forProvider"`
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
	InitProvider TriggerInitParameters `json:"initProvider,omitempty"`
}

// TriggerStatus defines the observed state of Trigger.
type TriggerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TriggerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Trigger is the Schema for the Triggers API. Manages a Glue Trigger resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Trigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.actions) || (has(self.initProvider) && has(self.initProvider.actions))",message="spec.forProvider.actions is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   TriggerSpec   `json:"spec"`
	Status TriggerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TriggerList contains a list of Triggers
type TriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Trigger `json:"items"`
}

// Repository type metadata.
var (
	Trigger_Kind             = "Trigger"
	Trigger_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Trigger_Kind}.String()
	Trigger_KindAPIVersion   = Trigger_Kind + "." + CRDGroupVersion.String()
	Trigger_GroupVersionKind = CRDGroupVersion.WithKind(Trigger_Kind)
)

func init() {
	SchemeBuilder.Register(&Trigger{}, &TriggerList{})
}
