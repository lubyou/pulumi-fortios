// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSystemAutomationDestination(ctx *pulumi.Context, args *LookupSystemAutomationDestinationArgs, opts ...pulumi.InvokeOption) (*LookupSystemAutomationDestinationResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemAutomationDestinationResult
	err := ctx.Invoke("fortios:index/getSystemAutomationDestination:GetSystemAutomationDestination", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemAutomationDestination.
type LookupSystemAutomationDestinationArgs struct {
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemAutomationDestination.
type LookupSystemAutomationDestinationResult struct {
	Destinations []GetSystemAutomationDestinationDestination `pulumi:"destinations"`
	HaGroupId    int                                         `pulumi:"haGroupId"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Name      string  `pulumi:"name"`
	Type      string  `pulumi:"type"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupSystemAutomationDestinationOutput(ctx *pulumi.Context, args LookupSystemAutomationDestinationOutputArgs, opts ...pulumi.InvokeOption) LookupSystemAutomationDestinationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemAutomationDestinationResult, error) {
			args := v.(LookupSystemAutomationDestinationArgs)
			r, err := LookupSystemAutomationDestination(ctx, &args, opts...)
			var s LookupSystemAutomationDestinationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemAutomationDestinationResultOutput)
}

// A collection of arguments for invoking GetSystemAutomationDestination.
type LookupSystemAutomationDestinationOutputArgs struct {
	Name      pulumi.StringInput    `pulumi:"name"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemAutomationDestinationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemAutomationDestinationArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemAutomationDestination.
type LookupSystemAutomationDestinationResultOutput struct{ *pulumi.OutputState }

func (LookupSystemAutomationDestinationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemAutomationDestinationResult)(nil)).Elem()
}

func (o LookupSystemAutomationDestinationResultOutput) ToLookupSystemAutomationDestinationResultOutput() LookupSystemAutomationDestinationResultOutput {
	return o
}

func (o LookupSystemAutomationDestinationResultOutput) ToLookupSystemAutomationDestinationResultOutputWithContext(ctx context.Context) LookupSystemAutomationDestinationResultOutput {
	return o
}

func (o LookupSystemAutomationDestinationResultOutput) Destinations() GetSystemAutomationDestinationDestinationArrayOutput {
	return o.ApplyT(func(v LookupSystemAutomationDestinationResult) []GetSystemAutomationDestinationDestination {
		return v.Destinations
	}).(GetSystemAutomationDestinationDestinationArrayOutput)
}

func (o LookupSystemAutomationDestinationResultOutput) HaGroupId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemAutomationDestinationResult) int { return v.HaGroupId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemAutomationDestinationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationDestinationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSystemAutomationDestinationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationDestinationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSystemAutomationDestinationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationDestinationResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSystemAutomationDestinationResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemAutomationDestinationResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemAutomationDestinationResultOutput{})
}
