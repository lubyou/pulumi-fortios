// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSystemVdomException(ctx *pulumi.Context, args *LookupSystemVdomExceptionArgs, opts ...pulumi.InvokeOption) (*LookupSystemVdomExceptionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSystemVdomExceptionResult
	err := ctx.Invoke("fortios:index/getSystemVdomException:GetSystemVdomException", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemVdomException.
type LookupSystemVdomExceptionArgs struct {
	Fosid     int     `pulumi:"fosid"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemVdomException.
type LookupSystemVdomExceptionResult struct {
	Fosid int `pulumi:"fosid"`
	// The provider-assigned unique ID for this managed resource.
	Id        string                       `pulumi:"id"`
	Object    string                       `pulumi:"object"`
	Oid       int                          `pulumi:"oid"`
	Scope     string                       `pulumi:"scope"`
	Vdomparam *string                      `pulumi:"vdomparam"`
	Vdoms     []GetSystemVdomExceptionVdom `pulumi:"vdoms"`
}

func LookupSystemVdomExceptionOutput(ctx *pulumi.Context, args LookupSystemVdomExceptionOutputArgs, opts ...pulumi.InvokeOption) LookupSystemVdomExceptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemVdomExceptionResult, error) {
			args := v.(LookupSystemVdomExceptionArgs)
			r, err := LookupSystemVdomException(ctx, &args, opts...)
			var s LookupSystemVdomExceptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemVdomExceptionResultOutput)
}

// A collection of arguments for invoking GetSystemVdomException.
type LookupSystemVdomExceptionOutputArgs struct {
	Fosid     pulumi.IntInput       `pulumi:"fosid"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemVdomExceptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemVdomExceptionArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemVdomException.
type LookupSystemVdomExceptionResultOutput struct{ *pulumi.OutputState }

func (LookupSystemVdomExceptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemVdomExceptionResult)(nil)).Elem()
}

func (o LookupSystemVdomExceptionResultOutput) ToLookupSystemVdomExceptionResultOutput() LookupSystemVdomExceptionResultOutput {
	return o
}

func (o LookupSystemVdomExceptionResultOutput) ToLookupSystemVdomExceptionResultOutputWithContext(ctx context.Context) LookupSystemVdomExceptionResultOutput {
	return o
}

func (o LookupSystemVdomExceptionResultOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemVdomExceptionResult) int { return v.Fosid }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemVdomExceptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemVdomExceptionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSystemVdomExceptionResultOutput) Object() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemVdomExceptionResult) string { return v.Object }).(pulumi.StringOutput)
}

func (o LookupSystemVdomExceptionResultOutput) Oid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemVdomExceptionResult) int { return v.Oid }).(pulumi.IntOutput)
}

func (o LookupSystemVdomExceptionResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemVdomExceptionResult) string { return v.Scope }).(pulumi.StringOutput)
}

func (o LookupSystemVdomExceptionResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemVdomExceptionResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o LookupSystemVdomExceptionResultOutput) Vdoms() GetSystemVdomExceptionVdomArrayOutput {
	return o.ApplyT(func(v LookupSystemVdomExceptionResult) []GetSystemVdomExceptionVdom { return v.Vdoms }).(GetSystemVdomExceptionVdomArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemVdomExceptionResultOutput{})
}
