// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

func GetRouterAspathListList(ctx *pulumi.Context, args *GetRouterAspathListListArgs, opts ...pulumi.InvokeOption) (*GetRouterAspathListListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRouterAspathListListResult
	err := ctx.Invoke("fortios:index/getRouterAspathListList:GetRouterAspathListList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterAspathListList.
type GetRouterAspathListListArgs struct {
	Filter    *string `pulumi:"filter"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterAspathListList.
type GetRouterAspathListListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetRouterAspathListListOutput(ctx *pulumi.Context, args GetRouterAspathListListOutputArgs, opts ...pulumi.InvokeOption) GetRouterAspathListListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRouterAspathListListResult, error) {
			args := v.(GetRouterAspathListListArgs)
			r, err := GetRouterAspathListList(ctx, &args, opts...)
			var s GetRouterAspathListListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRouterAspathListListResultOutput)
}

// A collection of arguments for invoking GetRouterAspathListList.
type GetRouterAspathListListOutputArgs struct {
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetRouterAspathListListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouterAspathListListArgs)(nil)).Elem()
}

// A collection of values returned by GetRouterAspathListList.
type GetRouterAspathListListResultOutput struct{ *pulumi.OutputState }

func (GetRouterAspathListListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouterAspathListListResult)(nil)).Elem()
}

func (o GetRouterAspathListListResultOutput) ToGetRouterAspathListListResultOutput() GetRouterAspathListListResultOutput {
	return o
}

func (o GetRouterAspathListListResultOutput) ToGetRouterAspathListListResultOutputWithContext(ctx context.Context) GetRouterAspathListListResultOutput {
	return o
}

func (o GetRouterAspathListListResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetRouterAspathListListResult] {
	return pulumix.Output[GetRouterAspathListListResult]{
		OutputState: o.OutputState,
	}
}

func (o GetRouterAspathListListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouterAspathListListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRouterAspathListListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRouterAspathListListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRouterAspathListListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRouterAspathListListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetRouterAspathListListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouterAspathListListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRouterAspathListListResultOutput{})
}
