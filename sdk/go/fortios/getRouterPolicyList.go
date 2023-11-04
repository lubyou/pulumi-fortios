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

func GetRouterPolicyList(ctx *pulumi.Context, args *GetRouterPolicyListArgs, opts ...pulumi.InvokeOption) (*GetRouterPolicyListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRouterPolicyListResult
	err := ctx.Invoke("fortios:index/getRouterPolicyList:GetRouterPolicyList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterPolicyList.
type GetRouterPolicyListArgs struct {
	Filter    *string `pulumi:"filter"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterPolicyList.
type GetRouterPolicyListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	SeqNumlists []int   `pulumi:"seqNumlists"`
	Vdomparam   *string `pulumi:"vdomparam"`
}

func GetRouterPolicyListOutput(ctx *pulumi.Context, args GetRouterPolicyListOutputArgs, opts ...pulumi.InvokeOption) GetRouterPolicyListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRouterPolicyListResult, error) {
			args := v.(GetRouterPolicyListArgs)
			r, err := GetRouterPolicyList(ctx, &args, opts...)
			var s GetRouterPolicyListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRouterPolicyListResultOutput)
}

// A collection of arguments for invoking GetRouterPolicyList.
type GetRouterPolicyListOutputArgs struct {
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetRouterPolicyListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouterPolicyListArgs)(nil)).Elem()
}

// A collection of values returned by GetRouterPolicyList.
type GetRouterPolicyListResultOutput struct{ *pulumi.OutputState }

func (GetRouterPolicyListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouterPolicyListResult)(nil)).Elem()
}

func (o GetRouterPolicyListResultOutput) ToGetRouterPolicyListResultOutput() GetRouterPolicyListResultOutput {
	return o
}

func (o GetRouterPolicyListResultOutput) ToGetRouterPolicyListResultOutputWithContext(ctx context.Context) GetRouterPolicyListResultOutput {
	return o
}

func (o GetRouterPolicyListResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetRouterPolicyListResult] {
	return pulumix.Output[GetRouterPolicyListResult]{
		OutputState: o.OutputState,
	}
}

func (o GetRouterPolicyListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouterPolicyListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRouterPolicyListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRouterPolicyListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRouterPolicyListResultOutput) SeqNumlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetRouterPolicyListResult) []int { return v.SeqNumlists }).(pulumi.IntArrayOutput)
}

func (o GetRouterPolicyListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouterPolicyListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRouterPolicyListResultOutput{})
}
