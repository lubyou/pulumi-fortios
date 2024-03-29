// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupUserSaml(ctx *pulumi.Context, args *LookupUserSamlArgs, opts ...pulumi.InvokeOption) (*LookupUserSamlResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupUserSamlResult
	err := ctx.Invoke("fortios:index/getUserSaml:GetUserSaml", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetUserSaml.
type LookupUserSamlArgs struct {
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetUserSaml.
type LookupUserSamlResult struct {
	AdfsClaim      string `pulumi:"adfsClaim"`
	AuthUrl        string `pulumi:"authUrl"`
	Cert           string `pulumi:"cert"`
	ClockTolerance int    `pulumi:"clockTolerance"`
	DigestMethod   string `pulumi:"digestMethod"`
	EntityId       string `pulumi:"entityId"`
	GroupClaimType string `pulumi:"groupClaimType"`
	GroupName      string `pulumi:"groupName"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string  `pulumi:"id"`
	IdpCert            string  `pulumi:"idpCert"`
	IdpEntityId        string  `pulumi:"idpEntityId"`
	IdpSingleLogoutUrl string  `pulumi:"idpSingleLogoutUrl"`
	IdpSingleSignOnUrl string  `pulumi:"idpSingleSignOnUrl"`
	LimitRelaystate    string  `pulumi:"limitRelaystate"`
	Name               string  `pulumi:"name"`
	Reauth             string  `pulumi:"reauth"`
	SingleLogoutUrl    string  `pulumi:"singleLogoutUrl"`
	SingleSignOnUrl    string  `pulumi:"singleSignOnUrl"`
	UserClaimType      string  `pulumi:"userClaimType"`
	UserName           string  `pulumi:"userName"`
	Vdomparam          *string `pulumi:"vdomparam"`
}

func LookupUserSamlOutput(ctx *pulumi.Context, args LookupUserSamlOutputArgs, opts ...pulumi.InvokeOption) LookupUserSamlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserSamlResult, error) {
			args := v.(LookupUserSamlArgs)
			r, err := LookupUserSaml(ctx, &args, opts...)
			var s LookupUserSamlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserSamlResultOutput)
}

// A collection of arguments for invoking GetUserSaml.
type LookupUserSamlOutputArgs struct {
	Name      pulumi.StringInput    `pulumi:"name"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupUserSamlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserSamlArgs)(nil)).Elem()
}

// A collection of values returned by GetUserSaml.
type LookupUserSamlResultOutput struct{ *pulumi.OutputState }

func (LookupUserSamlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserSamlResult)(nil)).Elem()
}

func (o LookupUserSamlResultOutput) ToLookupUserSamlResultOutput() LookupUserSamlResultOutput {
	return o
}

func (o LookupUserSamlResultOutput) ToLookupUserSamlResultOutputWithContext(ctx context.Context) LookupUserSamlResultOutput {
	return o
}

func (o LookupUserSamlResultOutput) AdfsClaim() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.AdfsClaim }).(pulumi.StringOutput)
}

func (o LookupUserSamlResultOutput) AuthUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.AuthUrl }).(pulumi.StringOutput)
}

func (o LookupUserSamlResultOutput) Cert() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.Cert }).(pulumi.StringOutput)
}

func (o LookupUserSamlResultOutput) ClockTolerance() pulumi.IntOutput {
	return o.ApplyT(func(v LookupUserSamlResult) int { return v.ClockTolerance }).(pulumi.IntOutput)
}

func (o LookupUserSamlResultOutput) DigestMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.DigestMethod }).(pulumi.StringOutput)
}

func (o LookupUserSamlResultOutput) EntityId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.EntityId }).(pulumi.StringOutput)
}

func (o LookupUserSamlResultOutput) GroupClaimType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.GroupClaimType }).(pulumi.StringOutput)
}

func (o LookupUserSamlResultOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.GroupName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupUserSamlResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupUserSamlResultOutput) IdpCert() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.IdpCert }).(pulumi.StringOutput)
}

func (o LookupUserSamlResultOutput) IdpEntityId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.IdpEntityId }).(pulumi.StringOutput)
}

func (o LookupUserSamlResultOutput) IdpSingleLogoutUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.IdpSingleLogoutUrl }).(pulumi.StringOutput)
}

func (o LookupUserSamlResultOutput) IdpSingleSignOnUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.IdpSingleSignOnUrl }).(pulumi.StringOutput)
}

func (o LookupUserSamlResultOutput) LimitRelaystate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.LimitRelaystate }).(pulumi.StringOutput)
}

func (o LookupUserSamlResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupUserSamlResultOutput) Reauth() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.Reauth }).(pulumi.StringOutput)
}

func (o LookupUserSamlResultOutput) SingleLogoutUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.SingleLogoutUrl }).(pulumi.StringOutput)
}

func (o LookupUserSamlResultOutput) SingleSignOnUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.SingleSignOnUrl }).(pulumi.StringOutput)
}

func (o LookupUserSamlResultOutput) UserClaimType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.UserClaimType }).(pulumi.StringOutput)
}

func (o LookupUserSamlResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.UserName }).(pulumi.StringOutput)
}

func (o LookupUserSamlResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserSamlResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserSamlResultOutput{})
}
