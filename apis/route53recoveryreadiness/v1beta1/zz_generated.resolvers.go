/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/cloudwatch/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ResourceSet.
func (mg *ResourceSet) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Resources); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Resources[i3].ResourceArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.Resources[i3].ResourceArnRef,
			Selector:     mg.Spec.ForProvider.Resources[i3].ResourceArnSelector,
			To: reference.To{
				List:    &v1beta1.MetricAlarmList{},
				Managed: &v1beta1.MetricAlarm{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Resources[i3].ResourceArn")
		}
		mg.Spec.ForProvider.Resources[i3].ResourceArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Resources[i3].ResourceArnRef = rsp.ResolvedReference

	}

	return nil
}
