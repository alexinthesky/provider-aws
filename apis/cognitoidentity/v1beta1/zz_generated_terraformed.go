// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/crossplane/upjet/pkg/resource"
	"github.com/crossplane/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this Pool
func (mg *Pool) GetTerraformResourceType() string {
	return "aws_cognito_identity_pool"
}

// GetConnectionDetailsMapping for this Pool
func (tr *Pool) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Pool
func (tr *Pool) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Pool
func (tr *Pool) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Pool
func (tr *Pool) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Pool
func (tr *Pool) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Pool
func (tr *Pool) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this Pool
func (tr *Pool) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this Pool using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Pool) LateInitialize(attrs []byte) (bool, error) {
	params := &PoolParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Pool) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this CognitoIdentityPoolProviderPrincipalTag
func (mg *CognitoIdentityPoolProviderPrincipalTag) GetTerraformResourceType() string {
	return "aws_cognito_identity_pool_provider_principal_tag"
}

// GetConnectionDetailsMapping for this CognitoIdentityPoolProviderPrincipalTag
func (tr *CognitoIdentityPoolProviderPrincipalTag) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this CognitoIdentityPoolProviderPrincipalTag
func (tr *CognitoIdentityPoolProviderPrincipalTag) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this CognitoIdentityPoolProviderPrincipalTag
func (tr *CognitoIdentityPoolProviderPrincipalTag) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this CognitoIdentityPoolProviderPrincipalTag
func (tr *CognitoIdentityPoolProviderPrincipalTag) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this CognitoIdentityPoolProviderPrincipalTag
func (tr *CognitoIdentityPoolProviderPrincipalTag) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this CognitoIdentityPoolProviderPrincipalTag
func (tr *CognitoIdentityPoolProviderPrincipalTag) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this CognitoIdentityPoolProviderPrincipalTag
func (tr *CognitoIdentityPoolProviderPrincipalTag) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this CognitoIdentityPoolProviderPrincipalTag using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *CognitoIdentityPoolProviderPrincipalTag) LateInitialize(attrs []byte) (bool, error) {
	params := &CognitoIdentityPoolProviderPrincipalTagParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *CognitoIdentityPoolProviderPrincipalTag) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this PoolRolesAttachment
func (mg *PoolRolesAttachment) GetTerraformResourceType() string {
	return "aws_cognito_identity_pool_roles_attachment"
}

// GetConnectionDetailsMapping for this PoolRolesAttachment
func (tr *PoolRolesAttachment) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this PoolRolesAttachment
func (tr *PoolRolesAttachment) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this PoolRolesAttachment
func (tr *PoolRolesAttachment) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this PoolRolesAttachment
func (tr *PoolRolesAttachment) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this PoolRolesAttachment
func (tr *PoolRolesAttachment) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this PoolRolesAttachment
func (tr *PoolRolesAttachment) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this PoolRolesAttachment
func (tr *PoolRolesAttachment) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this PoolRolesAttachment using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *PoolRolesAttachment) LateInitialize(attrs []byte) (bool, error) {
	params := &PoolRolesAttachmentParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *PoolRolesAttachment) GetTerraformSchemaVersion() int {
	return 0
}
