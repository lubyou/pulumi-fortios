// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSystemAutoupdatePushUpdate(ctx *pulumi.Context, args *LookupSystemAutoupdatePushUpdateArgs, opts ...pulumi.InvokeOption) (*LookupSystemAutoupdatePushUpdateResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemAutoupdatePushUpdateResult
	err := ctx.Invoke("fortios:index/getSystemAutoupdatePushUpdate:GetSystemAutoupdatePushUpdate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemAutoupdatePushUpdate.
type LookupSystemAutoupdatePushUpdateArgs struct {
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemAutoupdatePushUpdate.
type LookupSystemAutoupdatePushUpdateResult struct {
	Address string `pulumi:"address"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Override  string  `pulumi:"override"`
	Port      int     `pulumi:"port"`
	Status    string  `pulumi:"status"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupSystemAutoupdatePushUpdateOutput(ctx *pulumi.Context, args LookupSystemAutoupdatePushUpdateOutputArgs, opts ...pulumi.InvokeOption) LookupSystemAutoupdatePushUpdateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemAutoupdatePushUpdateResult, error) {
			args := v.(LookupSystemAutoupdatePushUpdateArgs)
			r, err := LookupSystemAutoupdatePushUpdate(ctx, &args, opts...)
			var s LookupSystemAutoupdatePushUpdateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemAutoupdatePushUpdateResultOutput)
}

// A collection of arguments for invoking GetSystemAutoupdatePushUpdate.
type LookupSystemAutoupdatePushUpdateOutputArgs struct {
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemAutoupdatePushUpdateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemAutoupdatePushUpdateArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemAutoupdatePushUpdate.
type LookupSystemAutoupdatePushUpdateResultOutput struct{ *pulumi.OutputState }

func (LookupSystemAutoupdatePushUpdateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemAutoupdatePushUpdateResult)(nil)).Elem()
}

func (o LookupSystemAutoupdatePushUpdateResultOutput) ToLookupSystemAutoupdatePushUpdateResultOutput() LookupSystemAutoupdatePushUpdateResultOutput {
	return o
}

func (o LookupSystemAutoupdatePushUpdateResultOutput) ToLookupSystemAutoupdatePushUpdateResultOutputWithContext(ctx context.Context) LookupSystemAutoupdatePushUpdateResultOutput {
	return o
}

func (o LookupSystemAutoupdatePushUpdateResultOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutoupdatePushUpdateResult) string { return v.Address }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemAutoupdatePushUpdateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutoupdatePushUpdateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSystemAutoupdatePushUpdateResultOutput) Override() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutoupdatePushUpdateResult) string { return v.Override }).(pulumi.StringOutput)
}

func (o LookupSystemAutoupdatePushUpdateResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemAutoupdatePushUpdateResult) int { return v.Port }).(pulumi.IntOutput)
}

func (o LookupSystemAutoupdatePushUpdateResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutoupdatePushUpdateResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupSystemAutoupdatePushUpdateResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemAutoupdatePushUpdateResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemAutoupdatePushUpdateResultOutput{})
}
