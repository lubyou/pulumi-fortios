// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetRouterStatic6List(ctx *pulumi.Context, args *GetRouterStatic6ListArgs, opts ...pulumi.InvokeOption) (*GetRouterStatic6ListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRouterStatic6ListResult
	err := ctx.Invoke("fortios:index/getRouterStatic6List:GetRouterStatic6List", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterStatic6List.
type GetRouterStatic6ListArgs struct {
	Filter    *string `pulumi:"filter"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterStatic6List.
type GetRouterStatic6ListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	SeqNumlists []int   `pulumi:"seqNumlists"`
	Vdomparam   *string `pulumi:"vdomparam"`
}

func GetRouterStatic6ListOutput(ctx *pulumi.Context, args GetRouterStatic6ListOutputArgs, opts ...pulumi.InvokeOption) GetRouterStatic6ListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRouterStatic6ListResult, error) {
			args := v.(GetRouterStatic6ListArgs)
			r, err := GetRouterStatic6List(ctx, &args, opts...)
			var s GetRouterStatic6ListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRouterStatic6ListResultOutput)
}

// A collection of arguments for invoking GetRouterStatic6List.
type GetRouterStatic6ListOutputArgs struct {
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetRouterStatic6ListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouterStatic6ListArgs)(nil)).Elem()
}

// A collection of values returned by GetRouterStatic6List.
type GetRouterStatic6ListResultOutput struct{ *pulumi.OutputState }

func (GetRouterStatic6ListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouterStatic6ListResult)(nil)).Elem()
}

func (o GetRouterStatic6ListResultOutput) ToGetRouterStatic6ListResultOutput() GetRouterStatic6ListResultOutput {
	return o
}

func (o GetRouterStatic6ListResultOutput) ToGetRouterStatic6ListResultOutputWithContext(ctx context.Context) GetRouterStatic6ListResultOutput {
	return o
}

func (o GetRouterStatic6ListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouterStatic6ListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRouterStatic6ListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRouterStatic6ListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRouterStatic6ListResultOutput) SeqNumlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetRouterStatic6ListResult) []int { return v.SeqNumlists }).(pulumi.IntArrayOutput)
}

func (o GetRouterStatic6ListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouterStatic6ListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRouterStatic6ListResultOutput{})
}
